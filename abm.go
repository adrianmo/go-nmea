package nmea

const (
	// TypeABM type of ABM sentence for AIS addressed binary and safety related message
	TypeABM = "ABM"
)

// ABM - AIS addressed binary and safety related message
// https://fcc.report/FCC-ID/ADB9ZWRTR100/2768717.pdf (page 6) FURUNO MARINE RADAR, model FAR-15XX manual
//
// Format: !--ABM,x,x,x,xxxxxxxxx,x,xx,s--s,x,*hh<CR><LF>
// Example: !AIABM,26,2,1,3381581370,3,8,177KQJ5000G?tO`K>RA1wUbN0TKH,0*02
type ABM struct {
	BaseSentence

	// NumFragments is total number of fragments/sentences need to transfer the message (1 - 9)
	NumFragments int64 // 0

	//  FragmentNumber is current fragment/sentence number (1 - 9)
	FragmentNumber int64 // 1

	// MessageID is sequential message identifier (0 - 3)
	MessageID int64 // 2

	// MMSI is The MMSI of destination AIS unit for the ITU-R M.1371 message (10 digits or empty)
	MMSI string // 3

	// Channel is AIS channel for broadcast of the radio message (0 - 3)
	// 0 - no broadcast
	// 1 - on AIS channel A
	// 2 - on AIS channel B
	// 3 - broadcast on both AIS channels
	Channel string // 4

	// VDLMessageNumber is VDL message number (6/12), see ITU-R M.1371
	VDLMessageNumber int64 // 5

	// Payload is encapsulated data (6 bit binary-converted data) (1 - 63 bytes)
	Payload []byte // 6
	// 7 - Number of fill bits (0 - 5)
}

// newABM constructor
func newABM(s BaseSentence) (ABM, error) {
	p := NewParser(s)
	p.AssertType(TypeABM)
	return ABM{
		BaseSentence:     s,
		NumFragments:     p.Int64(0, "number of fragments"),
		FragmentNumber:   p.Int64(1, "fragment number"),
		MessageID:        p.Int64(2, "message ID"),
		MMSI:             p.String(3, "MMSI"),
		Channel:          p.String(4, "channel"),
		VDLMessageNumber: p.Int64(5, "VDL message number"),
		Payload:          p.SixBitASCIIArmour(6, int(p.Int64(7, "number of padding bits")), "payload"),
	}, p.Err()
}
