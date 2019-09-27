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
type GLL struct {
	BaseSentence
	Latitude     float64 // Latitude
	Longitude    float64 // Longitude
	LatDirection string  // Latitude direction.
	LonDirection string  // Longitude direction.
	Time         Time    // Time Stamp
	Validity     string  // validity - A-valid
}

// newGLL constructor
func newGLL(s BaseSentence) (GLL, error) {
	p := newParser(s)
	p.AssertType(TypeGLL)
	return GLL{
		BaseSentence: s,
		Latitude:     p.LatLong(0, 1, "latitude"),
		Longitude:    p.LatLong(2, 3, "longitude"),
		LatDirection: p.EnumString(1, "latitude direction", North, South),
		LonDirection: p.EnumString(3, "longitude direction", East, West),
		Time:         p.Time(4, "time"),
		Validity:     p.EnumString(5, "validity", ValidGLL, InvalidGLL),
	}, p.Err()
}
