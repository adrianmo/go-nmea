package nmea

const (
	// TypeDBK type of DBK sentence for Depth Below Keel
	TypeDBK = "DBK"
)

// DBK - Depth Below Keel (obsolete, use DPT instead)
// https://gpsd.gitlab.io/gpsd/NMEA.html#_dbk_depth_below_keel
// https://wiki.openseamap.org/wiki/OpenSeaMap-dev:NMEA#DBK_-_Depth_below_keel
//
// Format: $--DBK,x.x,f,x.x,M,x.x,F*hh<CR><LF>
// Example: $SDDBK,12.3,f,3.7,M,2.0,F*2F
type DBK struct {
	BaseSentence
	DepthFeet        float64 // Depth, feet
	DepthFeetUnit    string  // f = feet
	DepthMeters      float64 // Depth, meters
	DepthMetersUnit  string  // M = meters
	DepthFathoms     float64 // Depth, Fathoms
	DepthFathomsUnit string  // F = Fathoms
}

// newDBK constructor
func newDBK(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeDBK)
	return DBK{
		BaseSentence:     s,
		DepthFeet:        p.Float64(0, "depth feet"),
		DepthFeetUnit:    p.EnumString(1, "depth feet unit", DistanceUnitFeet),
		DepthMeters:      p.Float64(2, "depth meters"),
		DepthMetersUnit:  p.EnumString(3, "depth meters unit", DistanceUnitMetre),
		DepthFathoms:     p.Float64(4, "depth fathom"),
		DepthFathomsUnit: p.EnumString(5, "depth fathom unit", DistanceUnitFathom),
	}, p.Err()
}
