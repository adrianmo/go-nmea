package nmea

import (
	"fmt"
	"strconv"
)

const (
	// PrefixGPVTG prefix
	PrefixGPVTG = "GPVTG"
)

// GPVTG represents track & speed data.
// http://aprs.gids.nl/nmea/#vtg
type GPVTG struct {
	Sentence
	TrueTrack float64
	MagneticTrack float64
	GroundSpeedKnots float64
	GroundSpeedKPH float64
}

// NewGPVTG constructor
func NewGPVTG(sentence Sentence) GPVTG {
	s := new(GPVTG)
	s.Sentence = sentence
	return *s
}

// GetSentence getter
func (s GPVTG) GetSentence() Sentence {
	return s.Sentence
}

// Parse parses the GPVTG sentence into this struct.
// e.g: $GPVTG,360.0,T,348.7,M,000.0,N,000.0,K*43
func (s *GPVTG) parse() error {
	var err error

	if s.Type != PrefixGPVTG {
		return fmt.Errorf("%s is not a %s", s.Type, PrefixGPVTG)
	}

	if s.Fields[0] != "" {
		s.TrueTrack, err = strconv.ParseFloat(s.Fields[0], 64)
		if err != nil {
			return fmt.Errorf("GPVTG decode true track error: %s", s.Fields[0])
		}
	}

	if s.Fields[2] != "" {
		s.MagneticTrack, err = strconv.ParseFloat(s.Fields[2], 64)
		if err != nil {
			return fmt.Errorf("GPVTG decode magnetic track error: %s", s.Fields[2])
		}
	}

	if s.Fields[4] != "" {
		s.GroundSpeedKnots, err = strconv.ParseFloat(s.Fields[4], 64)
		if err != nil {
			return fmt.Errorf("GPVTG decode ground speed (knots) error: %s", s.Fields[4])
		}
	}

	if s.Fields[6] != "" {
		s.GroundSpeedKPH, err = strconv.ParseFloat(s.Fields[6], 64)
		if err != nil {
			return fmt.Errorf("GPVTG decode ground speed (km/h) error: %s", s.Fields[6])
		}
	}

	return nil
}
