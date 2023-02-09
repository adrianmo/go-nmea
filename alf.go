package nmea

const (
	// TypeALF type of ALF sentence for alert sentence
	TypeALF = "ALF"
)

// ALF - Alert sentence
// https://fcc.report/FCC-ID/ADB9ZWRTR100/2768717.pdf (page 6) FURUNO MARINE RADAR, model FAR-15XX manual
// https://www.rcom.nl/wp-content/uploads/2017/09/NSR-NGR-3000-Users-Manual.pdf (page 47)
//
// Format: $--ALF,x,x,x,hhmmss.ss,a,a,a,aaa,x.x,x.x,x.x,x,c--c,*hh<CR><LF>
// Example: $VDALF,1,0,1,220516,B,A,S,SAL,001,1,2,0,My alarm*2c
type ALF struct {
	BaseSentence

	// NumFragments is total number of ALF sentences this message (1, 2)
	NumFragments int64 // 0

	//  FragmentNumber is current fragment/sentence number (1 - 2)
	FragmentNumber int64 // 1

	// MessageID is sequential message identifier (0 - 9)
	MessageID int64 // 2

	// Time is time of last change (000000.00 - 240001.00 / null)
	Time Time // 3

	// Category is alert category (A/B/C/null)
	// A - Alert category A,
	// B - Alert category B,
	// C - Alert category C,
	Category string // 4

	// Priority is alert priority (A/W/C/null)
	// E - Emergency Alarm: E, for use with Bridge alert management
	// A - Alarm,
	// W - Warning,
	// C - Caution,
	// null
	Priority string // 5

	// State is alert state (A/S/O/U/V/N/null)
	// A - Acknowledged
	// S - Silence
	// O - Active-responsiblity transferred
	// U - Rectified-unacknowledged
	// V - Not acknowledged
	// N - Normal state
	// null
	State string // 6

	// ManufacturerMnemonicCode is manufacturer mnemonic code
	ManufacturerMnemonicCode string // 7

	// AlertIdentifier is alert identifier (001 to 99999)
	AlertIdentifier int64 // 8

	// AlertInstance is alert instance
	AlertInstance int64 // 9

	// RevisionCounter is revision counter (1 - 99)
	RevisionCounter int64 // 10

	// EscalationCounter is escalation counter (0 - 9)
	EscalationCounter int64 // 11

	// Text is alarm text
	Text string // 12
}

// newALF constructor
func newALF(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeALF)
	return ALF{
		BaseSentence:             s,
		NumFragments:             p.Int64(0, "number of fragments"),
		FragmentNumber:           p.Int64(1, "fragment number"),
		MessageID:                p.Int64(2, "message ID"),
		Time:                     p.Time(3, "time"),
		Category:                 p.EnumString(4, "alarm category", "A", "B", "C"),
		Priority:                 p.EnumString(5, "alarm priority", "E", "A", "C", "W"),
		State:                    p.EnumString(6, "alarm state", "A", "S", "O", "U", "V", "N"),
		ManufacturerMnemonicCode: p.String(7, "manufacturer mnemonic code"),
		AlertIdentifier:          p.Int64(8, "alert identifier"),
		AlertInstance:            p.Int64(9, "alert instance"),
		RevisionCounter:          p.Int64(10, "revision counter"),
		EscalationCounter:        p.Int64(11, "escalation counter"),
		Text:                     p.String(12, "alert text"),
	}, p.Err()
}
