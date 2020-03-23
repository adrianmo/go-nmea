package nmea

const (
	// TypeDBT type for DBT sentences
	TypeDBT = "DBT"
)

// DBT - Depth below transducer
// https://gpsd.gitlab.io/gpsd/NMEA.html#_dbt_depth_below_transducer
type DBT struct {
	BaseSentence
	DepthFeet    float64
	DepthMeters  float64
	DepthFathoms float64
}

// newDBT constructor
func newDBT(s BaseSentence) (DBT, error) {
	p := newParser(s)
	p.AssertType(TypeDBT)
	return DBT{
		BaseSentence: s,
		DepthFeet:    p.Float64(0, "depth_feet"),
		DepthMeters:  p.Float64(2, "depth_meters"),
		DepthFathoms: p.Float64(4, "depth_fathoms"),
	}, p.Err()
}
