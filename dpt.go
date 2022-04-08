package nmea

const (
	// TypeDPT type for DPT sentences
	TypeDPT = "DPT"
)

// DPT - Depth of Water
// https://gpsd.gitlab.io/gpsd/NMEA.html#_dpt_depth_of_water
//
// Format: $--DPT,x.x,x.x,x.x*hh<CR><LF>
// Example: $SDDPT,0.5,0.5,*7B
//          $INDPT,2.3,0.0*46
type DPT struct {
	BaseSentence
	Depth      float64 // Water depth relative to transducer, meters
	Offset     float64 // offset from transducer
	RangeScale float64 // OPTIONAL, Maximum range scale in use (NMEA 3.0 and above)
}

// newDPT constructor
func newDPT(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeDPT)
	dpt := DPT{
		BaseSentence: s,
		Depth:        p.Float64(0, "depth"),
		Offset:       p.Float64(1, "offset"),
	}
	if len(p.Fields) > 2 {
		dpt.RangeScale = p.Float64(2, "range scale")
	}
	return dpt, p.Err()
}
