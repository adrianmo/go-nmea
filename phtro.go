package nmea

const (
	// TypePHTRO type of PHTRO sentence for vessel pitch and roll
	TypePHTRO = "HTRO"
	// PHTROBowUP for bow up
	PHTROBowUP = "M"
	// PHTROBowDown for bow down
	PHTROBowDown = "P"
	// PHTROPortUP for port up
	PHTROPortUP = "T"
	// PHTROPortDown for port down
	PHTROPortDown = "B"
)

// PHTRO is proprietary sentence for vessel pitch and roll.
// https://www.igp.de/manuals/7-INS-InterfaceLibrary_MU-INSIII-AN-001-O.pdf (page 172)
//
// Format: $PHTRO,x.xx,a,y.yy,b*hh<CR><LF>
// Example: $PHTRO,10.37,P,177.62,T*65
type PHTRO struct {
	BaseSentence
	Pitch float64 // Pitch in degrees
	Bow   string  // "M" for bow up and "P" for bow down (2 digits after the decimal point)
	Roll  float64 // Roll in degrees
	Port  string  // "B" for port down and "T" for port up (2 digits after the decimal point)
}

// newPHTRO constructor
func newPHTRO(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypePHTRO)
	m := PHTRO{
		BaseSentence: s,
		Pitch:        p.Float64(0, "pitch"),
		Bow:          p.EnumString(1, "bow", PHTROBowUP, PHTROBowDown),
		Roll:         p.Float64(2, "roll"),
		Port:         p.EnumString(3, "port", PHTROPortUP, PHTROPortDown),
	}
	return m, p.Err()
}
