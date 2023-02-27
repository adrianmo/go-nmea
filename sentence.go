package nmea

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

const (
	// TagBlockSep is the separator (slash `\`) that indicates start and end of tag block
	TagBlockSep = '\\'

	// SentenceStart is the token to indicate the start of a sentence.
	SentenceStart = "$"

	// SentenceStartEncapsulated is the token to indicate the start of encapsulated data.
	SentenceStartEncapsulated = "!"

	// ProprietarySentencePrefix is the token to indicate the start of parametric sentences.
	ProprietarySentencePrefix = 'P'

	// QuerySentencePostfix is the suffix token to indicate the Query sentences.
	QuerySentencePostfix = 'Q'

	// FieldSep is the token to delimit fields of a sentence.
	FieldSep = ","

	// ChecksumSep is the token to delimit the checksum of a sentence.
	ChecksumSep = "*"
)

var (
	// defaultSentenceParser exists for backwards compatibility reasons to allow global Parse/RegisterParser/MustRegisterParser
	// to work as they did before SentenceParser was added.
	defaultSentenceParser   = SentenceParser{Parsers: DefaultParsers()}
	defaultSentenceParserMu sync.Mutex
	defaultParsers          = DefaultParsers()
	customParsers           = map[string]struct{}{}
)

func init() {
	// for backwards compatibility support MTK. PMTK001 is correct an supported when using SentenceParser instance
	defaultSentenceParser.Parsers["$PMTK001"] = newMTK
}

// ParserFunc callback used to parse specific sentence variants
type ParserFunc func(s BaseSentence) (Sentence, error)

// NotSupportedError is returned when parsed sentence is not supported
type NotSupportedError struct {
	Prefix string
}

// Error returns error message
func (p *NotSupportedError) Error() string {
	return fmt.Sprintf("nmea: sentence prefix '%s' not supported", p.Prefix)
}

// Sentence interface for all NMEA sentence
type Sentence interface {
	fmt.Stringer
	Prefix() string
	DataType() string
	TalkerID() string
}

// BaseSentence contains the information about the NMEA sentence
type BaseSentence struct {
	Talker   string   // The talker id (e.g GP)
	Type     string   // The data type (e.g GSA)
	Fields   []string // Array of fields
	Checksum string   // The (raw) Checksum
	Raw      string   // The raw NMEA sentence received
	TagBlock TagBlock // NMEA tagblock
}

// Prefix returns the talker and type of message
func (s BaseSentence) Prefix() string {
	return s.Talker + s.Type
}

// DataType returns the type of the message
func (s BaseSentence) DataType() string {
	return s.Type
}

// TalkerID returns the talker of the message
func (s BaseSentence) TalkerID() string {
	return s.Talker
}

// String formats the sentence into a string
func (s BaseSentence) String() string { return s.Raw }

// SentenceParser is configurable parser instance to parse raw input into NMEA0183 Sentence
//
// SentenceParser fields/methods are not co-routine safe!
type SentenceParser struct {
	// Parsers allows registering additional parsers
	Parsers map[string]ParserFunc

	// ParsePrefix takes in the sentence first field (NMEA0183 address) and splits it into a talker id and sentence type
	ParsePrefix func(prefix string) (talkerID string, sentence string, err error)

	// CheckCRC allows custom handling of checksum checking based on parsed sentence
	CheckCRC func(sentence BaseSentence, rawFields string) error

	// OnTagBlock is callback to handle all parsed tag blocks even for lines containing only a tag block and
	// allows to track multiline tag group separate lines. Logic how to combine/assemble multiline tag group
	// should be implemented outside SentenceParser
	// OnTagBlock is called before actual sentence part is parsed. When callback returns an error sentence parsing will
	// not be done and Parse returns early with the returned error.
	//
	// Example of group of 3:
	// \g:1-3-1234,s:r3669961,c:1120959341*hh\
	// \g:2-3-1234*hh\!ABVDM,1,1,1,B,.....,0*hh
	// \g:3-3-1234*hh\$ABVSI,r3669961,1,013536.96326433,1386,-98,,*hh
	OnTagBlock func(tagBlock TagBlock) error
}

func (p *SentenceParser) parseBaseSentence(raw string) (BaseSentence, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return BaseSentence{}, errors.New("nmea: can not parse empty input")
	}

	var (
		tagBlock TagBlock
		err      error
	)

	if startOfTagBlock := strings.IndexByte(raw, TagBlockSep); startOfTagBlock != -1 {
		// tag block is always at the start of line (unless IEC 61162-450). Starts with `\` and ends with `\` and has valid sentence
		// following or <CR><LF>
		//
		// Note: tag block group can span multiple lines but we only parse ones that have sentence
		endOfTagBlock := strings.LastIndexByte(raw, TagBlockSep)
		if endOfTagBlock <= startOfTagBlock {
			return BaseSentence{}, fmt.Errorf("nmea: sentence tag block is missing '\\' at the end")
		}
		tagBlock, err = parseTagBlock(raw[startOfTagBlock+1 : endOfTagBlock])
		if err != nil {
			return BaseSentence{}, err
		}
		if p.OnTagBlock != nil {
			if err := p.OnTagBlock(tagBlock); err != nil {
				return BaseSentence{}, err
			}
		}
		raw = raw[endOfTagBlock+1:]
	}

	startIndex := strings.IndexAny(raw, SentenceStart+SentenceStartEncapsulated)
	if startIndex != 0 {
		return BaseSentence{}, errors.New("nmea: sentence does not start with a '$' or '!'")
	}
	checksumSepIndex := strings.Index(raw, ChecksumSep)
	rawFields := raw[startIndex+1:]
	checksumRaw := ""
	if checksumSepIndex != -1 {
		rawFields = raw[startIndex+1 : checksumSepIndex]
		checksumRaw = strings.ToUpper(raw[checksumSepIndex+1:])
	}
	// TODO: fields can contain encoded chars that need to be unescaped. `^` 0x5e is escape character for HEX representation of ISO/IEC 8859-1 (ASCII) characters.
	// TODO: `^` itself is escaped as `^5e` and `,` is escaped as `^2c` etc. All reserved characters should be escaped/unescaped (See wikipedia https://en.wikipedia.org/wiki/NMEA_0183#Message_structure)
	fields := strings.Split(rawFields, FieldSep)

	var (
		talkerID string
		typ      string
	)
	if p.ParsePrefix == nil {
		talkerID, typ, err = ParsePrefix(fields[0])
	} else {
		talkerID, typ, err = p.ParsePrefix(fields[0])
	}
	if err != nil {
		return BaseSentence{}, err
	}
	sentence := BaseSentence{
		Talker:   talkerID,
		Type:     typ,
		Fields:   fields[1:],
		Checksum: checksumRaw,
		Raw:      raw,
		TagBlock: tagBlock,
	}
	if p.CheckCRC == nil {
		err = CheckCRC(sentence, rawFields)
	} else {
		err = p.CheckCRC(sentence, rawFields)
	}
	if err != nil {
		return BaseSentence{}, err
	}
	return sentence, nil
}

// CheckCRC is default implementation for checking sentence Checksum
func CheckCRC(sentence BaseSentence, rawFields string) error {
	if sentence.Checksum == "" {
		return fmt.Errorf("nmea: sentence does not contain checksum separator")
	}
	if checksum := Checksum(rawFields); checksum != sentence.Checksum {
		return fmt.Errorf("nmea: sentence checksum mismatch [%s != %s]", checksum, sentence.Checksum)
	}
	return nil
}

// ParsePrefix is default implementation for prefix parsing. It takes in the sentence first field and splits it into
// a talker id and sentence type.
func ParsePrefix(prefix string) (string, string, error) {
	if prefix == "" {
		return "", "", errors.New("nmea: sentence prefix is empty")
	}
	// proprietary sentences start with `P` + sentence type. By NMEA0183 spec they should be 5 character long,
	// In this case we allow sentence type to be longer as there are plenty of examples with longer sentence
	// types (PSKPDPT, PMTK001 etc)
	if prefix[0] == ProprietarySentencePrefix {
		return string(ProprietarySentencePrefix), prefix[1:], nil
	}
	// valid prefix, by the NMEA0183 standard, is 5 character long:
	// a) talkerID (2) + sentence identifier (3)
	// b) talkerID (2) + destinationID (2) + 'Q' (1)
	if len(prefix) == 5 {
		// query sentence is a special type of sentence in NMEA standard that is used for a listener to request a
		// particular sentence from a talker.
		// Query prefix consist of: XXYYQ, XX - requester talker ID, YY - requestee/destination talker ID, `Q` - query type
		// Destination talker ID is handled/parsed by newQuery function
		if prefix[4] == QuerySentencePostfix {
			return prefix[:2], string(QuerySentencePostfix), nil
		}
		return prefix[:2], prefix[2:], nil
	}
	// this is catch all for other invalid NMEA0183 implementations. When prefix is shorter or longer than 5 characters
	// we use everything as sentence type. This way custom parser could be created that matches this off-spec prefix.
	return "", prefix, nil
}

// Checksum xor all the bytes in a string and return it
// as an uppercase hex string
func Checksum(s string) string {
	var checksum uint8
	for i := 0; i < len(s); i++ {
		checksum ^= s[i]
	}
	return fmt.Sprintf("%02X", checksum)
}

// MustRegisterParser register a custom parser or panic
func MustRegisterParser(sentenceType string, parser ParserFunc) {
	if err := RegisterParser(sentenceType, parser); err != nil {
		panic(err)
	}
}

// RegisterParser register a custom parser
func RegisterParser(sentenceType string, parser ParserFunc) error {
	defaultSentenceParserMu.Lock()
	defer defaultSentenceParserMu.Unlock()
	if _, ok := customParsers[sentenceType]; ok {
		return fmt.Errorf("nmea: parser for sentence type '%q' already exists", sentenceType)
	}
	defaultSentenceParser.Parsers[sentenceType] = parser
	customParsers[sentenceType] = struct{}{}
	return nil
}

// Parse parses the given string into the correct sentence type.
func Parse(raw string) (Sentence, error) {
	defaultSentenceParserMu.Lock()
	defer defaultSentenceParserMu.Unlock()

	return defaultSentenceParser.Parse(raw)
}

// Parse parses the given string into the correct sentence type.
// This method is not co-routine safe!
func (p *SentenceParser) Parse(raw string) (Sentence, error) {
	s, err := p.parseBaseSentence(raw)
	if err != nil {
		return nil, err
	}
	parsers := p.Parsers
	if parsers == nil {
		parsers = defaultParsers
	}
	// lookup parser in decreasing order of specificity
	start := "$"
	if strings.HasPrefix(s.Raw, "!") {
		start = "!"
	}
	if parse, ok := parsers[start+s.Prefix()]; ok {
		return parse(s)
	}
	if parse, ok := parsers[start+s.Type]; ok {
		return parse(s)
	}
	if parser, ok := parsers[s.Type]; ok {
		return parser(s)
	}
	return nil, &NotSupportedError{Prefix: s.Prefix()}
}

// DefaultParsers returns the default sentence parsers.
func DefaultParsers() map[string]ParserFunc {
	return map[string]ParserFunc{
		"$RMC":     newRMC,
		"$AAM":     newAAM,
		"$ACK":     newACK,
		"$ACN":     newACN,
		"$ALA":     newALA,
		"$ALC":     newALC,
		"$ALF":     newALF,
		"$ALR":     newALR,
		"$APB":     newAPB,
		"$ARC":     newARC,
		"$BEC":     newBEC,
		"$BOD":     newBOD,
		"$BWC":     newBWC,
		"$BWR":     newBWR,
		"$BWW":     newBWW,
		"$DOR":     newDOR,
		"$DSC":     newDSC,
		"$DSE":     newDSE,
		"$DTM":     newDTM,
		"$EVE":     newEVE,
		"$FIR":     newFIR,
		"$GGA":     newGGA,
		"$GSA":     newGSA,
		"$GLL":     newGLL,
		"$VTG":     newVTG,
		"$ZDA":     newZDA,
		"$PGRME":   newPGRME,
		"$PHTRO":   newPHTRO,
		"$PMTK001": newPMTK001,
		"$PRDID":   newPRDID,
		"$PSKPDPT": newPSKPDPT,
		"$PSONCMS": newPSONCMS,
		"$Q":       newQuery,
		"$GSV":     newGSV,
		"$HBT":     newHBT,
		"$HDG":     newHDG,
		"$HDT":     newHDT,
		"$HDM":     newHDM,
		"$HSC":     newHSC,
		"$GNS":     newGNS,
		"$THS":     newTHS,
		"$TLB":     newTLB,
		"$TLL":     newTLL,
		"$TTM":     newTTM,
		"$TXT":     newTXT,
		"$WPL":     newWPL,
		"$RMB":     newRMB,
		"$RPM":     newRPM,
		"$RSA":     newRSA,
		"$RSD":     newRSD,
		"$RTE":     newRTE,
		"$ROT":     newROT,
		"$VBW":     newVBW,
		"$VDR":     newVDR,
		"$VHW":     newVHW,
		"$VSD":     newVSD,
		"$VPW":     newVPW,
		"$VLW":     newVLW,
		"$VWR":     newVWR,
		"$VWT":     newVWT,
		"$DPT":     newDPT,
		"$DBT":     newDBT,
		"$DBK":     newDBK,
		"$DBS":     newDBS,
		"$MDA":     newMDA,
		"$MTA":     newMTA,
		"$MTW":     newMTW,
		"$MWD":     newMWD,
		"$MWV":     newMWV,
		"$OSD":     newOSD,
		"$XDR":     newXDR,
		"$XTE":     newXTE,
		"!ABM":     newABM,
		"!BBM":     newBBM,
		"!TTD":     newTTD,
		"!VDM":     newVDMVDO,
		"!VDO":     newVDMVDO,
	}
}
