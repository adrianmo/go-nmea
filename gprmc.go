package nmea

import (
	"fmt"
	"strconv"
)

const (
	// PrefixGPRMC prefix of GPRMC sentence type
	PrefixGPRMC = "GPRMC"
	// ValidRMC character
	ValidRMC = "A"
	// InvalidRMC character
	InvalidRMC = "V"
)

// GPRMC is the Recommended Minimum Specific GNSS data.
// http://aprs.gids.nl/nmea/#rmc
type GPRMC struct {
	Sentence
	Time      Time    // Time Stamp
	Validity  string  // validity - A-ok, V-invalid
	Latitude  LatLong // Latitude
	Longitude LatLong // Longitude
	Speed     float64 // Speed in knots
	Course    float64 // True course
	Date      string  // Date
	Variation float64 // Magnetic variation
}

// NewGPRMC constructor
func NewGPRMC(sentence Sentence) GPRMC {
	s := new(GPRMC)
	s.Sentence = sentence
	return *s
}

// GetSentence getter
func (s GPRMC) GetSentence() Sentence {
	return s.Sentence
}

func (s *GPRMC) parse() error {
	var err error

	if s.Type != PrefixGPRMC {
		return fmt.Errorf("%s is not a %s", s.Type, PrefixGPRMC)
	}
	s.Time = ParseTime(s.Fields[0])
	s.Validity = s.Fields[1]

	if s.Validity != ValidRMC && s.Validity != InvalidRMC {
		return fmt.Errorf("GPRMC decode, invalid validity '%s'", s.Validity)
	}

	s.Latitude, err = NewLatLong(fmt.Sprintf("%s %s", s.Fields[2], s.Fields[3]))
	if err != nil {
		return fmt.Errorf("GPRMC decode latitude error: %s", err)
	}
	s.Longitude, err = NewLatLong(fmt.Sprintf("%s %s", s.Fields[4], s.Fields[5]))
	if err != nil {
		return fmt.Errorf("GPRMC decode longitude error: %s", err)
	}
	if s.Fields[6] != "" {
		s.Speed, err = strconv.ParseFloat(s.Fields[6], 64)
		if err != nil {
			return fmt.Errorf("GPRMC decode speed error: %s", s.Fields[6])
		}
	}
	if s.Fields[7] != "" {
		s.Course, err = strconv.ParseFloat(s.Fields[7], 64)
		if err != nil {
			return fmt.Errorf("GPRMC decode course error: %s", s.Fields[7])
		}
	}
	s.Date = s.Fields[8]

	if s.Fields[9] != "" {
		s.Variation, err = strconv.ParseFloat(s.Fields[9], 64)
		if err != nil {
			return fmt.Errorf("GPRMC decode variation error: %s", s.Fields[9])
		}
		if s.Fields[10] == "W" {
			s.Variation = 0 - s.Variation
		} else if s.Fields[10] != "E" {
			return fmt.Errorf("GPRMC decode variation error for: %s", s.Fields[10])
		}
	}

	return nil
}
