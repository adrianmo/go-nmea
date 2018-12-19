package nmea

const (
	// PrefixGLL prefix for GLL sentence type
	PrefixGLL = "GLL"
	// ValidGLL character
	ValidGLL = "A"
	// InvalidGLL character
	InvalidGLL = "V"
)

// GLL is Geographic Position, Latitude / Longitude and time.
// http://aprs.gids.nl/nmea/#gll
type GLL struct {
	BaseSentence
	Latitude  float64 // Latitude
	Longitude float64 // Longitude
	Time      Time    // Time Stamp
	Validity  string  // validity - A-valid
}

// newGLL constructor
func newGLL(s BaseSentence) (GLL, error) {
	p := newParser(s, "")
	p.AssertType(PrefixGLL)
	p.AssertTalker("GP")
	return GLL{
		BaseSentence: s,
		Latitude:     p.LatLong(0, 1, "latitude"),
		Longitude:    p.LatLong(2, 3, "longitude"),
		Time:         p.Time(4, "time"),
		Validity:     p.EnumString(5, "validity", ValidGLL, InvalidGLL),
	}, p.Err()
}
