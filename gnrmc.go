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
	Time      Time    // Time Stamp
	Validity  string  // validity - A-ok, V-invalid
	Latitude  LatLong // Latitude
	Longitude LatLong // Longitude
	Speed     float64 // Speed in knots
	Course    float64 // True course
	Date      string  // Date
	Variation float64 // Magnetic variation
}

// NewGNRMC constructor
func NewGNRMC(sentence Sentence) (GNRMC, error) {
	s := GNRMC{Sentence: sentence}
	p := newParser(sentence, PrefixGNRMC)

	s.Time = p.Time(0, "time")
	s.Validity = p.String(1, "validity")
	if s.Validity != ValidRMC && s.Validity != InvalidRMC {
		return s, fmt.Errorf("GNRMC invalid validity: %s", s.Validity)
	}

	s.Latitude = p.LatLong(2, 3, "latitude")
	s.Longitude = p.LatLong(4, 5, "longitude")
	s.Speed = p.Float64(6, "speed")
	s.Course = p.Float64(7, "course")
	s.Date = p.String(8, "date")

	if err := p.Err(); err != nil {
		return s, err
	}

	var err error
	if s.Fields[9] != "" {
		s.Variation, err = strconv.ParseFloat(s.Fields[9], 64)
		if err != nil {
			return s, fmt.Errorf("GNRMC invalid variation: %s", s.Fields[9])
		}
		if s.Fields[10] == "W" {
			s.Variation = 0 - s.Variation
		} else if s.Fields[10] != "E" {
			return s, fmt.Errorf("GNRMC invalid variation: %s", s.Fields[10])
		}
	}

	return s, p.Err()
}

// GetSentence getter
func (s GNRMC) GetSentence() Sentence {
	return s.Sentence
}
