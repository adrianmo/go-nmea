package nmea

const (
	// TypeMCP type for MCP sentences
	TypeMCP = "MCP"
)

// MCP - Micropilot Joystick Controller Message
// Format: $--MCP,x.x,f,x.x,f,x.x,f,x.x,f,x.x,f,x,i,x,i,x,i*hh<CR><LF>
// Example: $IIMCP,50.0,%f,30.0,%f,45.0,%f,0,0,0*hh
type MCP struct {
	BaseSentence
	JoystickSurgeAxisCommandSetValue float64 // Joystick surge axis (ahead - astern) command set value.
	JoystickSwayAxisCommandSetValue  float64 // Joystick sway axis (sideways port - sideways starboard) command set value.
	JoystickYawAxisCommandSetValue   float64 // Joystick yaw axis (rotation counter-clockwise - clockwise) command set value.
	Reserved1                        float64 // Reserved for future use.
	Reserved2                        float64 // Reserved for future use.
	ValueErrorStatusWord             int64   // Value error status word for values 1...5, each bit indicates a single value error condition state either OK or ERROR.
	ControlStateWord1                int64   // Control state word 1, each bit indicates a separate condition state either ON or OFF.
	ControlStateWord2                int64   // Control state word 2, each bit indicates a separate condition state either ON or OFF.
}

// newMCP constructor
func newMCP(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeMCP)
	return MCP{
		BaseSentence:                     s,
		JoystickSurgeAxisCommandSetValue: p.Float64(0, "value1"),
		JoystickSwayAxisCommandSetValue:  p.Float64(1, "value2"),
		JoystickYawAxisCommandSetValue:   p.Float64(2, "value3"),
		Reserved1:                        p.Float64(3, "value4"),
		Reserved2:                        p.Float64(4, "value5"),
		ValueErrorStatusWord:             p.Int64(5, "vesw"),
		ControlStateWord1:                p.Int64(6, "csw1"),
		ControlStateWord2:                p.Int64(7, "csw2"),
	}, p.Err()
}
