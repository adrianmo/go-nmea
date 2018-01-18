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
func NewGNRMC(s Sentence) (GNRMC, error) {
	p := newParser(s, PrefixGNRMC)
	m := GNRMC{
		Sentence:  s,
		Time:      p.Time(0, "time"),
		Validity:  p.EnumString(1, "validity", ValidRMC, InvalidRMC),
		Latitude:  p.LatLong(2, 3, "latitude"),
		Longitude: p.LatLong(4, 5, "longitude"),
		Speed:     p.Float64(6, "speed"),
		Course:    p.Float64(7, "course"),
		Date:      p.String(8, "date"),
	}
	if !p.Empty(9, "variation") {
		m.Variation = p.Float64(9, "variation")
		if p.EnumString(10, "direction", "W", "E") == "W" {
			m.Variation = 0 - m.Variation
		}
	}
	return m, p.Err()
}

// GetSentence getter
func (s GNRMC) GetSentence() Sentence {
	return s.Sentence
}
