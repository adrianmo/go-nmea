package nmea

import (
	"fmt"
	"strings"
	"sync"
)

const (
	// SentenceStart is the token to indicate the start of a sentence.
	SentenceStart = "$"

	// SentenceStartEncapsulated is the token to indicate the start of encapsulated data.
	SentenceStartEncapsulated = "!"

	// FieldSep is the token to delimit fields of a sentence.
	FieldSep = ","

	// ChecksumSep is the token to delimit the checksum of a sentence.
	ChecksumSep = "*"
)

var (
	customParsersMu = &sync.Mutex{}
	customParsers   = map[string]ParserFunc{}
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
	Checksum string   // The Checksum
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

// parseSentence parses a raw message into it's fields
func parseSentence(raw string) (BaseSentence, error) {
	raw = strings.TrimSpace(raw)
	tagBlockParts := strings.SplitN(raw, `\`, 3)

	var (
		tagBlock TagBlock
		err      error
	)
	if len(tagBlockParts) == 3 {
		tags := tagBlockParts[1]
		raw = tagBlockParts[2]
		tagBlock, err = parseTagBlock(tags)
		if err != nil {
			return BaseSentence{}, err
		}
	}

	startIndex := strings.IndexAny(raw, SentenceStart+SentenceStartEncapsulated)
	if startIndex != 0 {
		return BaseSentence{}, fmt.Errorf("nmea: sentence does not start with a '$' or '!'")
	}
	sumSepIndex := strings.Index(raw, ChecksumSep)
	if sumSepIndex == -1 {
		return BaseSentence{}, fmt.Errorf("nmea: sentence does not contain checksum separator")
	}
	var (
		fieldsRaw   = raw[startIndex+1 : sumSepIndex]
		fields      = strings.Split(fieldsRaw, FieldSep)
		checksumRaw = strings.ToUpper(raw[sumSepIndex+1:])
		checksum    = Checksum(fieldsRaw)
	)
	// Validate the checksum
	if checksum != checksumRaw {
		return BaseSentence{}, fmt.Errorf(
			"nmea: sentence checksum mismatch [%s != %s]", checksum, checksumRaw)
	}
	talker, typ := parsePrefix(fields[0])
	return BaseSentence{
		Talker:   talker,
		Type:     typ,
		Fields:   fields[1:],
		Checksum: checksumRaw,
		Raw:      raw,
		TagBlock: tagBlock,
	}, nil
}

// parsePrefix takes the first field and splits it into a talker id and data type.
func parsePrefix(s string) (string, string) {
	if strings.HasPrefix(s, "PMTK") {
		return "PMTK", s[4:]
	}
	if strings.HasPrefix(s, "P") {
		return "P", s[1:]
	}
	if len(s) < 2 {
		return s, ""
	}
	return s[:2], s[2:]
}

// Checksum xor all the bytes in a string an return it
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
	customParsersMu.Lock()
	defer customParsersMu.Unlock()

	if _, ok := customParsers[sentenceType]; ok {
		return fmt.Errorf("nmea: parser for sentence type '%q' already exists", sentenceType)
	}

	customParsers[sentenceType] = parser
	return nil
}

// Parse parses the given string into the correct sentence type.
func Parse(raw string) (Sentence, error) {
	s, err := parseSentence(raw)
	if err != nil {
		return nil, err
	}

	// Custom parser allow overriding of existing parsers
	if parser, ok := customParsers[s.Type]; ok {
		return parser(s)
	}

	if strings.HasPrefix(s.Raw, SentenceStart) {
		// MTK message types share the same format
		// so we return the same struct for all types.
		switch s.Talker {
		case TypeMTK:
			return newMTK(s)
		}

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
		case TypePRDID:
			return newPRDID(s)
		case TypePSKPDPT:
			return newPSKPDPT(s)
		case TypePSONCMS:
			return newPSONCMS(s)
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
	if strings.HasPrefix(s.Raw, SentenceStartEncapsulated) {
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
