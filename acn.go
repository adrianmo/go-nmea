package nmea

const (
	// TypeACN type of ACN sentence for alert command
	TypeACN = "ACN"
)

// ACN - Alert command. Used for acknowledge, silence, responsibility transfer and to request repeat of alert.
// https://www.furuno.it/docs/INSTALLATION%20MANUALIME44900D_FA170.pdf Furuno CLASS A AIS Model FA-170 (page 49)
// https://www.furuno.it/docs/INSTALLATION%20MANUALgp170_installation_manual.pdf GPS NAVIGATOR Model GP-170 (page 42)
//
// Format: $--ACN,hhmmss.ss,AAA,x.x,x.x,A,A*hh<CR><LF>
// Example: $VRACN,220516,BPMP1,A,A,Bilge pump alarm1*43
type ACN struct {
	BaseSentence

	// Time is time of alarm condition change, UTC (000000.00 - 240001.00)
	Time Time // 0

	// ManufacturerMnemonicCode is manufacturer mnemonic code
	ManufacturerMnemonicCode string // 1

	// AlertIdentifier is alert identifier (001 to 99999)
	AlertIdentifier int64 // 2

	// AlertInstance is alert instance
	AlertInstance int64 // 3

	// Command is Alert command
	// * A - acknowledge,
	// * Q - request/repeat information
	// * O - responsibility transfer
	// * S - silence
	Command string // 4

	// State is alarm state
	// * C - command
	// * possible more classifier values but these are not mentioned in manual
	State string // 5
}

// newACN constructor
func newACN(s BaseSentence) (ACN, error) {
	p := NewParser(s)
	p.AssertType(TypeACN)
	return ACN{
		BaseSentence:             s,
		Time:                     p.Time(0, "time"),
		ManufacturerMnemonicCode: p.String(1, "manufacturer mnemonic code"),
		AlertIdentifier:          p.Int64(2, "alert identifier"),
		AlertInstance:            p.Int64(3, "alert instance"),
		Command:                  p.EnumString(4, "alert command", AlertCommandAcknowledge, AlertCommandRequestRepeatInformation, AlertCommandResponsibilityTransfer, AlertCommandSilence),
		State:                    p.String(5, "alarm state"),
	}, p.Err()
}
