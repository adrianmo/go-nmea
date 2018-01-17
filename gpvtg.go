package nmea

const (
	// PrefixGPVTG prefix
	PrefixGPVTG = "GPVTG"
)

// GPVTG represents track & speed data.
// http://aprs.gids.nl/nmea/#vtg
type GPVTG struct {
	Sentence
	TrueTrack        float64
	MagneticTrack    float64
	GroundSpeedKnots float64
	GroundSpeedKPH   float64
}

// NewGPVTG constructor
func NewGPVTG(sentence Sentence) GPVTG {
	s := new(GPVTG)
	s.Sentence = sentence
	return *s
}

// GetSentence getter
func (s GPVTG) GetSentence() Sentence {
	return s.Sentence
}

// Parse parses the GPVTG sentence into this struct.
// e.g: $GPVTG,360.0,T,348.7,M,000.0,N,000.0,K*43
func (s *GPVTG) parse() error {
	p := newParser(s.Sentence, PrefixGPVTG)
	s.TrueTrack = p.Float64(0, "true track")
	s.MagneticTrack = p.Float64(2, "magnetic track")
	s.GroundSpeedKnots = p.Float64(4, "ground speed (knots)")
	s.GroundSpeedKPH = p.Float64(6, "ground speed (km/h)")
	return p.Err()
}
