package nmea

const (
	// PrefixGPZDA prefix
	PrefixGPZDA = "GPZDA"
)

// GPZDA represents date & time data.
// http://aprs.gids.nl/nmea/#zda
type GPZDA struct {
	Sentence
	Time  Time
	Day   int64
	Month int64
	Year  int64
	// Local time zone offset from GMT, hours
	OffsetHours int64
	// Local time zone offset from GMT, minutes
	OffsetMinutes int64
}

// NewGPZDA constructor
func NewGPZDA(sentence Sentence) (GPZDA, error) {
	s := GPZDA{Sentence: sentence}
	p := newParser(sentence, PrefixGPZDA)
	s.Time = p.Time(0, "time")
	s.Day = p.Int64(1, "day")
	s.Month = p.Int64(2, "month")
	s.Year = p.Int64(3, "year")
	s.OffsetHours = p.Int64(4, "offset (hours)")
	s.OffsetMinutes = p.Int64(5, "offset (minutes)")
	return s, p.Err()
}

// GetSentence getter
func (s GPZDA) GetSentence() Sentence {
	return s.Sentence
}
