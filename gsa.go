package nmea

const (
	// TypeGSA type for GSA sentences
	TypeGSA = "GSA"
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

// GSA represents overview satellite data.
// http://aprs.gids.nl/nmea/#gsa
type GSA struct {
	BaseSentence
	Mode    string   // The selection mode.
	FixType string   // The fix type.
	SV      []string // List of satellite PRNs used for this fix.
	PDOP    float64  // Dilution of precision.
	HDOP    float64  // Horizontal dilution of precision.
	VDOP    float64  // Vertical dilution of precision.
}

// newGSA parses the GSA sentence into this struct.
func newGSA(s BaseSentence) (GSA, error) {
	p := newParser(s)
	p.AssertType(TypeGSA)
	m := GSA{
		BaseSentence: s,
		Mode:         p.EnumString(0, "selection mode", Auto, Manual),
		FixType:      p.EnumString(1, "fix type", FixNone, Fix2D, Fix3D),
	}
	// Satellites in view.
	for i := 2; i < 14; i++ {
		if v := p.String(i, "satellite in view"); v != "" {
			m.SV = append(m.SV, v)
		}
	}
	// Dilution of precision.
	m.PDOP = p.Float64(14, "pdop")
	m.HDOP = p.Float64(15, "hdop")
	m.VDOP = p.Float64(16, "vdop")
	return m, p.Err()
}
