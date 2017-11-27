package nmea

import (
	"fmt"
	"strconv"
)

const (
	// PrefixGNRMC prefix of GNRMC sentence type
	PrefixGNRMC = "GNRMC"
)

// GNRMC is the Recommended Minimum Specific GNSS data.
// http://aprs.gids.nl/nmea/#rmc
type GNRMC struct {
	Sentence
	Time      string  // Time Stamp
	Validity  string  // validity - A-ok, V-invalid
	Latitude  LatLong // Latitude
	Longitude LatLong // Longitude
	Speed     float64 // Speed in knots
	Course    float64 // True course
	Date      string  // Date
	Variation float64 // Magnetic variation
}

// NewGNRMC constructor
func NewGNRMC(sentence Sentence) GNRMC {
	s := new(GNRMC)
	s.Sentence = sentence
	return *s
}

// GetSentence getter
func (s GNRMC) GetSentence() Sentence {
	return s.Sentence
}

func (s *GNRMC) parse() error {
	var err error

	if s.Type != PrefixGNRMC {
		return fmt.Errorf("%s is not a %s", s.Type, PrefixGNRMC)
	}
	s.Time = s.Fields[0]
	s.Validity = s.Fields[1]

	if s.Validity != ValidRMC && s.Validity != InvalidRMC {
		return fmt.Errorf("GNRMC decode, invalid validity '%s'", s.Validity)
	}

	s.Latitude, err = NewLatLong(fmt.Sprintf("%s %s", s.Fields[2], s.Fields[3]))
	if err != nil {
		return fmt.Errorf("GNRMC decode latitude error: %s", err)
	}
	s.Longitude, err = NewLatLong(fmt.Sprintf("%s %s", s.Fields[4], s.Fields[5]))
	if err != nil {
		return fmt.Errorf("GNRMC decode longitude error: %s", err)
	}
	if s.Fields[6] != "" {
		s.Speed, err = strconv.ParseFloat(s.Fields[6], 64)
		if err != nil {
			return fmt.Errorf("GNRMC decode speed error: %s", s.Fields[6])
		}
	}
	if s.Fields[7] != "" {
		s.Course, err = strconv.ParseFloat(s.Fields[7], 64)
		if err != nil {
			return fmt.Errorf("GNRMC decode course error: %s", s.Fields[7])
		}
	}
	s.Date = s.Fields[8]

	if s.Fields[9] != "" {
		s.Variation, err = strconv.ParseFloat(s.Fields[9], 64)
		if err != nil {
			return fmt.Errorf("GNRMC decode variation error: %s", s.Fields[9])
		}
		if s.Fields[10] == "W" {
			s.Variation = 0 - s.Variation
		} else if s.Fields[10] != "E" {
			return fmt.Errorf("GNRMC decode variation error for: %s", s.Fields[10])
		}
	}

	return nil
}
