package nmea

const (
	// TypeVDR type of VDR sentence for Set and Drift
	TypeVDR = "VDR"
)

// VDR - Set and Drift
// In navigation, set and drift are characteristics of the current and velocity of water over the ground in which a ship
// is sailing. Set is the bearing the current is flowing. Drift is the magnitude of the current.
// https://gpsd.gitlab.io/gpsd/NMEA.html#_vdr_set_and_drift
//
// Format: $--VDR,x.x,T,x.x,M,x.x,N*hh<CR><LF>
// Example: $IIVDR,10.1,T,12.3,M,1.2,N*3A
type VDR struct {
	BaseSentence
	SetDegreesTrue         float64 // Direction degrees, True
	SetDegreesTrueUnit     string  // T = True
	SetDegreesMagnetic     float64 // Direction degrees, True
	SetDegreesMagneticUnit string  // M = Magnetic
	DriftKnots             float64 // Current speed, knots
	DriftUnit              string  // N = Knots
}

// newVDR constructor
func newVDR(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeVDR)
	return VDR{
		BaseSentence:           s,
		SetDegreesTrue:         p.Float64(0, "true set degrees"),
		SetDegreesTrueUnit:     p.EnumString(1, "true set unit", BearingTrue),
		SetDegreesMagnetic:     p.Float64(2, "magnetic set degrees"),
		SetDegreesMagneticUnit: p.EnumString(3, "magnetic set unit", BearingMagnetic),
		DriftKnots:             p.Float64(4, "drift knots"),
		DriftUnit:              p.EnumString(5, "drift unit", SpeedKnots),
	}, p.Err()
}
