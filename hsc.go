package nmea

const (
	// TypeHSC type of HSC sentence for Heading steering command
	TypeHSC = "HSC"
)

// HSC - Heading steering command
// https://gpsd.gitlab.io/gpsd/NMEA.html#_hsc_heading_steering_command
// https://www.tronico.fi/OH6NT/docs/NMEA0183.pdf (page 11)
//
// Format: $--HSC, x.x, T, x.x, M,a*hh<CR><LF>
// Example: $FTHSC,40.12,T,39.11,M*5E
type HSC struct {
	BaseSentence
	TrueHeading         float64 //  Heading Degrees, True
	TrueHeadingType     string  //  T = True
	MagneticHeading     float64 // Heading Degrees, Magnetic
	MagneticHeadingType string  // M = Magnetic
}

// newHSC constructor
func newHSC(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeHSC)
	return HSC{
		BaseSentence:        s,
		TrueHeading:         p.Float64(0, "true heading"),
		TrueHeadingType:     p.EnumString(1, "true heading type", HeadingTrue),
		MagneticHeading:     p.Float64(2, "magnetic heading"),
		MagneticHeadingType: p.EnumString(3, "magnetic heading type", HeadingMagnetic),
	}, p.Err()
}
