package nmea

const (
	// TypeGLL type for GLL sentences
	TypeGLL = "GLL"
	// ValidGLL character
	ValidGLL = "A"
	// InvalidGLL character
	InvalidGLL = "V"
)

// GLL is Geographic Position, Latitude / Longitude and time.
// http://aprs.gids.nl/nmea/#gll
// https://gpsd.gitlab.io/gpsd/NMEA.html#_gll_geographic_position_latitudelongitude
//
// Format            : $--GLL,ddmm.mm,a,dddmm.mm,a,hhmmss.ss,a*hh<CR><LF>
// Format (NMEA 2.3+): $--GLL,ddmm.mm,a,dddmm.mm,a,hhmmss.ss,a,m*hh<CR><LF>
// Example: $IIGLL,5924.462,N,01030.048,E,062216,A*38
// Example: $GNGLL,4404.14012,N,12118.85993,W,001037.00,A,A*67
type GLL struct {
	BaseSentence
	Latitude  float64 // Latitude
	Longitude float64 // Longitude
	Time      Time    // Time Stamp
	Validity  string  // validity - A=valid, V=invalid
	FFAMode   string  // FAA mode indicator (filled in NMEA 2.3 and later)
}

// newGLL constructor
func newGLL(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeGLL)
	gll := GLL{
		BaseSentence: s,
		Latitude:     p.LatLong(0, 1, "latitude"),
		Longitude:    p.LatLong(2, 3, "longitude"),
		Time:         p.Time(4, "time"),
		Validity:     p.EnumString(5, "validity", ValidGLL, InvalidGLL),
	}
	if len(p.Fields) > 6 {
		gll.FFAMode = p.String(6, "FAA mode")
	}
	return gll, p.Err()
}
