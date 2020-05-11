package nmea

const (
	// TypeDPT type for DPT sentences
	TypeDPT = "DPT"
)

// DPT - Depth of Water
// https://gpsd.gitlab.io/gpsd/NMEA.html#_dpt_depth_of_water
type DPT struct {
	BaseSentence
	Depth      float64
	Offset     float64
	RangeScale float64
}

// newDPT constructor
func newDPT(s BaseSentence) (DPT, error) {
	p := NewParser(s)
	p.AssertType(TypeDPT)
	return DPT{
		BaseSentence: s,
		Depth:        p.Float64(0, "depth"),
		Offset:       p.Float64(1, "offset"),
		RangeScale:   p.Float64(2, "range scale"),
	}, p.Err()
}
