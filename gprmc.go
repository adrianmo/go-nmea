package nmea

const (
	// PrefixGPRMC prefix of GPRMC sentence type
	PrefixGPRMC = "GPRMC"
	// ValidRMC character
	ValidRMC = "A"
	// InvalidRMC character
	InvalidRMC = "V"
)

// GPRMC is the Recommended Minimum Specific GNSS data.
// http://aprs.gids.nl/nmea/#rmc
type GPRMC struct {
	Sentence
	Time      Time    // Time Stamp
	Validity  string  // validity - A-ok, V-invalid
	Latitude  LatLong // Latitude
	Longitude LatLong // Longitude
	Speed     float64 // Speed in knots
	Course    float64 // True course
	Date      string  // Date
	Variation float64 // Magnetic variation
}

// NewGPRMC constructor
func NewGPRMC(sentence Sentence) (GPRMC, error) {
	p := newParser(sentence, PrefixGPRMC)
	s := GPRMC{
		Sentence:  sentence,
		Time:      p.Time(0, "time"),
		Validity:  p.EnumString(1, "validity", ValidRMC, InvalidRMC),
		Latitude:  p.LatLong(2, 3, "latitude"),
		Longitude: p.LatLong(4, 5, "longitude"),
		Speed:     p.Float64(6, "speed"),
		Course:    p.Float64(7, "course"),
		Date:      p.String(8, "date"),
	}
	if !p.Empty(9, "variation") {
		s.Variation = p.Float64(9, "variation")
		if p.EnumString(10, "variation", "W", "E") == "W" {
			s.Variation = 0 - s.Variation
		}
	}
	return s, p.Err()
}

// GetSentence getter
func (s GPRMC) GetSentence() Sentence {
	return s.Sentence
}
