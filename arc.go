package nmea

const (
	// TypeARC type of ARC sentence for alert command refused
	TypeARC = "ARC"
)

const (
	// AlertCommandAcknowledge means acknowledge
	AlertCommandAcknowledge = "A"
	// AlertCommandRequestRepeatInformation means request/repeat information
	AlertCommandRequestRepeatInformation = "Q"
	// AlertCommandResponsibilityTransfer means responsibility transfer
	AlertCommandResponsibilityTransfer = "O"
	// AlertCommandSilence means silence
	AlertCommandSilence = "S"
)

// ARC - Alert command refused
// https://fcc.report/FCC-ID/ADB9ZWRTR100/2768717.pdf (page 7) FURUNO MARINE RADAR, model FAR-15XX manual
//
// Format: $--ARC,hhmmss.ss,aaa,x.x,x.x,c*hh<CR><LF>
// Example: $RAARC,220516,TCK,002,1,A*73
type ARC struct {
	BaseSentence

	// Time is UTC Time
	Time Time // 0

	// ManufacturerMnemonicCode is manufacturer mnemonic code
	ManufacturerMnemonicCode string // 1

	// AlertIdentifier is alert identifier (001 to 99999)
	AlertIdentifier int64 // 2

	// AlertInstance is alert instance
	AlertInstance int64 // 3

	// Command is Refused alert command
	// A - acknowledge
	// Q - request/repeat information
	// O - responsibility transfer
	// S - silence
	Command string // 4
}

// newARC constructor
func newARC(s BaseSentence) (ARC, error) {
	p := NewParser(s)
	p.AssertType(TypeARC)
	return ARC{
		BaseSentence:             s,
		Time:                     p.Time(0, "time"),
		ManufacturerMnemonicCode: p.String(1, "manufacturer mnemonic code"),
		AlertIdentifier:          p.Int64(2, "alert identifier"),
		AlertInstance:            p.Int64(3, "alert instance"),
		Command:                  p.EnumString(4, "refused alert command", AlertCommandAcknowledge, AlertCommandRequestRepeatInformation, AlertCommandResponsibilityTransfer, AlertCommandSilence),
	}, p.Err()
}
