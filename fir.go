package nmea

const (
	// TypeFIR type of FIR sentence for Fire Detection
	TypeFIR = "FIR"

	// TypeEventOrAlarmFIR is Event, Fire Alarm type
	TypeEventOrAlarmFIR = "E"
	// TypeFaultFIR is type for Fault
	TypeFaultFIR = "F"
	// TypeDisablementFIR is type for detector disablement
	TypeDisablementFIR = "D"

	// ConditionActivationFIR is activation condition
	ConditionActivationFIR = "A"
	// ConditionNonActivationFIR is non-activation condition
	ConditionNonActivationFIR = "V"
	// ConditionUnknownFIR is unknown condition
	ConditionUnknownFIR = "X"

	// AlarmStateAcknowledgedFIR is value for alarm acknowledgement
	AlarmStateAcknowledgedFIR = "A"
	// AlarmStateNotAcknowledgedFIR is value for alarm being not acknowledged
	AlarmStateNotAcknowledgedFIR = "V"
)

// FIR - Fire Detection event with time and location
// Source: "Interfacing Voyage Data Recorder Systems, AutroSafe Interactive Fire-Alarm System, 116-P-BSL336/EE, RevA 2007-01-25,
// Autronica Fire and Security AS " (page 39 | p.8.1.6)
// https://product.autronicafire.com/fileshare/fileupload/14251/bsl336_ee.pdf
//
// Format: $FRFIR,a,hhmmss,aa,aa,xxx,xxx,a,a,c--c*hh<CR><LF>
// Example: $FRFIR,E,103000,FD,PT,000,007,A,V,Fire Alarm : TEST PT7 Name TEST DZ2 Name*7A
type FIR struct {
	BaseSentence

	// Type is type of the message
	// * E – Event, Fire Alarm
	// * F – Fault
	// * D – Disablement
	Type string

	// Time is Event Time
	Time Time

	// SystemIndicator is system indicator. Detector system type with 2 char identifier.
	// * FD Generic fire detector
	// * FH Heat detector
	// * FS Smoke detector
	// * FD Smoke and heat detector
	// * FM Manual call point
	// * GD Any gas detector
	// * GO Oxygen gas detector
	// * GS Hydrogen sulphide gas detector
	// * GH Hydro-carbon gas detector
	// * SF Sprinkler flow switch
	// * SV Sprinkler manual valve release
	// * CO CO2 manual release
	// * OT Other
	SystemIndicator string

	// DivisionIndicator1 is first division indicator for locating origin detector for this message
	DivisionIndicator1 string

	// DivisionIndicator2 is second division indicator for locating origin detector for this message
	DivisionIndicator2 int64

	// FireDetectorNumberOrCount is Fire detector number or activated detectors count (seems to be field with overloaded meaning)
	FireDetectorNumberOrCount int64

	// Condition describes the condition triggering current message
	// * A – Activation
	// * V – Non-activation
	// * X – State unknown
	Condition string

	// AlarmAckState is Alarm's acknowledge state
	// * A – Acknowledged
	// * V – Not acknowledged
	AlarmAckState string

	// Message's description text (could be cut to fit max packet length)
	Message string
}

// newFIR constructor
func newFIR(s BaseSentence) (FIR, error) {
	p := NewParser(s)
	p.AssertType(TypeFIR)
	return FIR{
		BaseSentence:              s,
		Type:                      p.EnumString(0, "message type", TypeEventOrAlarmFIR, TypeFaultFIR, TypeDisablementFIR),
		Time:                      p.Time(1, "time"),
		SystemIndicator:           p.String(2, "system indicator"),
		DivisionIndicator1:        p.String(3, "division indicator 1"),
		DivisionIndicator2:        p.Int64(4, "division indicator 2"),
		FireDetectorNumberOrCount: p.Int64(5, "fire detector number or count"),
		Condition:                 p.EnumString(6, "condition", ConditionActivationFIR, ConditionNonActivationFIR, ConditionUnknownFIR),
		AlarmAckState:             p.EnumString(7, "alarm acknowledgement state", AlarmStateAcknowledgedFIR, AlarmStateNotAcknowledgedFIR),
		Message:                   p.String(8, "message"),
	}, p.Err()
}
