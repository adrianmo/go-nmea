package nmea

import "fmt"

const (
	// PrefixGPGLL prefix for GPGLL sentence type
	PrefixGPGLL = "GPGLL"
	// ValidGLL character
	ValidGLL = "A"
	// InvalidGLL character
	InvalidGLL = "V"
)

// GPGLL is Geographic Position, Latitude / Longitude and time.
// http://aprs.gids.nl/nmea/#gll
type GPGLL struct {
	Sentence
	Latitude  LatLong // Latitude
	Longitude LatLong // Longitude
	Time      Time    // Time Stamp
	Validity  string  // validity - A-valid
}

// NewGPGLL constructor
func NewGPGLL(sentence Sentence) GPGLL {
	s := new(GPGLL)
	s.Sentence = sentence
	return *s
}

// GetSentence getter
func (s GPGLL) GetSentence() Sentence {
	return s.Sentence
}

func (s *GPGLL) parse() error {
	var err error

	if s.Type != PrefixGPGLL {
		return fmt.Errorf("%s is not a %s", s.Type, PrefixGPGLL)
	}

	s.Latitude, err = NewLatLong(fmt.Sprintf("%s %s", s.Fields[0], s.Fields[1]))
	if err != nil {
		return fmt.Errorf("GPGLL decode latitude error: %s", err)
	}
	s.Longitude, err = NewLatLong(fmt.Sprintf("%s %s", s.Fields[2], s.Fields[3]))
	if err != nil {
		return fmt.Errorf("GPGLL decode longitude error: %s", err)
	}

	s.Time = ParseTime(s.Fields[4])
	s.Validity = s.Fields[5]

	if s.Validity != ValidGLL && s.Validity != InvalidGLL {
		return fmt.Errorf("GPGLL decode, invalid validity '%s'", s.Validity)
	}

	return nil
}
