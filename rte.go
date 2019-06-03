package nmea

const (
	// TypeRTE type for RTE sentences
	TypeRTE = "RTE"

	// ActiveRoute active route
	ActiveRoute = "c"

	// WaypointList list containing waypoints
	WaypointList = "w"
)

// RTE is a route of waypoints
type RTE struct {
	BaseSentence
	NumberOfSentences         int64    // Number of sentences in sequence
	SentenceNumber            int64    // Sentence number
	ActiveRouteOrWaypointList string   // Current active route or waypoint list
	Name                      string   // Name or number of active route
	Idents                    []string // List of ident of waypoints
}

// newRTE constructor
func newRTE(s BaseSentence) (RTE, error) {
	p := newParser(s)
	p.AssertType(TypeRTE)
	return RTE{
		BaseSentence:              s,
		NumberOfSentences:         p.Int64(0, "number of sentences"),
		SentenceNumber:            p.Int64(1, "sentence number"),
		ActiveRouteOrWaypointList: p.EnumString(2, "active route or waypoint list", ActiveRoute, WaypointList),
		Name:                      p.String(3, "name or number"),
		Idents:                    p.ListString(4, "ident of waypoints"),
	}, p.Err()
}
