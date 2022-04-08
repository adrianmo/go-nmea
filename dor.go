package nmea

const (
	// TypeDOR type of DOR sentence for Door Status Detection
	TypeDOR = "DOR"

	// TypeSingleDoorDOR is type for single door related event
	TypeSingleDoorDOR = "E"
	// TypeFaultDOR is type for fault with door
	TypeFaultDOR = "F"
	// TypeSectionDOR is type for section of doors related event
	TypeSectionDOR = "S"

	// DoorStatusOpenDOR is status for open door
	DoorStatusOpenDOR = "O"
	// DoorStatusClosedDOR is status for closed door
	DoorStatusClosedDOR = "C"
	// DoorStatusFaultDOR is status for fault with door
	DoorStatusFaultDOR = "X"

	// SwitchSettingHarbourModeDOR is setting for Harbour mode (allowed open)
	SwitchSettingHarbourModeDOR = "O"
	// SwitchSettingSeaModeDOR is setting for Sea mode (ordered closed)
	SwitchSettingSeaModeDOR = "C"
)

// DOR - Door Status Detection
// Source: "Interfacing Voyage Data Recorder Systems, AutroSafe Interactive Fire-Alarm System, 116-P-BSL336/EE, RevA 2007-01-25,
// Autronica Fire and Security AS " (page 32 | p.8.1.4)
// https://product.autronicafire.com/fileshare/fileupload/14251/bsl336_ee.pdf
//
// Format: $FRDOR,a,hhmmss,aa,aa,xxx,xxx,a,a,c--c*hh<CR><LF>
// Example: $FRDOR,E,233042,FD,FP,000,010,C,C,Door Closed : TEST FPA Name*4D
type DOR struct {
	BaseSentence

	// Type is type of the message
	// * E – Single door
	// * F – Fault
	// * S – Section (whole or part of section)
	Type string

	// Time is Event Time
	Time Time

	// SystemIndicator is system indicator. Detector system type with 2 char identifier.
	// * WT - watertight
	// * WS - semi watertight
	// * FD - fire door
	// * HD - hull door
	// * OT - other
	// could be more
	// https://www.nmea.org/Assets/20190303%20nmea%200183%20talker%20identifier%20mnemonics.pdf
	SystemIndicator string

	// DivisionIndicator1 is first division indicator for locating origin detector for this message
	DivisionIndicator1 string

	// DivisionIndicator2 is second division indicator for locating origin detector for this message
	DivisionIndicator2 int64

	// DoorNumberOrCount is Door number or activated door count (seems to be field with overloaded meaning)
	DoorNumberOrCount int64

	// DoorStatus is Door status
	// * O – Open
	// * C – Closed
	// * X – Fault
	// could be more
	DoorStatus string

	// SwitchSetting is  Mode switch setting
	// * O – Harbour mode (allowed open)
	// * C – Sea mode (ordered closed)
	SwitchSetting string

	// Message's description text (could be cut to fit max packet length)
	Message string
}

// newDOR constructor
func newDOR(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeDOR)
	return DOR{
		BaseSentence:       s,
		Type:               p.EnumString(0, "message type", TypeSingleDoorDOR, TypeFaultDOR, TypeSectionDOR),
		Time:               p.Time(1, "time"),
		SystemIndicator:    p.String(2, "system indicator"),
		DivisionIndicator1: p.String(3, "division indicator 1"),
		DivisionIndicator2: p.Int64(4, "division indicator 2"),
		DoorNumberOrCount:  p.Int64(5, "door number or count"),
		DoorStatus:         p.EnumString(6, "door state", DoorStatusOpenDOR, DoorStatusClosedDOR, DoorStatusFaultDOR),
		SwitchSetting:      p.EnumString(7, "switch setting mode", SwitchSettingHarbourModeDOR, SwitchSettingSeaModeDOR),
		Message:            p.String(8, "message"),
	}, p.Err()
}
