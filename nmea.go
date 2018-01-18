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

// GetSentence getter
func (s Sentence) GetSentence() Sentence {
	return s
}

func (s *Sentence) parse(input string) error {
	s.Raw = input

	// Start the sentence from the $ character
	startPosition := strings.Index(s.Raw, sentenceStart)
	if startPosition != 0 {
		return fmt.Errorf("Sentence does not start with a '$'")
	}

	sentence := s.Raw[startPosition+1:]

	fieldSum := strings.Split(sentence, checksumSep)
	if len(fieldSum) != 2 {
		return fmt.Errorf("Sentence does not contain single checksum separator")
	}

	fields := strings.Split(fieldSum[0], fieldSep)
	s.Type = fields[0]
	s.Fields = fields[1:]
	s.Checksum = strings.ToUpper(fieldSum[1])

	if err := s.sumOk(); err != nil {
		return fmt.Errorf("Sentence checksum mismatch %s", err)
	}

	return nil
}

// SumOk returns whether the calculated checksum matches the message checksum.
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
func Parse(s string) (SentenceI, error) {
	sentence := Sentence{}
	if err := sentence.parse(s); err != nil {
		return nil, err
	}

	switch sentence.Type {
	case PrefixGPRMC:
		gprmc := NewGPRMC(sentence)
		if err := gprmc.parse(); err != nil {
			return nil, err
		}
		return gprmc, nil
	case PrefixGNRMC:
		gnrmc := NewGNRMC(sentence)
		if err := gnrmc.parse(); err != nil {
			return nil, err
		}
		return gnrmc, nil
	case PrefixGPGGA:
		gpgga := NewGPGGA(sentence)
		if err := gpgga.parse(); err != nil {
			return nil, err
		}
		return gpgga, nil
	case PrefixGNGGA:
		gngga := NewGNGGA(sentence)
		if err := gngga.parse(); err != nil {
			return nil, err
		}
		return gngga, nil
	case PrefixGPGSA:
		gpgsa := NewGPGSA(sentence)
		if err := gpgsa.parse(); err != nil {
			return nil, err
		}
		return gpgsa, nil
	case PrefixGPGLL:
		gpgll := NewGPGLL(sentence)
		if err := gpgll.parse(); err != nil {
			return nil, err
		}
		return gpgll, nil
	case PrefixGPVTG:
		gpvtg := NewGPVTG(sentence)
		if err := gpvtg.parse(); err != nil {
			return nil, err
		}
		return gpvtg, nil
	case PrefixGPZDA:
		gpzda := NewGPZDA(sentence)
		if err := gpzda.parse(); err != nil {
			return nil, err
		}
		return gpzda, nil
	case PrefixPGRME:
		pgrme := NewPGRME(sentence)
		if err := pgrme.parse(); err != nil {
			return nil, err
		}
		return pgrme, nil
	case PrefixGPGSV:
		gpgsv := NewGPGSV(sentence)
		if err := gpgsv.parse(); err != nil {
			return nil, err
		}
		return gpgsv, nil
	case PrefixGLGSV:
		glgsv := NewGLGSV(sentence)
		if err := glgsv.parse(); err != nil {
			return nil, err
		}
		return glgsv, nil
	default:
		return nil, fmt.Errorf("Sentence type '%s' not implemented", sentence.Type)
	}
}
