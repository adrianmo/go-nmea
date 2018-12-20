package nmea

const (
	// TypeVTG type for VTG sentences
	TypeVTG = "VTG"
)

// VTG represents track & speed data.
// http://aprs.gids.nl/nmea/#vtg
type VTG struct {
	BaseSentence
	TrueTrack        float64
	MagneticTrack    float64
	GroundSpeedKnots float64
	GroundSpeedKPH   float64
}

// newVTG parses the VTG sentence into this struct.
// e.g: $GPVTG,360.0,T,348.7,M,000.0,N,000.0,K*43
func newVTG(s BaseSentence) (VTG, error) {
	p := newParser(s)
	p.AssertType(TypeVTG)
	p.AssertTalker("GP")
	return VTG{
		BaseSentence:     s,
		TrueTrack:        p.Float64(0, "true track"),
		MagneticTrack:    p.Float64(2, "magnetic track"),
		GroundSpeedKnots: p.Float64(4, "ground speed (knots)"),
		GroundSpeedKPH:   p.Float64(6, "ground speed (km/h)"),
	}, p.Err()
}
