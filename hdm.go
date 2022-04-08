package nmea

const (
	// TypeHDM type of HDM sentence for vessel heading in degrees with respect to magnetic north
	TypeHDM = "HDM"
	// MagneticHDM for valid Magnetic heading
	MagneticHDM = "M"
)

// HDM is vessel heading in degrees with respect to magnetic north produced by any device or system producing magnetic heading.
// https://gpsd.gitlab.io/gpsd/NMEA.html#_hdm_heading_magnetic
//
// Format: $--HDM,xxx.xx,M*hh<CR><LF>
// Example: $HCHDM,093.8,M*2B
type HDM struct {
	BaseSentence
	Heading       float64 // Heading in degrees
	MagneticValid bool    // Heading is respect to magnetic north
}

// newHDM constructor
func newHDM(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeHDM)
	m := HDM{
		BaseSentence:  s,
		Heading:       p.Float64(0, "heading"),
		MagneticValid: p.EnumString(1, "magnetic", MagneticHDM) == MagneticHDM,
	}
	return m, p.Err()
}
