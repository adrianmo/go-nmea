package nmea

const (
	// TypeVPW type of VPW sentence for Speed Measured Parallel to Wind
	TypeVPW = "VPW"
)

// VPW - Speed Measured Parallel to Wind
// https://gpsd.gitlab.io/gpsd/NMEA.html#_vpw_speed_measured_parallel_to_wind
//
// Format: $--VPW,x.x,N,x.x,M*hh<CR><LF>
// Example: $IIVPW,4.5,N,6.7,M*52
type VPW struct {
	BaseSentence
	SpeedKnots     float64 // Speed, "-" means downwind, knots
	SpeedKnotsUnit string  // N = knots
	SpeedMPS       float64 // Speed, "-" means downwind, m/s
	SpeedMPSUnit   string  // M = m/s
}

// newVPW constructor
func newVPW(s BaseSentence) (VPW, error) {
	p := NewParser(s)
	p.AssertType(TypeVPW)
	return VPW{
		BaseSentence:   s,
		SpeedKnots:     p.Float64(0, "wind speed in knots"),
		SpeedKnotsUnit: p.EnumString(1, "wind speed in knots unit", SpeedKnots),
		SpeedMPS:       p.Float64(2, "wind speed in meters per second"),
		SpeedMPSUnit:   p.EnumString(3, "wind speed in meters per second unit", SpeedMeterPerSecond),
	}, p.Err()
}
