package nmea

const (
	// TypeRSA type of RSA sentence for Rudder Sensor Angle
	TypeRSA = "RSA"
)

// RSA - Rudder Sensor Angle
// https://gpsd.gitlab.io/gpsd/NMEA.html#_rsa_rudder_sensor_angle
//
// Format: $--RSA,x.x,A,x.x,A*hh<CR><LF>
// Example: $IIRSA,10.5,A,,V*4D
type RSA struct {
	BaseSentence
	StarboardRudderAngle       float64 // Starboard (or single) rudder sensor, "-" means Turn To Port
	StarboardRudderAngleStatus string  // Status, A = valid, V = Invalid
	PortRudderAngle            float64 // Port rudder sensor
	PortRudderAngleStatus      string  // Status, A = valid, V = Invalid
}

// newRSA constructor
func newRSA(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeRSA)
	return RSA{
		BaseSentence:               s,
		StarboardRudderAngle:       p.Float64(0, "starboard rudder angle"),
		StarboardRudderAngleStatus: p.EnumString(1, "starboard rudder angle status", StatusValid, StatusInvalid),
		PortRudderAngle:            p.Float64(2, "port rudder angle"),
		PortRudderAngleStatus:      p.EnumString(3, "port rudder angle status", StatusValid, StatusInvalid),
	}, p.Err()
}
