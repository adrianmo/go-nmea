package nmea

const (
	// TypeGNS type for GNS sentences
	TypeGNS = "GNS"
)

// GNS mode values. These are same values ans GLL/RMC FAAMode* values.
// Note: there can be other values (proprietary).
const (
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
// https://gpsd.gitlab.io/gpsd/NMEA.html#_gns_fix_data
//
// Format: $--GNS,hhmmss.ss,ddmm.mm,a,dddmm.mm,a,c--c,xx,x.x,x.x,x.x,x.x,x.x*hh<CR><LF>
// Example: $GNGNS,014035.00,4332.69262,S,17235.48549,E,RR,13,0.9,25.63,11.24,,*70
//          $GPGNS,224749.00,3333.4268304,N,11153.3538273,W,D,19,0.6,406.110,-26.294,6.0,0138,S*6A
type GNS struct {
	BaseSentence
	Time      Time // UTC of position
	Latitude  float64
	Longitude float64
	// FAA mode indicator for each satellite navigation system (constellation) supported by device.
	//
	// May be up to six characters (according to GPSD).
	// '1' - GPS
	// '2' - GLONASS
	// '3' - Galileo
	// '4' - BDS
	// '5' - QZSS
	// '6' - NavIC (IRNSS)
	Mode       []string
	SVs        int64   // Total number of satellites in use, 00-99
	HDOP       float64 // Horizontal Dilution of Precision
	Altitude   float64 // Antenna altitude, meters, re:mean-sea-level(geoid).
	Separation float64 // Geoidal separation meters
	Age        float64 // Age of differential data
	Station    int64   // Differential reference station ID
	NavStatus  string  // Navigation status (NMEA 4.1+). See NavStats* (`NavStatusAutonomous` etc) constants for possible values.
}

// newGNS Constructor
func newGNS(s BaseSentence) (GNS, error) {
	p := NewParser(s)
	p.AssertType(TypeGNS)
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
	if len(p.Fields) >= 13 {
		m.NavStatus = p.EnumString(
			12,
			"navigation status",
			NavStatusAutonomous,
			NavStatusDifferential,
			NavStatusEstimated,
			NavStatusManualInput,
			NavStatusSimulated,
			NavStatusDataNotValid,
			NavStatusDataValid,
		)
	}
	return m, p.Err()
}
