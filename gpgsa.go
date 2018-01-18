package nmea

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
func NewGPGSA(sentence Sentence) (GPGSA, error) {
	s := new(GPGSA)
	s.Sentence = sentence
	return *s, s.parse()
}

// GetSentence getter
func (s GPGSA) GetSentence() Sentence {
	return s.Sentence
}

// Parse parses the GPGSA sentence into this struct.
func (s *GPGSA) parse() error {

	p := newParser(s.Sentence, PrefixGPGSA)

	s.Mode = p.String(0, "selection mode")
	if s.Mode != Auto && s.Mode != Manual {
		p.SetErr("selection mode", s.Mode)
	}
	s.FixType = p.String(1, "fix type")
	if s.FixType != FixNone && s.FixType != Fix2D && s.FixType != Fix3D {
		p.SetErr("fix type", s.FixType)
	}
	// Satellites in view.
	for i := 2; i < 14; i++ {
		if v := p.String(i, "satelite in view"); v != "" {
			s.SV = append(s.SV, v)
		}
	}
	// Dilution of precision.
	s.PDOP = p.String(14, "pdop")
	s.HDOP = p.String(15, "hdop")
	s.VDOP = p.String(16, "vdop")

	return p.Err()
}
