package nmea

const (
	// TypeHBT type of HBT sentence for heartbeat supervision sentence.
	TypeHBT = "HBT"
)

// HBT is heartbeat supervision sentence to indicate if equipment is operating normally.
// https://fcc.report/FCC-ID/ADB9ZWRTR100/2768717.pdf (page 1) FURUNO MARINE RADAR, model FAR-15XX manual
//
// Format: $--HBT,x.x,A,x*hh<CR><LF>
// Example: $HCHBT,98.3,0.0,E,12.6,W*57
type HBT struct {
	BaseSentence
	// Interval is configured repeat interval in seconds (1 - 999, null)
	Interval float64
	// OperationStatus is equipment operation status: A = ok, V = not ok
	OperationStatus string
	// MessageID is sequential message identifier (0 - 9). Counts to 9 and resets to 0.
	MessageID int64
}

// newHBT constructor
func newHBT(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeHBT)
	m := HBT{
		BaseSentence:    s,
		Interval:        p.Float64(0, "interval"),
		OperationStatus: p.EnumString(1, "operation status", StatusValid, StatusInvalid),
		MessageID:       p.Int64(2, "message ID"),
	}
	return m, p.Err()
}
