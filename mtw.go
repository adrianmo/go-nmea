package nmea

const (
	// TypeMTW type of MWT sentence describing mean temperature of water
	TypeMTW = "MTW"
	// CelsiusMTW is MTW unit of measurement in celsius
	CelsiusMTW = "C"
)

// MTW is sentence for mean temperature of water.
// https://gpsd.gitlab.io/gpsd/NMEA.html#_mtw_mean_temperature_of_water
//
// Format: $--MTW,TT.T,C*hh<CR><LF>
// Example: $INMTW,17.9,C*1B
type MTW struct {
	BaseSentence
	Temperature  float64 // Temperature, degrees
	CelsiusValid bool    // Is unit of measurement Celsius
}

// newMTW constructor
func newMTW(s BaseSentence) (MTW, error) {
	p := NewParser(s)
	p.AssertType(TypeMTW)
	return MTW{
		BaseSentence: s,
		Temperature:  p.Float64(0, "temperature"),
		CelsiusValid: p.EnumString(1, "unit of measurement celsius", CelsiusMTW) == CelsiusMTW,
	}, p.Err()
}
