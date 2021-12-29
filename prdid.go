package nmea

const (
	// TypePRDID type of PRDID sentence for vessel pitch, roll and heading
	TypePRDID = "RDID"
)

// PRDID is proprietary sentence for vessel pitch, roll and heading.
// https://www.xsens.com/hubfs/Downloads/Manuals/MT_Low-Level_Documentation.pdf (page 37)
//
// Format: $PRDID,aPPP.PP,bRRR.RR,HHH.HH*hh<CR><LF>
// Example: $PRDID,-10.37,2.34,230.34*AA
type PRDID struct {
	BaseSentence
	Pitch   float64 // Pitch in degrees (positive bow up)
	Roll    float64 // Roll in degrees (positive port up)
	Heading float64 // True heading in degrees
}

// newPRDID constructor
func newPRDID(s BaseSentence) (PRDID, error) {
	p := NewParser(s)
	p.AssertType(TypePRDID)
	m := PRDID{
		BaseSentence: s,
		Pitch:        p.Float64(0, "pitch"),
		Roll:         p.Float64(1, "roll"),
		Heading:      p.Float64(2, "heading"),
	}
	return m, p.Err()
}
