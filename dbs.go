package nmea

const (
	// TypeDBS type for DBS sentences
	TypeDBS = "DBS"
)

// DBS - Depth Below Surface
// https://gpsd.gitlab.io/gpsd/NMEA.html#_dbs_depth_below_surface
type DBS struct {
	BaseSentence
	DepthFeet    float64
	DepthMeters  float64
	DepthFathoms float64
}

// newDBS constructor
func newDBS(s BaseSentence) (DBS, error) {
	p := newParser(s)
	p.AssertType(TypeDBS)
	return DBS{
		BaseSentence: s,
		DepthFeet:    p.Float64(0, "depth_feet"),
		DepthMeters:  p.Float64(2, "depth_meters"),
		DepthFathoms: p.Float64(4, "depth_fathoms"),
	}, p.Err()
}
