package nmea

import (
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

var defaultSentenceParser = SentenceParser{
	parserLock:    &sync.RWMutex{},
	customParsers: map[string]ParserFunc{},
}

// SentenceParser provides configurable functionality to parse raw input into Sentence
type SentenceParser struct {
	parserLock    *sync.RWMutex
	customParsers map[string]ParserFunc

	// CheckCRC allows custom handling of checksum checking based on parsed sentence
	CheckCRC func(rawCRC string, calculatedCRC string, sentence BaseSentence) error

	// OnTagBlock is callback to handle all parsed tag blocks even for lines containing only a tag block and
	// allows to track multiline tag group separate lines
	//
	// Example of group of 3:
	// \g:1-3-1234,s:r3669961,c:1120959341*hh\
	// \g:2-3-1234*hh\!ABVDM,1,1,1,B,.....,0*hh
	// \g:3-3-1234*hh\$ABVSI,r3669961,1,013536.96326433,1386,-98,,*hh
	OnTagBlock func(tagBlock TagBlock)
}

// NewSentenceParser creates new instance of SentenceParser
func NewSentenceParser() *SentenceParser {
	return &SentenceParser{
		parserLock:    &sync.RWMutex{},
		customParsers: map[string]ParserFunc{},
	}
}

// ParseBaseSentence parses BaseSentence from input
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
		return BaseSentence{}, fmt.Errorf("nmea: sentence does not start with a '$' or '!'")
	}
	checksumSepIndex := strings.Index(raw, ChecksumSep)
	fieldsRaw := raw[startIndex+1:]
	checksumRaw := ""
	if checksumSepIndex != -1 {
		fieldsRaw = raw[startIndex+1 : checksumSepIndex]
		checksumRaw = strings.ToUpper(raw[checksumSepIndex+1:])
	}
	if p.CheckCRC == nil {
		if checksumSepIndex == -1 {
			return BaseSentence{}, fmt.Errorf("nmea: sentence does not contain checksum separator")
		}
		if checksum := Checksum(fieldsRaw); checksum != checksumRaw {
			return BaseSentence{}, fmt.Errorf("nmea: sentence checksum mismatch [%s != %s]", checksum, checksumRaw)
		}
	}

	fields := strings.Split(fieldsRaw, FieldSep)
	talker, typ, err := parsePrefix(fields[0])
	if err != nil {
		return BaseSentence{}, err
	}
	sentence := BaseSentence{
		Talker:   talker,
		Type:     typ,
		Fields:   fields[1:],
		Checksum: checksumRaw,
		Raw:      raw,
		TagBlock: tagBlock,
	}
	if p.CheckCRC != nil {
		if err := p.CheckCRC(checksumRaw, Checksum(fieldsRaw), sentence); err != nil {
			return BaseSentence{}, err
		}
	}

	return sentence, nil
}

// parsePrefix takes the first field and splits it into a talker id and data type.
func parsePrefix(s string) (string, string, error) {
	// proprietary sentences
	if s[0] == ProprietarySentencePrefix {
		return string(ProprietarySentencePrefix), s[1:], nil
	}
	// valid NMEA talkerID (2) + sentence identifier (3+) must be at least 5 characters long
	if len(s) < 5 {
		return "", "", fmt.Errorf("nmea: sentence prefix too short: '%s'", s)
	}
	// query sentence is a special type of sentence in NMEA standard that is used for a listener to request a
	// particular sentence from a talker.
	if len(s) == 5 && s[4] == QuerySentencePostfix {
		return s[:2], string(QuerySentencePostfix), nil
	}
	return s[:2], s[2:], nil
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
	defaultSentenceParser.MustRegisterParser(sentenceType, parser)
}

// MustRegisterParser register a custom parser or panic
func (p *SentenceParser) MustRegisterParser(sentenceType string, parser ParserFunc) {
	if err := p.RegisterParser(sentenceType, parser); err != nil {
		panic(err)
	}
}

// RegisterParser register a custom parser
func RegisterParser(sentenceType string, parser ParserFunc) error {
	return defaultSentenceParser.RegisterParser(sentenceType, parser)
}

// RegisterParser register a custom parser
func (p *SentenceParser) RegisterParser(sentenceType string, parser ParserFunc) error {
	p.parserLock.Lock()
	defer p.parserLock.Unlock()

	if _, ok := p.customParsers[sentenceType]; ok {
		return fmt.Errorf("nmea: parser for sentence type '%q' already exists", sentenceType)
	}

	p.customParsers[sentenceType] = parser
	return nil
}

// Parse parses the given string into the correct sentence type.
func Parse(raw string) (Sentence, error) {
	return defaultSentenceParser.Parse(raw)
}

// Parse parses the given string into the correct sentence type.
func (p *SentenceParser) Parse(raw string) (Sentence, error) {
	s, err := p.ParseBaseSentence(raw)
	if err != nil {
		return nil, err
	}

	// Custom parser allow overriding of existing parsers
	p.parserLock.RLock()
	if parser, ok := p.customParsers[s.Type]; ok {
		p.parserLock.RUnlock()
		return parser(s)
	}
	p.parserLock.RUnlock()

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
	TypeAAM:     newAAM,
	TypeALA:     newALA,
	TypeAPB:     newAPB,
	TypeBEC:     newBEC,
	TypeBOD:     newBOD,
	TypeBWC:     newBWC,
	TypeBWR:     newBWR,
	TypeBWW:     newBWW,
	TypeDBK:     newDBK,
	TypeDBS:     newDBS,
	TypeDBT:     newDBT,
	TypeDOR:     newDOR,
	TypeDPT:     newDPT,
	TypeDSC:     newDSC,
	TypeDSE:     newDSE,
	TypeDTM:     newDTM,
	TypeEVE:     newEVE,
	TypeFIR:     newFIR,
	TypeGGA:     newGGA,
	TypeGLL:     newGLL,
	TypeGNS:     newGNS,
	TypeGSA:     newGSA,
	TypeGSV:     newGSV,
	TypeHDG:     newHDG,
	TypeHDM:     newHDM,
	TypeHDT:     newHDT,
	TypeHSC:     newHSC,
	TypeMDA:     newMDA,
	TypeMTA:     newMTA,
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
	TypeTLL:     newTLL,
	TypeTTM:     newTTM,
	TypeTXT:     newTXT,
	TypeVBW:     newVBW,
	TypeVDR:     newVDR,
	TypeVHW:     newVHW,
	TypeVLW:     newVLW,
	TypeVPW:     newVPW,
	TypeVTG:     newVTG,
	TypeVWR:     newVWR,
	TypeVWT:     newVWT,
	TypeWPL:     newWPL,
	TypeXDR:     newXDR,
	TypeXTE:     newXTE,
	TypeZDA:     newZDA,

	// ones that have `!` as a prefix (encapsulation sentences)
	TypeVDM: newVDMVDO,
	TypeVDO: newVDMVDO,

	// Query is special case as sentence type ends always with `Q` and "destination"/recipient is encoded into
	// sentence type/identifier
	TypeQuery: newQuery,
}
