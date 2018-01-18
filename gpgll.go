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
func NewGPGLL(sentence Sentence) (GPGLL, error) {
	p := newParser(sentence, PrefixGPGLL)
	return GPGLL{
		Sentence:  sentence,
		Latitude:  p.LatLong(0, 1, "latitude"),
		Longitude: p.LatLong(2, 3, "longitude"),
		Time:      p.Time(4, "time"),
		Validity:  p.EnumString(5, "validity", ValidGLL, InvalidGLL),
	}, p.Err()
}

// GetSentence getter
func (s GPGLL) GetSentence() Sentence {
	return s.Sentence
}
