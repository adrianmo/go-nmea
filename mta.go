package nmea

const (
	// TypeMTA type of MTA sentence for Air Temperature
	TypeMTA = "MTA"
)

// MTA - Air Temperature (obsolete, use XDR instead)
// https://www.nmea.org/Assets/100108_nmea_0183_sentences_not_recommended_for_new_designs.pdf (page 7)
//
// Format: $--MTA,x.x,C*hh<CR><LF>
// Example: $IIMTA,13.3,C*04
type MTA struct {
	BaseSentence
	Temperature float64 // temperature
	Unit        string  // unit of temperature, should be degrees Celsius
}

// newMTA constructor
func newMTA(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeMTA)
	return MTA{
		BaseSentence: s,
		Temperature:  p.Float64(0, "temperature"),
		Unit:         p.EnumString(1, "temperature unit", TemperatureCelsius),
	}, p.Err()
}
