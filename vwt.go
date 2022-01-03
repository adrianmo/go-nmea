package nmea

const (
	// TypeVWT type of VWT sentence for True Wind Speed and Angle
	TypeVWT = "VWT"
)

// VWT - True Wind Speed and Angle
// https://www.nmea.org/Assets/100108_nmea_0183_sentences_not_recommended_for_new_designs.pdf
// https://www.rubydoc.info/gems/nmea_plus/1.0.20/NMEAPlus/Message/NMEA/VWT
// https://lists.gnu.org/archive/html/gpsd-dev/2012-04/msg00048.html
//
// Format: $--VWT,x.x,a,x.x,N,x.x,M,x.x,K*hh<CR><LF>
// Example: $IIVWT,75,x,1.0,N,0.51,M,1.85,K*40
type VWT struct {
	BaseSentence
	TrueAngle        float64 // true Wind direction magnitude in degrees (0 to 180 deg)
	TrueDirectionBow string  // true Wind direction Left/Right of bow
	SpeedKnots       float64 // true wind Speed, knots
	SpeedKnotsUnit   string  // N = knots
	SpeedMPS         float64 // Wind speed, meters/second
	SpeedMPSUnit     string  // M = m/s
	SpeedKPH         float64 // Wind speed, km/hour
	SpeedKPHUnit     string  // M = km/h
}

// newVWT constructor
func newVWT(s BaseSentence) (VWT, error) {
	p := NewParser(s)
	p.AssertType(TypeVWT)
	return VWT{
		BaseSentence:     s,
		TrueAngle:        p.Float64(0, "true wind angle"),
		TrueDirectionBow: p.EnumString(1, "true wind direction to bow", Left, Right),
		SpeedKnots:       p.Float64(2, "wind speed in knots"),
		SpeedKnotsUnit:   p.EnumString(3, "wind speed in knots unit", SpeedKnots),
		SpeedMPS:         p.Float64(4, "wind speed in meters per second"),
		SpeedMPSUnit:     p.EnumString(5, "wind speed in meters per second unit", SpeedMeterPerSecond),
		SpeedKPH:         p.Float64(6, "wind speed in kilometers per hour"),
		SpeedKPHUnit:     p.EnumString(7, "wind speed in kilometers per hour unit", SpeedKilometerPerHour),
	}, p.Err()
}
