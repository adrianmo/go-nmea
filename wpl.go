package nmea

const (
	// TypeWPL type for WPL sentences
	TypeWPL = "WPL"
)

// WPL contains information about a waypoint location
// http://aprs.gids.nl/nmea/#wpl
// https://gpsd.gitlab.io/gpsd/NMEA.html#_wpl_waypoint_location
//
// Format: $--WPL,llll.ll,a,yyyyy.yy,a,c--c*hh<CR><LF>
// Example:  $IIWPL,5503.4530,N,01037.2742,E,411*6F
type WPL struct {
	BaseSentence
	Latitude  float64 // Latitude
	Longitude float64 // Longitude
	Ident     string  // Ident of nth waypoint
}

// newWPL constructor
func newWPL(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeWPL)
	return WPL{
		BaseSentence: s,
		Latitude:     p.LatLong(0, 1, "latitude"),
		Longitude:    p.LatLong(2, 3, "longitude"),
		Ident:        p.String(4, "ident of nth waypoint"),
	}, p.Err()
}
