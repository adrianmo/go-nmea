package nmea

/**
$WIMWD

NMEA 0183 standard Wind Direction and Speed, with respect to north.

Syntax
$WIMWD,<1>,<2>,<3>,<4>,<5>,<6>,<7>,<8>*hh<CR><LF>

Fields
<1>    Wind direction, 0.0 to 359.9 degrees True, to the nearest 0.1 degree
<2>    T = True
<3>    Wind direction, 0.0 to 359.9 degrees Magnetic, to the nearest 0.1 degree
<4>    M = Magnetic
<5>    Wind speed, knots, to the nearest 0.1 knot.
<6>    N = Knots
<7>    Wind speed, meters/second, to the nearest 0.1 m/s.
<8>    M = Meters/second
*/

const (
	// TypeMWD type for MWD sentences
	TypeMWD = "MWD"
	// TrueMWD for valid True Direction
	TrueMWD = "T"
	// MagneticMWD for valid Magnetic direction
	MagneticMWD = "M"
	// KnotsMWD for valid Knots
	KnotsMWD = "N"
	// MetersSecondMWD for valid Meters per Second
	MetersSecondMWD = "M"
)

// MWD Wind Direction and Speed, with respect to north.
// https://www.tronico.fi/OH6NT/docs/NMEA0183.pdf
// http://gillinstruments.com/data/manuals/OMC-140_Operator_Manual_v1.04_131117.pdf
//
// Format: $--MWD,x.x,T,x.x,M,x.x,N,x.x,M*hh<CR><LF>
// Example: $WIMWD,10.1,T,10.1,M,12,N,40,M*5D
type MWD struct {
	BaseSentence
	WindDirectionTrue     float64
	TrueValid             bool
	WindDirectionMagnetic float64
	MagneticValid         bool
	WindSpeedKnots        float64
	KnotsValid            bool
	WindSpeedMeters       float64
	MetersValid           bool
}

func newMWD(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeMWD)
	return MWD{
		BaseSentence:          s,
		WindDirectionTrue:     p.Float64(0, "true wind direction"),
		TrueValid:             p.EnumString(1, "true wind valid", TrueMWD) == TrueMWD,
		WindDirectionMagnetic: p.Float64(2, "magnetic wind direction"),
		MagneticValid:         p.EnumString(3, "magnetic direction valid", MagneticMWD) == MagneticMWD,
		WindSpeedKnots:        p.Float64(4, "windspeed knots"),
		KnotsValid:            p.EnumString(5, "windspeed knots valid", KnotsMWD) == KnotsMWD,
		WindSpeedMeters:       p.Float64(6, "windspeed m/s"),
		MetersValid:           p.EnumString(7, "windspeed m/s valid", MetersSecondMWD) == MetersSecondMWD,
	}, p.Err()
}
