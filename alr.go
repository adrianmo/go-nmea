package nmea

const (
	// TypeALR type of ALR sentence for alert command refused
	TypeALR = "ALR"
)

// ALR - Set alarm state
// https://fcc.report/FCC-ID/ADB9ZWRTR100/2768717.pdf (page 7) FURUNO MARINE RADAR, model FAR-15XX manual
//
// Format: $--ALR,hhmmss.ss,xxx,A,A,c--c,*hh<CR><LF>
// Example: $RAALR,220516,BPMP1,A,A,Bilge pump alarm1*43
type ALR struct {
	BaseSentence

	// Time is time of alarm condition change, UTC (000000.00 - 240001.00)
	Time Time // 0

	// AlarmIdentifier is unique alarm number (identifier) at alarm source
	AlarmIdentifier int64 // 1

	// AlarmCondition is alarm condition (A/V)
	// A - threshold exceeded
	// V - not exceeded
	Condition string // 2

	// State is alarm state (A/V)
	// A - acknowledged
	// V - not acknowledged
	State string // 3

	// Description is alarm description text
	Description string // 4
}

// newALR constructor
func newALR(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeALR)
	return ALR{
		BaseSentence:    s,
		Time:            p.Time(0, "time"),
		AlarmIdentifier: p.Int64(1, "unique alarm number"),
		Condition:       p.EnumString(2, "alarm condition", StatusValid, StatusInvalid),
		State:           p.EnumString(3, "alarm state", StatusValid, StatusInvalid),
		Description:     p.String(4, "description"),
	}, p.Err()
}
