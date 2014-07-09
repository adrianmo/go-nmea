package nmea

import "fmt"

const (
	PrefixGPVTG = "GPVTG"
)

// GPVTG is the Track Made Good and Ground Speed.
// http://aprs.gids.nl/nmea/#vtg
type GPVTG struct {
	Sentence
	// True track
	TrueTrack string
	// Magnetic track
	MagneticTrack string
	// Ground speed, knots.
	SpeedKnots string
	// Ground speed kph.
	SpeedKPH string
}

// Parse parses the GPVTG into this struct.
// e.g: $GPVTG,356.10,T,,M,0.55,N,1.0,K,A*0D
func (g *GPVTG) Parse(s string) error {
	err := g.Sentence.Parse(s)
	if err != nil {
		// Sentence decode error.
		return err
	}
	if g.SType != PrefixGPVTG {
		return fmt.Errorf("%s is not a %s", g.SType, PrefixGPVTG)
	}
	g.TrueTrack = g.Fields[0]
	if g.Fields[1] != "T" {
		return fmt.Errorf("field expected 'T' got '%s'", g.Fields[1])
	}
	g.MagneticTrack = g.Fields[2]
	if g.Fields[3] != "M" {
		return fmt.Errorf("field expected 'M' got '%s'", g.Fields[3])
	}
	g.SpeedKnots = g.Fields[4]
	if g.Fields[5] != "N" {
		return fmt.Errorf("field expected 'N' got '%s'", g.Fields[5])
	}
	g.SpeedKPH = g.Fields[6]
	if g.Fields[7] != "K" {
		return fmt.Errorf("field expected 'K' got '%s'", g.Fields[7])
	}
	return nil
}
