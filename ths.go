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
	// InvalidTHS not valid THS heading (or standby)
	InvalidTHS = "V"
)

// THS is the Actual vessel heading in degrees True with status.
// http://www.nuovamarea.net/pytheas_9.html
// http://manuals.spectracom.com/VSP/Content/VSP/NMEA_THSmess.htm
//
// Format: $--THS,xxx.xx,c*hh<CR><LF>
// Example: $GPTHS,338.01,A*36
type THS struct {
	BaseSentence
	Heading float64 // Heading in degrees
	Status  string  // Heading status
}

// newTHS constructor
func newTHS(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeTHS)
	m := THS{
		BaseSentence: s,
		Heading:      p.Float64(0, "heading"),
		Status:       p.EnumString(1, "status", AutonomousTHS, EstimatedTHS, ManualTHS, SimulatorTHS, InvalidTHS),
	}
	return m, p.Err()
}
