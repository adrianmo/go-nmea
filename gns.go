package nmea

const (
	// TypeGNS prefix
	TypeGNS = "GNS"
	// NoFixGNS Character
	NoFixGNS = "N"
	// AutonomousGNS Character
	AutonomousGNS = "A"
	// DifferentialGNS Character
	DifferentialGNS = "D"
	// PreciseGNS Character
	PreciseGNS = "P"
	// RealTimeKinematicGNS Character
	RealTimeKinematicGNS = "R"
	// FloatRTKGNS RealTime Kinematic Character
	FloatRTKGNS = "F"
	// EstimatedGNS Fix Character
	EstimatedGNS = "E"
	// ManualGNS Fix Character
	ManualGNS = "M"
	// SimulatorGNS Character
	SimulatorGNS = "S"
)

// GNS is standard GNSS sentance that combined multiple constellations
type GNS struct {
	BaseSentence
	Time       Time
	Latitude   float64
	Longitude  float64
	Mode       []string
	SVs        int64
	HDOP       float64
	Altitude   float64
	Separation float64
	Age        float64
	Station    int64
}

// newGNS Constructor
func newGNS(s BaseSentence) (GNS, error) {
	p := newParser(s)
	p.AssertType(TypeGNS)
	p.AssertTalker("GN")
	m := GNS{
		BaseSentence: s,
		Time:         p.Time(0, "time"),
		Latitude:     p.LatLong(1, 2, "latitude"),
		Longitude:    p.LatLong(3, 4, "longitude"),
		Mode:         p.EnumChars(5, "mode", NoFixGNS, AutonomousGNS, DifferentialGNS, PreciseGNS, RealTimeKinematicGNS, FloatRTKGNS, EstimatedGNS, ManualGNS, SimulatorGNS),
		SVs:          p.Int64(6, "SVs"),
		HDOP:         p.Float64(7, "HDOP"),
		Altitude:     p.Float64(8, "altitude"),
		Separation:   p.Float64(9, "separation"),
		Age:          p.Float64(10, "age"),
		Station:      p.Int64(11, "station"),
	}
	return m, p.Err()
}
