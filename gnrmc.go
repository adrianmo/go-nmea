package nmea

const (
	// PrefixGNRMC prefix of GNRMC sentence type
	PrefixGNRMC = "GNRMC"
)

// GNRMC is the Recommended Minimum Specific GNSS data.
// http://aprs.gids.nl/nmea/#rmc
type GNRMC struct {
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

// NewGNRMC constructor
func NewGNRMC(sentence Sentence) (GNRMC, error) {
	p := newParser(sentence, PrefixGNRMC)
	s := GNRMC{
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
		if p.EnumString(10, "direction", "W", "E") == "W" {
			s.Variation = 0 - s.Variation
		}
	}
	return s, p.Err()
}

// GetSentence getter
func (s GNRMC) GetSentence() Sentence {
	return s.Sentence
}
