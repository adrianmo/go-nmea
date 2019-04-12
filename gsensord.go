package nmea

const (
	// TypeGSensord type for GSensord sentences
	TypeGSensord = "GSENSORD"
)

// GSensord represents measured g-loadings in the x, y and z axis.
// http://aprs.gids.nl/nmea/#gsa
type GSensord struct {
	BaseSentence
	X float64 // X-axis G value
	Y float64 // Y-axis G value
	Z float64 // Z-axis G valye
}

// newGSensord parses the GSensord sentence into this struct.
func newGSensord(s BaseSentence) (GSensord, error) {
	p := newParser(s)
	p.AssertType(TypeGSensord)
	m := GSensord{
		BaseSentence: s,
		X:            p.Float64(0, "x-axis g value"),
		Y:            p.Float64(1, "y-axis g value"),
		Z:            p.Float64(2, "z-axis g value"),
	}
	return m, p.Err()
}
