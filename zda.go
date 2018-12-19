package nmea

const (
	// PrefixZDA prefix
	PrefixZDA = "ZDA"
)

// ZDA represents date & time data.
// http://aprs.gids.nl/nmea/#zda
type ZDA struct {
	BaseSentence
	Time          Time
	Day           int64
	Month         int64
	Year          int64
	OffsetHours   int64 // Local time zone offset from GMT, hours
	OffsetMinutes int64 // Local time zone offset from GMT, minutes
}

// newZDA constructor
func newZDA(s BaseSentence) (ZDA, error) {
	p := newParser(s, PrefixZDA)
	p.AssertType(PrefixZDA)
	p.AssertTalker("GP")
	return ZDA{
		BaseSentence:  s,
		Time:          p.Time(0, "time"),
		Day:           p.Int64(1, "day"),
		Month:         p.Int64(2, "month"),
		Year:          p.Int64(3, "year"),
		OffsetHours:   p.Int64(4, "offset (hours)"),
		OffsetMinutes: p.Int64(5, "offset (minutes)"),
	}, p.Err()
}
