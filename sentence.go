package nmea

import (
	"fmt"
	"strings"
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
	customParsers          = map[string]parserCallbackType{}
	defaultSentenceParsers = map[string]parserCallbackType{
		TypeRMC: func(s BaseSentence) (Sentence, error) {
			return newRMC(s)
		},
		TypeGGA: func(s BaseSentence) (Sentence, error) {
			return newGGA(s)
		},
		TypeGSA: func(s BaseSentence) (Sentence, error) {
			return newGSA(s)
		},
		TypeGLL: func(s BaseSentence) (Sentence, error) {
			return newGLL(s)
		},
		TypeVTG: func(s BaseSentence) (Sentence, error) {
			return newVTG(s)
		},
		TypeZDA: func(s BaseSentence) (Sentence, error) {
			return newZDA(s)
		},
		TypePGRME: func(s BaseSentence) (Sentence, error) {
			return newPGRME(s)
		},
		TypeGSV: func(s BaseSentence) (Sentence, error) {
			return newGSV(s)
		},
		TypeHDT: func(s BaseSentence) (Sentence, error) {
			return newHDT(s)
		},
		TypeGNS: func(s BaseSentence) (Sentence, error) {
			return newGNS(s)
		},
		TypeTHS: func(s BaseSentence) (Sentence, error) {
			return newTHS(s)
		},
		TypeWPL: func(s BaseSentence) (Sentence, error) {
			return newWPL(s)
		},
		TypeRTE: func(s BaseSentence) (Sentence, error) {
			return newRTE(s)
		},
		TypeVHW: func(s BaseSentence) (Sentence, error) {
			return newVHW(s)
		},
	}
)

type parserCallbackType func(BaseSentence) (Sentence, error)

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
func MustRegisterParser(t string, parser func(BaseSentence) (Sentence, error)) {
	if err := RegisterParser(t, parser); err != nil {
		panic(err)
	}
}

// RegisterParser register a custom parser
func RegisterParser(t string, parser func(BaseSentence) (Sentence, error)) error {
	if _, ok := customParsers[t]; ok {
		return fmt.Errorf("nmea: parser for prefix '%s' already exists", t)
	}

	customParsers[t] = parser
	return nil
}

// Parse parses the given string into the correct sentence type.
func Parse(raw string) (Sentence, error) {
	s, err := parseSentence(raw)
	if err != nil {
		return nil, err
	}

	// Custom parser allow overriding of existing parsers
	if customParserCallback, ok := customParsers[s.Type]; ok {
		return customParserCallback(s)
	}

	if strings.HasPrefix(s.Raw, SentenceStart) {
		// MTK message types share the same format
		// so we return the same struct for all types.
		switch s.Talker {
		case TypeMTK:
			return newMTK(s)
		}

		if parserCallback, ok := defaultSentenceParsers[s.Type]; ok {
			return parserCallback(s)
		}
	}
	if strings.HasPrefix(s.Raw, SentenceStartEncapsulated) {
		switch s.Type {
		case TypeVDM, TypeVDO:
			return newVDMVDO(s)
		}
	}
	return nil, fmt.Errorf("nmea: sentence prefix '%s' not supported", s.Prefix())
}
