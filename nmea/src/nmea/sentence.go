/*
Package nmea implements decoding for NMEA GPS sentences.

The package is able to perform generic decoding of sentences, including
checksum validation. It also has some parsers for specific sentence
types.

You typically only need to call one function for complete parsing:

sentence, err := Parse("$GPVTG,054.7,T,034.4,M,005.5,N,010.2,K*48")

switch s := sentence.(type) {
case GPGSA:
  fmt.Printf("This is a GPGSA message.")
case GPGGA:
  fmt.Printf("This is a GPGGA message.")
}

*/
package nmea

import "fmt"
import "strings"

const (
	// The token to indicate the start of a sentence.
	sentenceStart = "$"
	// The token to delimit fields of a sentence.
	fieldSep = ","
	// The token to delimit the checksum of a sentence.
	checksumSep = "*"
)


// Sentence is an NMEA sentence. It represents the base sentence
// without message specific parsing.
type Sentence struct {
	// The raw string passed in for parsing.
	Raw string
	// The sentence type (e.g $GPGSA)
	SType string
	// The comma separate fields (excluding sentence type/checksum)
	Fields []string
	// The checksum in the sentence.
	Checksum string
}

// Parse parses the given string as an NMEA sentence.
func (s *Sentence) Parse(str string) error {
	s.Raw = str
	if strings.Count(str, checksumSep) != 1 {
		return fmt.Errorf("Sentence does not contain single checksum separator")
	}
	if strings.Index(str, sentenceStart) != 0 {
		return fmt.Errorf("Sentence does not start with a '$'")
	}
	sentence := strings.Split(str, sentenceStart)[1]
	fieldSum := strings.Split(sentence, checksumSep)
	fields := strings.Split(fieldSum[0], fieldSep)
	s.SType = fields[0]
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
	if calculated != s.Checksum {
		return fmt.Errorf("[%s != %s]", calculated, s.Checksum)
	}
	return nil
}

// Parse parses the given string into the correct sentence type.
func Parse(s string) (interface{}, error) {
	sentence := Sentence{}
	if err := sentence.Parse(s); err != nil {
		return nil, err
	}
	if sentence.SType == PrefixGPGGA {
		gpgga := GPGGA{}
		if err := gpgga.Parse(s); err != nil {
			return nil, err
		}
		return gpgga, nil
	}
	if sentence.SType == PrefixGPGSA {
		gpgsa := GPGSA{}
		if err := gpgsa.Parse(s); err != nil {
			return nil, err
		}
		return gpgsa, nil
	}
	if sentence.SType == PrefixGPRMC {
		gprmc := GPRMC{}
		if err := gprmc.Parse(s); err != nil {
			return nil, err
		}
		return gprmc, nil
	}
	if sentence.SType == PrefixGPVTG {
		gpvtg := GPVTG{}
		if err := gpvtg.Parse(s); err != nil {
			return nil, err
		}
		return gpvtg, nil
	}
	if sentence.SType == PrefixPSRFTXT {
		psrf := PSRFTXT{}
		if err := psrf.Parse(s); err != nil {
			return nil, err
		}
		return psrf, nil
	}
	return sentence, nil
}
