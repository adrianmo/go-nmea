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
	// CustomParsers allows registering additional parsers
	CustomParsers map[string]ParserFunc

	// ParsePrefix takes in the sentence first field (NMEA0183 address) and splits it into a talker id and sentence type
	ParsePrefix func(prefix string) (talkerID string, sentence string, err error)

	// CheckCRC allows custom handling of checksum checking based on parsed sentence
	CheckCRC func(sentence BaseSentence, rawFields string) error

	// OnTagBlock is callback to handle all parsed tag blocks even for lines containing only a tag block and
	// allows to track multiline tag group separate lines. Logic how to combine/assemble multiline tag group
	// should be implemented outside SentenceParser
	//
	// Example of group of 3:
	// \g:1-3-1234,s:r3669961,c:1120959341*hh\
	// \g:2-3-1234*hh\!ABVDM,1,1,1,B,.....,0*hh
	// \g:3-3-1234*hh\$ABVSI,r3669961,1,013536.96326433,1386,-98,,*hh
	OnTagBlock func(tagBlock TagBlock)
}

// ParseBaseSentence parses BaseSentence from input
// This method is not co-routine safe!
func (p *SentenceParser) ParseBaseSentence(raw string) (BaseSentence, error) {
	raw = strings.TrimSpace(raw)

	var (
		tagBlock TagBlock
		err      error
	)

	if raw[0] == TagBlockSep {
		// tag block is always at the start of line. Starts with `\` and ends with `\` and has valid sentence
		// following or <CR><LF>
		//
		// Note: tag block group can span multiple lines but we only parse ones that have sentence
		endOfTagBlock := strings.LastIndexByte(raw, TagBlockSep)
		if endOfTagBlock <= 0 {
			return BaseSentence{}, fmt.Errorf("nmea: sentence tag block is missing '\\' at the end")
		}
		tagBlock, err = parseTagBlock(raw[1:endOfTagBlock])
		if err != nil {
			return BaseSentence{}, err
		}
		if p.OnTagBlock != nil {
			p.OnTagBlock(tagBlock)
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
		talkerID, typ, err = DefaultParsePrefix(fields[0])
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
		err = DefaultCheckCRC(sentence, rawFields)
	} else {
		err = p.CheckCRC(sentence, rawFields)
	}
	if err != nil {
		return BaseSentence{}, err
	}
	return sentence, nil
}

// DefaultCheckCRC is default implementation for checking sentence Checksum
func DefaultCheckCRC(sentence BaseSentence, rawFields string) error {
	if sentence.Checksum == "" {
		return fmt.Errorf("nmea: sentence does not contain checksum separator")
	}
	if checksum := Checksum(rawFields); checksum != sentence.Checksum {
		return fmt.Errorf("nmea: sentence checksum mismatch [%s != %s]", checksum, sentence.Checksum)
	}
	return nil
}

// DefaultParsePrefix takes in the sentence first field (NMEA0183 Address) and splits it into a talker id and sentence type.
func DefaultParsePrefix(prefix string) (string, string, error) {
	// proprietary sentences
	if prefix[0] == ProprietarySentencePrefix {
		return string(ProprietarySentencePrefix), prefix[1:], nil
	}
	// valid NMEA talkerID (2) + sentence identifier (3+) must be at least 5 characters long
	if len(prefix) < 5 {
		return "", "", fmt.Errorf("nmea: sentence prefix too short: '%s'", prefix)
	}
	// query sentence is a special type of sentence in NMEA standard that is used for a listener to request a
	// particular sentence from a talker.
	// Query address consist of: XXYYQ, XX - requester talker ID, YY - requestee/destination talker ID, `Q` - query type
	// Destination talker ID is handled/parsed by newQuery function
	if len(prefix) == 5 && prefix[4] == QuerySentencePostfix {
		return prefix[:2], string(QuerySentencePostfix), nil
	}
	return prefix[:2], prefix[2:], nil
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

var defaultSentenceParserMu = new(sync.Mutex)

// defaultSentenceParser exists for backwards compatibility reasons to allow global Parse/RegisterParser/MustRegisterParser
// to work as they did before SentenceParser was added.
var defaultSentenceParser = SentenceParser{
	CustomParsers: map[string]ParserFunc{
		TypeMTK: newMTK, // for backwards compatibility support MTK. PMTK001 is correct an supported when using SentenceParser instance
	},
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

	if _, ok := defaultSentenceParser.CustomParsers[sentenceType]; ok {
		return fmt.Errorf("nmea: parser for sentence type '%q' already exists", sentenceType)
	}

	defaultSentenceParser.CustomParsers[sentenceType] = parser
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
	s, err := p.ParseBaseSentence(raw)
	if err != nil {
		return nil, err
	}

	// Custom parser allow overriding of existing parsers
	if parser, ok := p.CustomParsers[s.Type]; ok {
		return parser(s)
	}

	if s.Raw[0] == SentenceStart[0] {
		switch s.Type {
		case TypeRMC:
			return newRMC(s)
		case TypeAAM:
			return newAAM(s)
		case TypeACK:
			return newACK(s)
		case TypeACN:
			return newACN(s)
		case TypeALA:
			return newALA(s)
		case TypeALC:
			return newALC(s)
		case TypeALF:
			return newALF(s)
		case TypeALR:
			return newALR(s)
		case TypeAPB:
			return newAPB(s)
		case TypeARC:
			return newARC(s)
		case TypeBEC:
			return newBEC(s)
		case TypeBOD:
			return newBOD(s)
		case TypeBWC:
			return newBWC(s)
		case TypeBWR:
			return newBWR(s)
		case TypeBWW:
			return newBWW(s)
		case TypeDOR:
			return newDOR(s)
		case TypeDSC:
			return newDSC(s)
		case TypeDSE:
			return newDSE(s)
		case TypeDTM:
			return newDTM(s)
		case TypeEVE:
			return newEVE(s)
		case TypeFIR:
			return newFIR(s)
		case TypeGGA:
			return newGGA(s)
		case TypeGSA:
			return newGSA(s)
		case TypeGLL:
			return newGLL(s)
		case TypeVTG:
			return newVTG(s)
		case TypeZDA:
			return newZDA(s)
		case TypePGRME:
			return newPGRME(s)
		case TypePHTRO:
			return newPHTRO(s)
		case TypePMTK001:
			return newPMTK001(s)
		case TypePRDID:
			return newPRDID(s)
		case TypePSKPDPT:
			return newPSKPDPT(s)
		case TypePSONCMS:
			return newPSONCMS(s)
		case TypeQuery:
			return newQuery(s)
		case TypeGSV:
			return newGSV(s)
		case TypeHBT:
			return newHBT(s)
		case TypeHDG:
			return newHDG(s)
		case TypeHDT:
			return newHDT(s)
		case TypeHDM:
			return newHDM(s)
		case TypeHSC:
			return newHSC(s)
		case TypeGNS:
			return newGNS(s)
		case TypeTHS:
			return newTHS(s)
		case TypeTLB:
			return newTLB(s)
		case TypeTLL:
			return newTLL(s)
		case TypeTTM:
			return newTTM(s)
		case TypeTXT:
			return newTXT(s)
		case TypeWPL:
			return newWPL(s)
		case TypeRMB:
			return newRMB(s)
		case TypeRPM:
			return newRPM(s)
		case TypeRSA:
			return newRSA(s)
		case TypeRSD:
			return newRSD(s)
		case TypeRTE:
			return newRTE(s)
		case TypeROT:
			return newROT(s)
		case TypeVBW:
			return newVBW(s)
		case TypeVDR:
			return newVDR(s)
		case TypeVHW:
			return newVHW(s)
		case TypeVSD:
			return newVSD(s)
		case TypeVPW:
			return newVPW(s)
		case TypeVLW:
			return newVLW(s)
		case TypeVWR:
			return newVWR(s)
		case TypeVWT:
			return newVWT(s)
		case TypeDPT:
			return newDPT(s)
		case TypeDBT:
			return newDBT(s)
		case TypeDBK:
			return newDBK(s)
		case TypeDBS:
			return newDBS(s)
		case TypeMDA:
			return newMDA(s)
		case TypeMTA:
			return newMTA(s)
		case TypeMTW:
			return newMTW(s)
		case TypeMWD:
			return newMWD(s)
		case TypeMWV:
			return newMWV(s)
		case TypeOSD:
			return newOSD(s)
		case TypeXDR:
			return newXDR(s)
		case TypeXTE:
			return newXTE(s)
		}
	}
	if s.Raw[0] == SentenceStartEncapsulated[0] {
		switch s.Type {
		case TypeABM:
			return newABM(s)
		case TypeBBM:
			return newBBM(s)
		case TypeTTD:
			return newTTD(s)
		case TypeVDM, TypeVDO:
			return newVDMVDO(s)
		}
	}
	return nil, &NotSupportedError{Prefix: s.Prefix()}
}

// SupportedParsers list all available parsers provided by the library
var SupportedParsers = map[string]ParserFunc{
	// ones that have `$` as a prefix (parametric sentences)
	TypeAAM: newAAM,
	TypeACK: newACK,
	TypeACN: newACN,
	TypeALA: newALA,
	TypeALC: newALC,
	TypeALF: newALF,
	TypeALR: newALR,
	TypeAPB: newAPB,
	TypeARC: newARC,
	TypeBEC: newBEC,
	TypeBOD: newBOD,
	TypeBWC: newBWC,
	TypeBWR: newBWR,
	TypeBWW: newBWW,
	TypeDBK: newDBK,
	TypeDBS: newDBS,
	TypeDBT: newDBT,
	TypeDOR: newDOR,
	TypeDPT: newDPT,
	TypeDSC: newDSC,
	TypeDSE: newDSE,
	TypeDTM: newDTM,
	TypeEVE: newEVE,
	TypeFIR: newFIR,
	TypeGGA: newGGA,
	TypeGLL: newGLL,
	TypeGNS: newGNS,
	TypeGSA: newGSA,
	TypeGSV: newGSV,
	TypeHBT: newHBT,
	TypeHDG: newHDG,
	TypeHDM: newHDM,
	TypeHDT: newHDT,
	TypeHSC: newHSC,
	TypeMDA: newMDA,
	TypeMTA: newMTA,
	//TypeMTK: newMTK, // deprecated - we are not exposing it here
	TypeMTW:     newMTW,
	TypeMWD:     newMWD,
	TypeMWV:     newMWV,
	TypeOSD:     newOSD,
	TypePGRME:   newPGRME,
	TypePHTRO:   newPHTRO,
	TypePMTK001: newPMTK001,
	TypePRDID:   newPRDID,
	TypePSKPDPT: newPSKPDPT,
	TypePSONCMS: newPSONCMS,
	TypeRMB:     newRMB,
	TypeRMC:     newRMC,
	TypeROT:     newROT,
	TypeRPM:     newRPM,
	TypeRSA:     newRSA,
	TypeRSD:     newRSD,
	TypeRTE:     newRTE,
	TypeTHS:     newTHS,
	TypeTLB:     newTLB,
	TypeTLL:     newTLL,
	TypeTTM:     newTTM,
	TypeTXT:     newTXT,
	TypeVBW:     newVBW,
	TypeVDR:     newVDR,
	TypeVHW:     newVHW,
	TypeVLW:     newVLW,
	TypeVPW:     newVPW,
	TypeVSD:     newVSD,
	TypeVTG:     newVTG,
	TypeVWR:     newVWR,
	TypeVWT:     newVWT,
	TypeWPL:     newWPL,
	TypeXDR:     newXDR,
	TypeXTE:     newXTE,
	TypeZDA:     newZDA,

	// ones that have `!` as a prefix (encapsulation sentences)
	TypeABM: newABM,
	TypeBBM: newBBM,
	TypeTTD: newTTD,
	TypeVDM: newVDMVDO,
	TypeVDO: newVDMVDO,

	// Query is special case as sentence type ends always with `Q` and "destination"/recipient is encoded into
	// sentence type/identifier
	TypeQuery: newQuery,
}
