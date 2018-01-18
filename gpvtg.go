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

// NewGPVTG parses the GPVTG sentence into this struct.
// e.g: $GPVTG,360.0,T,348.7,M,000.0,N,000.0,K*43
func NewGPVTG(sentence Sentence) (GPVTG, error) {
	p := newParser(sentence, PrefixGPVTG)
	return GPVTG{
		Sentence:         sentence,
		TrueTrack:        p.Float64(0, "true track"),
		MagneticTrack:    p.Float64(2, "magnetic track"),
		GroundSpeedKnots: p.Float64(4, "ground speed (knots)"),
		GroundSpeedKPH:   p.Float64(6, "ground speed (km/h)"),
	}, p.Err()
}

// GetSentence getter
func (s GPVTG) GetSentence() Sentence {
	return s.Sentence
}
