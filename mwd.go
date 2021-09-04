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
	TypeMWD         = "MWD"
	TrueMWD         = "T"
	MagneticMWD     = "M"
	KnotsMWD        = "N"
	MetersSecondMWD = "M"
	EmptyMWD        = ""
)

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

func newMWD(s BaseSentence) (MWD, error) {
	p := NewParser(s)
	p.AssertType(TypeMWD)
	return MWD{
		BaseSentence:          s,
		WindDirectionTrue:     p.Float64(0, "true wind direction"),
		TrueValid:             p.EnumString(1, "true wind valid", TrueMWD, EmptyMWD) == TrueMWD,
		WindDirectionMagnetic: p.Float64(2, "magnetic wind direction"),
		MagneticValid:         p.EnumString(3, "magnetic direction valid", MagneticMWD, EmptyMWD) == MagneticMWD,
		WindSpeedKnots:        p.Float64(4, "windspeed knots"),
		KnotsValid:            p.EnumString(5, "windspeed knots valid", KnotsMWD, EmptyMWD) == KnotsMWD,
		WindSpeedMeters:       p.Float64(6, "windspeed m/s"),
		MetersValid:           p.EnumString(7, "windspeed m/s valid", MetersSecondMWD, EmptyMWD) == MetersSecondMWD,
	}, p.Err()
}
