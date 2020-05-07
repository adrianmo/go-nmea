package nmea

const (
	// TypeMTK type for PMTK sentences
	TypeMTK = "PMTK"
)

// MTK is the Time, position, and fix related data of the receiver.
type MTK struct {
	BaseSentence
	Cmd,
	Flag int64
}

// newMTK constructor
func newMTK(s BaseSentence) (MTK, error) {
	p := NewParser(s)
	cmd := p.Int64(0, "command")
	flag := p.Int64(1, "flag")
	return MTK{
		BaseSentence: s,
		Cmd:          cmd,
		Flag:         flag,
	}, p.Err()
}
