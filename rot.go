package nmea

const (
	// TypeROT type of ROT sentence for vessel rate of turn
	TypeROT = "ROT"
	// ValidROT data is valid
	ValidROT = "A"
	// InvalidROT data is invalid
	InvalidROT = "V"
)

// ROT is sentence for rate of turn.
// https://gpsd.gitlab.io/gpsd/NMEA.html#_rot_rate_of_turn
//
// Format: $HEROT,-xxx.x,A*hh<CR><LF>
// Example: $HEROT,-11.23,A*07
type ROT struct {
	BaseSentence
	RateOfTurn float64 // rate of turn Z in deg/min (- means bow turns to port)
	Valid      bool    // "A" data valid,  "V" invalid data
}

func newROT(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeROT)
	return ROT{
		BaseSentence: s,
		RateOfTurn:   p.Float64(0, "rate of turn"),
		Valid:        p.EnumString(1, "status valid", ValidROT, InvalidROT) == ValidROT,
	}, p.Err()
}
