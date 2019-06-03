package nmea

const (
	// TypeWPL type for WPL sentences
	TypeWPL = "WPL"
)

// WPL contains information about a waypoint location
type WPL struct {
	BaseSentence
	Latitude  float64 // Latitude
	Longitude float64 // Longitude
	Ident     string  // Ident of nth waypoint
}

// newWPL constructor
func newWPL(s BaseSentence) (WPL, error) {
	p := newParser(s)
	p.AssertType(TypeWPL)
	return WPL{
		BaseSentence: s,
		Latitude:     p.LatLong(0, 1, "latitude"),
		Longitude:    p.LatLong(2, 3, "longitude"),
		Ident:        p.String(4, "ident of nth waypoint"),
	}, p.Err()
}
