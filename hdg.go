package nmea

const (
	// TypeHDG type of HDG sentence for vessel heading, deviation and variation with respect to magnetic north.
	TypeHDG = "HDG"
)

// HDG is vessel heading (in degrees), deviation and variation with respect to magnetic north produced by any
// device or system producing magnetic reading.
// https://gpsd.gitlab.io/gpsd/NMEA.html#_hdg_heading_deviation_variation
//
// Format: $--HDG,x.x,y.y,a,z.z,a*hr<CR><LF>
// Example: $HCHDG,98.3,0.0,E,12.6,W*57
type HDG struct {
	BaseSentence
	Heading            float64 // Heading in degrees
	Deviation          float64 //  Magnetic Deviation in degrees
	DeviationDirection string  // Magnetic Deviation direction, E = Easterly, W = Westerly
	Variation          float64 //  Magnetic Variation in degrees
	VariationDirection string  // Magnetic Variation direction, E = Easterly, W = Westerly
}

// newHDG constructor
func newHDG(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeHDG)
	m := HDG{
		BaseSentence:       s,
		Heading:            p.Float64(0, "heading"),
		Deviation:          p.Float64(1, "deviation"),
		DeviationDirection: p.EnumString(2, "deviation direction", East, West),
		Variation:          p.Float64(3, "variation"),
		VariationDirection: p.EnumString(4, "variation direction", East, West),
	}
	return m, p.Err()
}
