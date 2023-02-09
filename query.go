package nmea

const (
	// TypeQuery type of Query sentence for a listener to request a particular sentence from a talker
	TypeQuery = "Q"
)

// Query sentences is special type of sentence for a listener to request a particular sentence from a talker.
// https://www.tronico.fi/OH6NT/docs/NMEA0183.pdf (page 3)
//
// Format: $ttllQ,sss*hh<CR><LF>
// Example: $CCGPQ,GGA*2B<CR><LF>
type Query struct {
	BaseSentence
	DestinationTalkerID string
	RequestedSentence   string
}

// newQuery constructor
func newQuery(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeQuery)

	return Query{
		BaseSentence:        s,
		DestinationTalkerID: s.Raw[3:5],
		RequestedSentence:   p.String(0, "requested sentence"),
	}, p.Err()
}
