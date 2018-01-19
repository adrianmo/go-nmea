package nmea

import (
	"fmt"
	"strings"
)

const (
	SentenceStart = "$" // The token to indicate the start of a sentence.
	FieldSep      = "," // The token to delimit fields of a sentence.
	ChecksumSep   = "*" // The token to delimit the checksum of a sentence.
)

//Message interface for all NMEA sentence
type Message interface {
	fmt.Stringer
	GetSentence() Sentence
	GetType() string
}

// Sentence contains the information about the NMEA sentence
type Sentence struct {
	Type     string   // The sentence type (e.g $GPGSA)
	Fields   []string // Array of fields
	Checksum string   // The Checksum
	Raw      string   // The raw NMEA sentence received
}

func (s Sentence) GetSentence() Sentence { return s }
func (s Sentence) GetType() string       { return s.Type }
func (s Sentence) String() string        { return s.Raw }

// ParseSentence parses a raw message into it's fields
func ParseSentence(raw string) (Sentence, error) {
	startIndex := strings.Index(raw, SentenceStart)
	if startIndex != 0 {
		return Sentence{}, fmt.Errorf("nmea: sentence does not start with a '$'")
	}
	sumSepIndex := strings.Index(raw, ChecksumSep)
	if sumSepIndex == -1 {
		return Sentence{}, fmt.Errorf("nmea: sentence does not contain checksum separator")
	}
	var (
		fieldsRaw   = raw[startIndex+1 : sumSepIndex]
		fields      = strings.Split(fieldsRaw, FieldSep)
		checksumRaw = strings.ToUpper(raw[sumSepIndex+1:])
		checksum    = xorChecksum(fieldsRaw)
	)
	// Validate the checksum
	if checksum != checksumRaw {
		return Sentence{}, fmt.Errorf("nmea: sentence checksum mismatch [%s != %s]", checksum, checksumRaw)
	}
	return Sentence{
		Type:     fields[0],
		Fields:   fields[1:],
		Checksum: checksumRaw,
		Raw:      raw,
	}, nil
}

// xor all the bytes in a string an return it
// as an uppercase hex string
func xorChecksum(s string) string {
	var checksum uint8
	for i := 0; i < len(s); i++ {
		checksum ^= s[i]
	}
	return fmt.Sprintf("%02X", checksum)
}

// Parse parses the given string into the correct sentence type.
func Parse(raw string) (Message, error) {
	s, err := ParseSentence(raw)
	if err != nil {
		return nil, err
	}
	switch s.Type {
	case PrefixGPRMC:
		return NewGPRMC(s)
	case PrefixGNRMC:
		return NewGNRMC(s)
	case PrefixGPGGA:
		return NewGPGGA(s)
	case PrefixGNGGA:
		return NewGNGGA(s)
	case PrefixGPGSA:
		return NewGPGSA(s)
	case PrefixGPGLL:
		return NewGPGLL(s)
	case PrefixGPVTG:
		return NewGPVTG(s)
	case PrefixGPZDA:
		return NewGPZDA(s)
	case PrefixPGRME:
		return NewPGRME(s)
	case PrefixGPGSV:
		return NewGPGSV(s)
	case PrefixGLGSV:
		return NewGLGSV(s)
	default:
		return nil, fmt.Errorf("nmea: sentence type '%s' not implemented", s.Type)
	}
}
