package nmea

const (
	// TypeBBM type of BBM sentence for AIS broadcast binary message
	TypeBBM = "BBM"
)

// BBM - AIS broadcast binary message
// https://fcc.report/FCC-ID/ADB9ZWRTR100/2768717.pdf (page 7) FURUNO MARINE RADAR, model FAR-15XX manual
//
// Format: !--BBM,x,x,x,x,xx,s—s,x*hh<CR><LF>
// Example: !AIBBM,26,2,1,3,8,177KQJ5000G?tO`K>RA1wUbN0TKH,0*2C
type BBM struct {
	BaseSentence

	// NumFragments is total number of fragments/sentences need to transfer the message (1 - 9)
	NumFragments int64 // 0

	//  FragmentNumber is current fragment/sentence number (1 - 9)
	FragmentNumber int64 // 1

	// MessageID is sequential message identifier (0 - 9)
	MessageID int64 // 2

	// Channel is AIS channel for broadcast of the radio message (0 - 3)
	// 0 - no broadcast
	// 1 - on AIS channel A
	// 2 - on AIS channel B
	// 3 - broadcast on both AIS channels
	Channel string // 3

	// VDLMessageNumber is ITU-r M.1371 message number (8/14)
	VDLMessageNumber int64 // 4

	// Payload is encapsulated data (6 bit binary-converted data) (1 - 63 bytes)
	Payload []byte // 5
	// 6 - Number of fill bits (0 - 5)
}

// newBBM constructor
func newBBM(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeBBM)
	m := BBM{
		BaseSentence:     s,
		NumFragments:     p.Int64(0, "number of fragments"),
		FragmentNumber:   p.Int64(1, "fragment number"),
		MessageID:        p.Int64(2, "message ID"),
		Channel:          p.String(3, "channel"),
		VDLMessageNumber: p.Int64(4, "VDL message number"),
		Payload:          p.SixBitASCIIArmour(5, int(p.Int64(6, "number of padding bits")), "payload"),
	}
	return m, p.Err()
}
