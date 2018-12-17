package nmea

const (
	// Prefix GNGNS prefix
	PrefixGNGNS            = "GNGNS"
	NoFixGNGNS             = "N"
	AutonomousGNGNS        = "A"
	DifferentialGNGNS      = "D"
	PreciseGNGNS           = "P"
	RealTimeKinematicGNGNS = "R"
	FloatRTKGNGNS          = "F"
	EstimatedGNGNS         = "E"
	ManualGNGNS            = "M"
	SimulatorGNGNS         = "S"
)

type GNGNS struct {
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

func newGNGNS(s BaseSentence) (GNGNS, error) {
	p := newParser(s, PrefixGNGNS)
	m := GNGNS{
		BaseSentence: s,
		Time:         p.Time(0, "time"),
		Latitude:     p.LatLong(1, 2, "latitude"),
		Longitude:    p.LatLong(3, 4, "longitude"),
		Mode:         p.EnumChars(5, "mode", NoFixGNGNS, AutonomousGNGNS, DifferentialGNGNS, PreciseGNGNS, RealTimeKinematicGNGNS, FloatRTKGNGNS, EstimatedGNGNS, ManualGNGNS, SimulatorGNGNS),
		SVs:          p.Int64(6, "SVs"),
		HDOP:         p.Float64(7, "HDOP"),
		Altitude:     p.Float64(8, "altitude"),
		Separation:   p.Float64(9, "separation"),
		Age:          p.Float64(10, "age"),
		Station:      p.Int64(11, "station"),
	}
	return m, p.Err()
}
