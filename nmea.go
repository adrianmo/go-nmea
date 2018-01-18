package nmea

import (
	"fmt"
	"strings"
)

const (
	// The token to indicate the start of a sentence.
	sentenceStart = "$"
	// The token to delimit fields of a sentence.
	fieldSep = ","
	// The token to delimit the checksum of a sentence.
	checksumSep = "*"
)

//SentenceI interface for all NMEA sentence
type SentenceI interface {
	GetSentence() Sentence
}

// Sentence contains the information about the NMEA sentence
type Sentence struct {
	Type     string   // The sentence type (e.g $GPGSA)
	Fields   []string // Array of fields
	Checksum string   // Checksum
	Raw      string   // The raw NMEA sentence received
}

// ParseSentence parses a raw message into it's fields
func ParseSentence(raw string) (Sentence, error) {
	var s Sentence

	// Start the sentence from the $ character
	startPosition := strings.Index(raw, sentenceStart)
	if startPosition != 0 {
		return s, fmt.Errorf("Sentence does not start with a '$'")
	}

	sentence := raw[startPosition+1:]

	fieldSum := strings.Split(sentence, checksumSep)
	if len(fieldSum) != 2 {
		return s, fmt.Errorf("Sentence does not contain single checksum separator")
	}

	fields := strings.Split(fieldSum[0], fieldSep)
	s.Type = fields[0]
	s.Fields = fields[1:]
	s.Checksum = strings.ToUpper(fieldSum[1])
	s.Raw = raw

	if err := s.sumOk(); err != nil {
		return s, fmt.Errorf("Sentence checksum mismatch %s", err)
	}
	return s, nil
}

// GetSentence getter
func (s Sentence) GetSentence() Sentence {
	return s
}

// sumOk returns whether the calculated checksum matches the message checksum.
func (s *Sentence) sumOk() error {
	var checksum uint8
	for i := 1; i < len(s.Raw) && string(s.Raw[i]) != checksumSep; i++ {
		checksum ^= s.Raw[i]
	}

	calculated := fmt.Sprintf("%X", checksum)
	if len(calculated) == 1 {
		calculated = "0" + calculated
	}
	if calculated != s.Checksum {
		return fmt.Errorf("[%s != %s]", calculated, s.Checksum)
	}
	return nil
}

// Parse parses the given string into the correct sentence type.
func Parse(raw string) (SentenceI, error) {
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
		return nil, fmt.Errorf("Sentence type '%s' not implemented", s.Type)
	}
}
