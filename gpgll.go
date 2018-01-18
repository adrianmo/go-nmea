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
	s := new(GPGLL)
	s.Sentence = sentence
	return *s, s.parse()
}

// GetSentence getter
func (s GPGLL) GetSentence() Sentence {
	return s.Sentence
}

func (s *GPGLL) parse() error {
	p := newParser(s.Sentence, PrefixGPGLL)
	s.Latitude = p.LatLong(0, 1, "latitude")
	s.Longitude = p.LatLong(2, 3, "longitude")
	s.Time = p.Time(4, "time")
	s.Validity = p.String(5, "validity")
	if s.Validity != ValidGLL && s.Validity != InvalidGLL {
		p.SetErr("validity", s.Validity)
	}
	return p.Err()
}
