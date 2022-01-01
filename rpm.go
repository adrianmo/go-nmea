package nmea

const (
	// TypeRPM type of RPM sentence for Engine or Shaft revolutions and pitch
	TypeRPM = "RPM"

	// SourceEngineRPM is value for case when source is Engine
	SourceEngineRPM = "E"
	// SourceShaftRPM is value for case when source is Shaft
	SourceShaftRPM = "S"
)

// RPM - Engine or Shaft revolutions and pitch
// https://gpsd.gitlab.io/gpsd/NMEA.html#_rpm_revolutions
//
// Format: $--RPM,a,x,x.x,x.x,A*hh<CR><LF>
// Example: $RCRPM,S,0,74.6,30.0,A*56
type RPM struct {
	BaseSentence
	Source       string  // Source, S = Shaft, E = Engine
	EngineNumber int64   // Engine or shaft number
	SpeedRPM     float64 // Speed, Revolutions per minute
	PitchPercent float64 // Propeller pitch, % of maximum, "-" means astern
	Status       string  // Status, A = Valid, V = Invalid
}

// newRPM constructor
func newRPM(s BaseSentence) (RPM, error) {
	p := NewParser(s)
	p.AssertType(TypeRPM)
	return RPM{
		BaseSentence: s,
		Source:       p.EnumString(0, "source", SourceEngineRPM, SourceShaftRPM),
		EngineNumber: p.Int64(1, "engine number"),
		SpeedRPM:     p.Float64(2, "speed"),
		PitchPercent: p.Float64(3, "pitch"),
		Status:       p.EnumString(4, "status", StatusValid, StatusInvalid),
	}, p.Err()
}
