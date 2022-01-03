package nmea

const (
	// TypeVWR type of VWR sentence for Relative Wind Speed and Angle
	TypeVWR = "VWR"
)

// VWR - Relative Wind Speed and Angle. Speed is measured relative to the moving vessel.
// According to NMEA: use of $--MWV is recommended.
// https://gpsd.gitlab.io/gpsd/NMEA.html#_vwr_relative_wind_speed_and_angle
// https://www.nmea.org/Assets/100108_nmea_0183_sentences_not_recommended_for_new_designs.pdf (page 16)
//
// Format: $--VWR,x.x,a,x.x,N,x.x,M,x.x,K*hh<CR><LF>
// Example: $IIVWR,75,R,1.0,N,0.51,M,1.85,K*6C
// 			$IIVWR,024,L,018,N,,,,*5e
//			$IIVWR,,,,,,,,*53
type VWR struct {
	BaseSentence
	MeasuredAngle        float64 // Measured Wind direction magnitude in degrees (0 to 180 deg)
	MeasuredDirectionBow string  // Measured Wind direction Left/Right of bow
	SpeedKnots           float64 // Measured wind Speed, knots
	SpeedKnotsUnit       string  // N = knots
	SpeedMPS             float64 // Wind speed, meters/second
	SpeedMPSUnit         string  // M = m/s
	SpeedKPH             float64 // Wind speed, km/hour
	SpeedKPHUnit         string  // M = km/h
}

// newVWR constructor
func newVWR(s BaseSentence) (VWR, error) {
	p := NewParser(s)
	p.AssertType(TypeVWR)
	return VWR{
		BaseSentence:         s,
		MeasuredAngle:        p.Float64(0, "measured wind angle"),
		MeasuredDirectionBow: p.EnumString(1, "measured wind direction to bow", Left, Right),
		SpeedKnots:           p.Float64(2, "wind speed in knots"),
		SpeedKnotsUnit:       p.EnumString(3, "wind speed in knots unit", SpeedKnots),
		SpeedMPS:             p.Float64(4, "wind speed in meters per second"),
		SpeedMPSUnit:         p.EnumString(5, "wind speed in meters per second unit", SpeedMeterPerSecond),
		SpeedKPH:             p.Float64(6, "wind speed in kilometers per hour"),
		SpeedKPHUnit:         p.EnumString(7, "wind speed in kilometers per hour unit", SpeedKilometerPerHour),
	}, p.Err()
}
