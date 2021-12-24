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
// https://gpsd.gitlab.io/gpsd/NMEA.html#_gsa_gps_dop_and_active_satellites
//
// Format:             $--GSA,a,a,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x.x,x.x,x.x*hh<CR><LF>
// Format (NMEA 4.1+): $--GSA,a,a,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x.x,x.x,x.x,x*hh<CR><LF>
// Example: $GNGSA,A,3,80,71,73,79,69,,,,,,,,1.83,1.09,1.47*17
// Example (NMEA 4.1+): $GNGSA,A,3,13,12,22,19,08,21,,,,,,,1.05,0.64,0.83,4*0B
type GSA struct {
	BaseSentence
	Mode    string   // The selection mode.
	FixType string   // The fix type.
	SV      []string // List of satellite PRNs used for this fix.
	PDOP    float64  // Dilution of precision.
	HDOP    float64  // Horizontal dilution of precision.
	VDOP    float64  // Vertical dilution of precision.
	// SystemID is (GNSS) System ID (NMEA 4.1+)
	// 1 - GPS
	// 2 - GLONASS
	// 3 - Galileo
	// 4 - BeiDou
	// 5 - QZSS
	// 6 - NavID (IRNSS)
	SystemID int64
}

// newGSA parses the GSA sentence into this struct.
func newGSA(s BaseSentence) (GSA, error) {
	p := NewParser(s)
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

	if len(p.Fields) > 17 {
		m.SystemID = p.Int64(17, "system ID")
	}
	return m, p.Err()
}
