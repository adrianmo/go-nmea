package nmea

import (
	"fmt"
	"strconv"
)

const (
	// PrefixPGRME prefix for PGRME sentence type
	PrefixPGRME = "PGRME"
	// ErrorUnit must be meters (M)
	ErrorUnit = "M"
)

// PGRME is Estimated Position Error (Garmin proprietary sentence)
// http://aprs.gids.nl/nmea/#rme
type PGRME struct {
	Sentence
	Horizontal float64 // Estimated horizontal position error (HPE) in metres
	Vertical   float64 // Estimated vertical position error (VPE) in metres
	Spherical  float64 // Overall spherical equivalent position error in meters
}

// NewPGRME constructor
func NewPGRME(sentence Sentence) PGRME {
	s := new(PGRME)
	s.Sentence = sentence
	return *s
}

// GetSentence getter
func (s PGRME) GetSentence() Sentence {
	return s.Sentence
}

func (s *PGRME) parse() error {
	var err error

	if s.Type != PrefixPGRME {
		return fmt.Errorf("%s is not a %s", s.Type, PrefixPGRME)
	}

	s.Horizontal, err = strconv.ParseFloat(s.Fields[0], 64)
	if err != nil {
		return fmt.Errorf("PGRME decode invalid horizontal error: '%s'", s.Fields[0])
	}

	if s.Fields[1] != ErrorUnit {
		return fmt.Errorf("PGRME decode invalid horizontal error unit: '%s'", s.Fields[1])
	}

	s.Vertical, err = strconv.ParseFloat(s.Fields[2], 64)
	if err != nil {
		return fmt.Errorf("PGRME decode invalid vertical error: '%s'", s.Fields[2])
	}

	if s.Fields[3] != ErrorUnit {
		return fmt.Errorf("PGRME decode invalid vertical error unit: '%s'", s.Fields[3])
	}

	s.Spherical, err = strconv.ParseFloat(s.Fields[4], 64)
	if err != nil {
		return fmt.Errorf("PGRME decode invalid spherical error: '%s'", s.Fields[4])
	}

	if s.Fields[5] != ErrorUnit {
		return fmt.Errorf("PGRME decode invalid spherical error unit: '%s'", s.Fields[5])
	}

	return nil
}
