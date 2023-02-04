package nmea

const (
	// TypeTTD type of TTD sentence for tracked target data.
	TypeTTD = "TTD"
)

// TTD is sentence used by radars to transmit tracked targets data.
// https://fcc.report/FCC-ID/ADB9ZWRTR100/2768717.pdf (page 1) FURUNO MARINE RADAR, model FAR-15XX manual
//
// Format: !--TTD,hh,hh,x,s--s,x*hh<CR><LF>
// Example: !RATTD,1A,01,1,177KQJ5000G?tO`K>RA1wUbN0TKH,0*5C
type TTD struct {
	BaseSentence
	// NumFragments is total hex number of fragments/sentences need to transfer the message (1 - FF)
	NumFragments int64 // 0
	//  FragmentNumber is current fragment/sentence number (1 - FF)
	FragmentNumber int64 // 1
	// MessageID is sequential message identifier (0 - 9, null)
	MessageID int64 // 2
	// Payload is encapsulated tracked target data (6 bit binary-converted data)
	Payload []byte // 3
	// 4 - Number of fill bits (0 - 5)
}

// newTTD constructor
func newTTD(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeTTD)
	m := TTD{
		BaseSentence:   s,
		NumFragments:   p.HexInt64(0, "number of fragments"),
		FragmentNumber: p.HexInt64(1, "fragment number"),
		MessageID:      p.Int64(2, "sequence number"),
		Payload:        p.SixBitASCIIArmour(3, int(p.Int64(4, "number of padding bits")), "payload"),
	}
	return m, p.Err()
}
