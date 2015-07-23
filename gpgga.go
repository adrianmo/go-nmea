package nmea

import "fmt"

const (
	// PrefixGPGGA prefix
	PrefixGPGGA = "GPGGA"
	// Invalid fix quality.
	Invalid = "0"
	// GPS fix quality
	GPS = "1"
	// DGPS fix quality
	DGPS = "2"
)

// GPGGA represents fix data.
// http://aprs.gids.nl/nmea/#gga
type GPGGA struct {
	Sentence
	// Time of fix.
	Time string
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

// NewGPGGA constructor
func NewGPGGA(sentence Sentence) GPGGA {
	s := new(GPGGA)
	s.Sentence = sentence
	return *s
}

// GetSentence getter
func (s GPGGA) GetSentence() Sentence {
	return s.Sentence
}

// Parse parses the GPGGA sentence into this struct.
// e.g: $GPGGA,034225.077,3356.4650,S,15124.5567,E,1,03,9.7,-25.0,M,21.0,M,,0000*58
func (s *GPGGA) parse() error {
	var err error

	if s.Type != PrefixGPGGA {
		return fmt.Errorf("%s is not a %s", s.Type, PrefixGPGGA)
	}
	s.Time = s.Fields[0]
	s.Latitude, err = NewLatLong(fmt.Sprintf("%s %s", s.Fields[1], s.Fields[2]))
	if err != nil {
		return fmt.Errorf("GPGGA decode error: %s", err)
	}
	s.Longitude, err = NewLatLong(fmt.Sprintf("%s %s", s.Fields[3], s.Fields[4]))
	if err != nil {
		return fmt.Errorf("GPGGA decode error: %s", err)
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
