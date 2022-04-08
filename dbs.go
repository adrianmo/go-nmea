package nmea

const (
	// TypeDBS is type of DBS sentence for Depth Below Surface
	TypeDBS = "DBS"
)

// DBS - Depth Below Surface (obsolete, use DPT instead)
// https://gpsd.gitlab.io/gpsd/NMEA.html#_dbs_depth_below_surface
// https://wiki.openseamap.org/wiki/OpenSeaMap-dev:NMEA#DBS_-_Depth_below_surface
//
// Format: $--DBS,x.x,f,x.x,M,x.x,F*hh<CR><LF>
// Example: $23DBS,01.9,f,0.58,M,00.3,F*21
type DBS struct {
	BaseSentence
	DepthFeet       float64 // Depth, feet
	DepthFeetUnit   string  // f = feet
	DepthMeters     float64 // Depth, meters
	DepthMeterUnit  string  // M = meters
	DepthFathoms    float64 // Depth, Fathoms
	DepthFathomUnit string  // F = Fathoms
}

// newDBS constructor
func newDBS(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeDBS)
	return DBS{
		BaseSentence:    s,
		DepthFeet:       p.Float64(0, "depth feet"),
		DepthFeetUnit:   p.EnumString(1, "depth feet unit", DistanceUnitFeet),
		DepthMeters:     p.Float64(2, "depth meters"),
		DepthMeterUnit:  p.EnumString(3, "depth feet unit", DistanceUnitMetre),
		DepthFathoms:    p.Float64(4, "depth fathoms"),
		DepthFathomUnit: p.EnumString(5, "depth fathom unit", DistanceUnitFathom),
	}, p.Err()
}
