package nmea

import "strconv"

const (
	// TypeMTK type for PMTK sentences
	TypeMTK = "PMTK"
)

// MTK is the Time, position, and fix related data of the receiver.
type MTK struct {
	BaseSentence
	Cmd,
	Flag int
}

// newMTK constructor
func newMTK(s BaseSentence) (MTK, error) {
	cmd, err := strconv.Atoi(s.Fields[0])
	if err != nil {
		return MTK{}, err
	}
	flag, err := strconv.Atoi(s.Fields[1])
	if err != nil {
		return MTK{}, err
	}
	return MTK{
		BaseSentence: s,
		Cmd:          cmd,
		Flag:         flag,
	}, nil
}
