package nmea

const (
	// TypeZDA type for ZDA sentences
	TypeZDA = "ZDA"
)

// ZDA represents date & time data.
// http://aprs.gids.nl/nmea/#zda
// https://gpsd.gitlab.io/gpsd/NMEA.html#_zda_time_date_utc_day_month_year_and_local_time_zone
//
// Format: $--ZDA,hhmmss.ss,xx,xx,xxxx,xx,xx*hh<CR><LF>
// Example: $GPZDA,172809.456,12,07,1996,00,00*57
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
	p := NewParser(s)
	p.AssertType(TypeZDA)
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
