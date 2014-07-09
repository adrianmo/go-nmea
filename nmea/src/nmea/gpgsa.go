package nmea

import "fmt"

const (
	PrefixGPGSA = "GPGSA"
	// Field 1, auto or manual fix.
	Auto   = "A"
	Manual = "M"
	// Field 2, fix type.
	FixNone = "1"
	Fix2D   = "2"
	Fix3D   = "3"
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

// Parse parses the GPGSA sentence into this struct.
func (g *GPGSA) Parse(s string) error {
	err := g.Sentence.Parse(s)
	if err != nil {
		// Sentence decode error.
		return err
	}
	if g.SType != PrefixGPGSA {
		return fmt.Errorf("%s is not a %s", g.SType, PrefixGPGSA)
	}
	// Selection mode.
	g.Mode = g.Fields[0]
	if g.Mode != Auto && g.Mode != Manual {
		return fmt.Errorf("Invalid selection mode [%s]", g.Mode)
	}
	// Fix type.
	g.FixType = g.Fields[1]
	if g.FixType != FixNone && g.FixType != Fix2D &&
		g.FixType != Fix3D {
		return fmt.Errorf("Invalid fix type [%s]", g.FixType)
	}
	// Satellites in view.
	for i := 2; i < 14; i++ {
		if g.Fields[i] != "" {
			g.SV = append(g.SV, g.Fields[i])
		}
	}
	// Dilution of precision.
	g.PDOP = g.Fields[14]
	g.HDOP = g.Fields[15]
	g.VDOP = g.Fields[16]
	return nil
}
