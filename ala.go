package nmea

const (
	// TypeALA type of ALA sentence for System Faults and alarms
	TypeALA = "ALA"
)

// ALA - System Faults and alarms
// Source: "Interfacing Voyage Data Recorder Systems, AutroSafe Interactive Fire-Alarm System, 116-P-BSL336/EE, RevA 2007-01-25,
// Autronica Fire and Security AS " (page 31 | p.8.1.3)
// https://product.autronicafire.com/fileshare/fileupload/14251/bsl336_ee.pdf
//
// Format: $FRALA,hhmmss,aa,aa,xx,xxx,a,a,c-cc*hh<CR><LF>
// Example: $FRALA,143955,FR,OT,00,901,N,V,Syst Fault : AutroSafe comm. OK*4F
type ALA struct {
	BaseSentence

	// Time is Event Time
	Time Time

	// SystemIndicator is system indicator of original alarm source. Detector system type with 2 char identifier.
	// Values not known
	// https://www.nmea.org/Assets/20190303%20nmea%200183%20talker%20identifier%20mnemonics.pdf
	SystemIndicator string

	// SubSystemIndicator is sub system equipment indicator of original alarm source
	SubSystemIndicator string

	// InstanceNumber is instance number of equipment/unit/item (00-99)
	InstanceNumber int64

	// Type is alarm type (000-999)
	Type int64

	// Condition describes the condition triggering current message
	// * N – Normal state (OK)
	// * H - Alarm state (fault);
	// could be more
	Condition string

	// AlarmAckState is Alarm's acknowledge state
	// * A – Acknowledged
	// * H - Harbour mode
	// * V – Not acknowledged
	// * O - Override
	// could be more
	AlarmAckState string

	// Message's description text (could be cut to fit max packet length)
	Message string
}

// newALA constructor
func newALA(s BaseSentence) (ALA, error) {
	p := NewParser(s)
	p.AssertType(TypeALA)
	return ALA{
		BaseSentence:       s,
		Time:               p.Time(0, "time"),
		SystemIndicator:    p.String(1, "system indicator"),
		SubSystemIndicator: p.String(2, "subsystem indicator"),
		InstanceNumber:     p.Int64(3, "instance number"),
		Type:               p.Int64(4, "type"),
		Condition:          p.String(5, "condition"),                   // string as there could be more
		AlarmAckState:      p.String(6, "alarm acknowledgement state"), // string as there could be more
		Message:            p.String(7, "message"),
	}, p.Err()
}
