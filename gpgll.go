package nmea

const (
	// PrefixGPGLL prefix for GPGLL sentence type
	PrefixGPGLL = "GPGLL"
	// ValidGLL character
	ValidGLL = "A"
	// InvalidGLL character
	InvalidGLL = "V"
)

// GPGLL is Geographic Position, Latitude / Longitude and time.
// http://aprs.gids.nl/nmea/#gll
type GPGLL struct {
	Sentence
	Latitude  LatLong // Latitude
	Longitude LatLong // Longitude
	Time      Time    // Time Stamp
	Validity  string  // validity - A-valid
}

// NewGPGLL constructor
func NewGPGLL(s Sentence) (GPGLL, error) {
	p := newParser(s, PrefixGPGLL)
	return GPGLL{
		Sentence:  s,
		Latitude:  p.LatLong(0, 1, "latitude"),
		Longitude: p.LatLong(2, 3, "longitude"),
		Time:      p.Time(4, "time"),
		Validity:  p.EnumString(5, "validity", ValidGLL, InvalidGLL),
	}, p.Err()
}
