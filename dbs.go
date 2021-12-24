package nmea

const (
	// TypeDBS type for DBS sentences
	TypeDBS = "DBS"
)

// DBS - Depth Below Surface
// https://gpsd.gitlab.io/gpsd/NMEA.html#_dbs_depth_below_surface
//
// Format: $--DBS,x.x,f,x.x,M,x.x,F*hh<CR><LF>
// Example: $23DBS,01.9,f,0.58,M,00.3,F*21
type DBS struct {
	BaseSentence
	DepthFeet    float64
	DepthMeters  float64
	DepthFathoms float64
}

// newDBS constructor
func newDBS(s BaseSentence) (DBS, error) {
	p := NewParser(s)
	p.AssertType(TypeDBS)
	return DBS{
		BaseSentence: s,
		DepthFeet:    p.Float64(0, "depth_feet"),
		DepthMeters:  p.Float64(2, "depth_meters"),
		DepthFathoms: p.Float64(4, "depth_fathoms"),
	}, p.Err()
}
