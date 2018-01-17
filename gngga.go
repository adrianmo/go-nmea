package nmea

import "fmt"

const (
	// PrefixGNGGA prefix
	PrefixGNGGA = "GNGGA"
)

type GNGGA struct {
	Sentence
	// Time of fix.
	Time Time
	// Latitude.
	Latitude LatLong
	// Longitude.
	Longitude LatLong
	// Quality of fix.
	FixQuality string
	// Number of satellites in use.
	NumSatellites string
	// Horizontal dilution of precision.
	HDOP string
	// Altitude.
	Altitude string
	// Geoidal separation
	Separation string
	// Age of differential GPD data.
	DGPSAge string
	// DGPS reference station ID.
	DGPSId string
}

func NewGNGGA(sentence Sentence) GNGGA {
	s := new(GNGGA)
	s.Sentence = sentence
	return *s
}

func (s GNGGA) GetSentence() Sentence {
	return s.Sentence
}

func (s *GNGGA) parse() error {
	var err error

	if s.Type != PrefixGNGGA {
		return fmt.Errorf("%s is not a %s", s.Type, PrefixGNGGA)
	}
	s.Time, err = ParseTime(s.Fields[0])
	if err != nil {
		return fmt.Errorf("GNGGA decode error: %s", err)
	}
	s.Latitude, err = NewLatLong(fmt.Sprintf("%s %s", s.Fields[1], s.Fields[2]))
	if err != nil {
		return fmt.Errorf("GNGGA decode error: %s", err)
	}
	s.Longitude, err = NewLatLong(fmt.Sprintf("%s %s", s.Fields[3], s.Fields[4]))
	if err != nil {
		return fmt.Errorf("GNGGA decode error: %s", err)
	}
	s.FixQuality = s.Fields[5]
	if s.FixQuality != Invalid && s.FixQuality != GPS && s.FixQuality != DGPS {
		return fmt.Errorf("Invalid fix quality [%s]", s.FixQuality)
	}
	s.NumSatellites = s.Fields[6]
	s.HDOP = s.Fields[7]
	s.Altitude = s.Fields[8]
	s.Separation = s.Fields[10]
	s.DGPSAge = s.Fields[12]
	s.DGPSId = s.Fields[13]
	return nil
}
