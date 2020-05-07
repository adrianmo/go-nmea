package nmea

const (
	// TypeVHW type for VHW sentences
	TypeVHW = "VHW"
)

// VHW contains information about water speed and heading
type VHW struct {
	BaseSentence
	TrueHeading            float64
	MagneticHeading        float64
	SpeedThroughWaterKnots float64
	SpeedThroughWaterKPH   float64
}

// newVHW constructor
func newVHW(s BaseSentence) (VHW, error) {
	p := NewParser(s)
	p.AssertType(TypeVHW)
	return VHW{
		BaseSentence:           s,
		TrueHeading:            p.Float64(0, "true heading"),
		MagneticHeading:        p.Float64(2, "magnetic heading"),
		SpeedThroughWaterKnots: p.Float64(4, "speed through water in knots"),
		SpeedThroughWaterKPH:   p.Float64(6, "speed through water in kilometers per hour"),
	}, p.Err()
}
