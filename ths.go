package nmea

const (
	// TypeTHS type for THS sentences
	TypeTHS = "THS"
	// AutonomousTHS autonomous ths heading
	AutonomousTHS = "A"
	// EstimatedTHS estimated (dead reckoning) THS heading
	EstimatedTHS = "E"
	// ManualTHS manual input THS heading
	ManualTHS = "M"
	// SimulatorTHS simulated THS heading
	SimulatorTHS = "S"
	// NotValidTHS not valid THS heading (or standby)
	NotValidTHS = "V"
)

// THS is the Actual vessel heading in degrees True with status.
// http://www.nuovamarea.net/pytheas_9.html
type THS struct {
	BaseSentence
	Heading float64 // Heading in degrees
	Status  string  // Heading status
}

// newTHS constructor
func newTHS(s BaseSentence) (THS, error) {
	p := newParser(s)
	p.AssertType(TypeTHS)
	m := THS{
		BaseSentence: s,
		Heading:      p.Float64(0, "heading"),
		Status:       p.EnumString(1, "status", AutonomousTHS, EstimatedTHS, ManualTHS, SimulatorTHS, NotValidTHS),
	}
	return m, p.Err()
}
