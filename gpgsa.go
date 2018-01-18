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
	Mode    string   // The selection mode.
	FixType string   // The fix type.
	SV      []string // List of satellite PRNs used for this fix.
	PDOP    string   // Dilution of precision.
	HDOP    string   // Horizontal dilution of precision.
	VDOP    string   // Vertical dilution of precision.
}

// NewGPGSA parses the GPGSA sentence into this struct.
func NewGPGSA(sentence Sentence) (GPGSA, error) {
	s := GPGSA{Sentence: sentence}
	p := newParser(s.Sentence, PrefixGPGSA)

	s.Mode = p.EnumString(0, "selection mode", Auto, Manual)
	s.FixType = p.EnumString(1, "fix type", FixNone, Fix2D, Fix3D)
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

	return s, p.Err()
}

// GetSentence getter
func (s GPGSA) GetSentence() Sentence {
	return s.Sentence
}
