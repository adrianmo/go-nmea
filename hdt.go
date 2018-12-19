package nmea

const (
	// PrefixHDT prefix of HDT sentence type
	PrefixHDT = "HDT"
)

// HDT is the Actual vessel heading in degrees True.
// http://aprs.gids.nl/nmea/#hdt
type HDT struct {
	BaseSentence
	Heading float64 // Heading in degrees
	True    bool    // Heading is relative to true north
}

// newHDT constructor
func newHDT(s BaseSentence) (HDT, error) {
	p := newParser(s, PrefixHDT)
	p.AssertType(PrefixHDT)
	p.AssertTalker("GP")
	m := HDT{
		BaseSentence: s,
		Heading:      p.Float64(0, "heading"),
		True:         p.EnumString(1, "true", "T") == "T",
	}
	return m, p.Err()
}
