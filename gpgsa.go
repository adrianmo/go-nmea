package nmea

import "fmt"

const (
	// PrefixGPGSA prefix of GPGSA sentence type
	PrefixGPGSA = "GPGSA"
	// Auto - Field 1, auto or manual fix.
	Auto = "A"
	// Manual - Field 1, auto or manual fix.
	Manual = "M"
	// FixNone - Field 2, fix type.
	FixNone = "1"
	// Fix2D - Field 2, fix type.
	Fix2D = "2"
	// Fix3D - Field 2, fix type.
	Fix3D = "3"
)

// GPGSA represents overview satellite data.
// http://aprs.gids.nl/nmea/#gsa
type GPGSA struct {
	Sentence
	// The selection mode.
	Mode string
	// The fix type.
	FixType string
	// List of satellite PRNs used for this fix.
	SV []string
	// Dilution of precision.
	PDOP string
	// Horizontal dilution of precision.
	HDOP string
	// Vertical dilution of precision.
	VDOP string
}

// NewGPGSA constructor
func NewGPGSA(sentence Sentence) GPGSA {
	s := new(GPGSA)
	s.Sentence = sentence
	return *s
}

// GetSentence getter
func (s GPGSA) GetSentence() Sentence {
	return s.Sentence
}

// Parse parses the GPGSA sentence into this struct.
func (s *GPGSA) parse() error {

	if s.Type != PrefixGPGSA {
		return fmt.Errorf("%s is not a %s", s.Type, PrefixGPGSA)
	}
	// Selection mode.
	s.Mode = s.Fields[0]
	if s.Mode != Auto && s.Mode != Manual {
		return fmt.Errorf("Invalid selection mode [%s]", s.Mode)
	}
	// Fix type.
	s.FixType = s.Fields[1]
	if s.FixType != FixNone && s.FixType != Fix2D &&
		s.FixType != Fix3D {
		return fmt.Errorf("Invalid fix type [%s]", s.FixType)
	}
	// Satellites in view.
	for i := 2; i < 14; i++ {
		if s.Fields[i] != "" {
			s.SV = append(s.SV, s.Fields[i])
		}
	}
	// Dilution of precision.
	s.PDOP = s.Fields[14]
	s.HDOP = s.Fields[15]
	s.VDOP = s.Fields[16]

	return nil
}
