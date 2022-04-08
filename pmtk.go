package nmea

const (
	// TypePMTK001 type of acknowledgement sentence for MTK protocol
	TypePMTK001 = "MTK001"
)

// PMTK001 is sentence for acknowledgement of previously sent command/packet
// https://www.rhydolabz.com/documents/25/PMTK_A11.pdf
// https://www.sparkfun.com/datasheets/GPS/Modules/PMTK_Protocol.pdf
//
// The maximum length of each packet is restricted to 255 bytes which is longer than NMEA0183 82 bytes.
//
// Format: $PMTK001,c-c,d*hh<CR><LF>
// Example: $PMTK001,101,0*33<CR><LF>
type PMTK001 struct {
	BaseSentence

	// Cmd is command/packet acknowledgement is sent for.
	// Three bytes character string. From "000" to "999".
	Cmd int64

	// Flag is acknowledgement status for previously sent command/packet
	// 0 = invalid command/packet type
	// 1 = unsupported command packet type
	// 2 = valid command/packet, but action failed
	// 3 = valid command/packet and action succeeded
	Flag int64
}

// newPMTK001 constructor
func newPMTK001(s BaseSentence) (Sentence, error) {
	p := NewParser(s)

	cmd := p.Int64(0, "command")
	flag := p.Int64(1, "flag")
	return PMTK001{
		BaseSentence: s,
		Cmd:          cmd,
		Flag:         flag,
	}, p.Err()
}
