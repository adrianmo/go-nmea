package nmea

/**
$WIMDA

NMEA 0183 standard Meteorological Composite.

Syntax
$WIMDA,<1>,<2>,<3>,<4>,<5>,<6>,<7>,<8>,<9>,<10>,<11>,
<12>,<13>,<14>,<15>,<16>,<17>,<18>,<19>,<20>*hh
<CR><LF>

Fields
<1> Barometric pressure, inches of mercury, to the nearest 0.01 inch
<2> I = inches of mercury
<3> Barometric pressure, bars, to the nearest .001 bar
<4> B = bars
<5> Air temperature, degrees C, to the nearest 0.1 degree C
<6> C = degrees C
<7> Water temperature, degrees C (this field left blank by some WeatherStations)
<8> C = degrees C (this field left blank by WeatherStation)
<9> Relative humidity, percent, to the nearest 0.1 percent
<10> Absolute humidity, percent (this field left blank by some WeatherStations)
<11> Dew point, degrees C, to the nearest 0.1 degree C
<12> C = degrees C
<13> Wind direction, degrees True, to the nearest 0.1 degree
<14> T = true
<15> Wind direction, degrees Magnetic, to the nearest 0.1 degree
<16> M = magnetic
<17> Wind speed, knots, to the nearest 0.1 knot
<18> N = knots
<19> Wind speed, meters per second, to the nearest 0.1 m/s
<20> M = meters per second
*/

const (
	// TypeMDA type for MDA sentences
	TypeMDA = "MDA"
	// InchMDA for valid pressure in Inches of mercury
	InchMDA = "I"
	// BarsMDA for valid pressure in Bars
	BarsMDA = "B"
	// DegreesCMDA for valid data in degrees C
	DegreesCMDA = "C"
	// TrueMDA for valid data in True direction
	TrueMDA = "T"
	// MagneticMDA for valid data in Magnetic direction
	MagneticMDA = "M"
	// KnotsMDA for valid data in Knots
	KnotsMDA = "N"
	// MetersSecondMDA for valid data in Meters per Second
	MetersSecondMDA = "M"
)

// MDA is the Meteorological Composite
// Data of air pressure, air and water temperatures and wind speed and direction
// https://gpsd.gitlab.io/gpsd/NMEA.html#_mda_meteorological_composite
// https://opencpn.org/wiki/dokuwiki/doku.php?id=opencpn:opencpn_user_manual:advanced_features:nmea_sentences#mda
//
// Format: $--MDA,n.nn,I,n.nnn,B,n.n,C,n.C,n.n,n,n.n,C,n.n,T,n.n,M,n.n,N,n.n,M*hh<CR><LF>
// Example: $WIMDA,3.02,I,1.01,B,23.4,C,,,40.2,,12.1,C,19.3,T,20.1,M,13.1,N,1.1,M*62
type MDA struct {
	BaseSentence
	PressureInch          float64
	InchesValid           bool // I
	PressureBar           float64
	BarsValid             bool // B
	AirTemp               float64
	AirTempValid          bool // C or empty if no data
	WaterTemp             float64
	WaterTempValid        bool    // C or empty if no data
	RelativeHum           float64 // percent to .1
	AbsoluteHum           float64 // percent to .1
	DewPoint              float64
	DewPointValid         bool // C or empty if no data
	WindDirectionTrue     float64
	TrueValid             bool // T
	WindDirectionMagnetic float64
	MagneticValid         bool // M
	WindSpeedKnots        float64
	KnotsValid            bool // N
	WindSpeedMeters       float64
	MetersValid           bool // M
}

func newMDA(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeMDA)
	return MDA{
		BaseSentence:          s,
		PressureInch:          p.Float64(0, "pressure in inch"),
		InchesValid:           p.EnumString(1, "inches valid", InchMDA) == InchMDA,
		PressureBar:           p.Float64(2, "pressure in bar"),
		BarsValid:             p.EnumString(3, "bars valid", BarsMDA) == BarsMDA,
		AirTemp:               p.Float64(4, "air temp"),
		AirTempValid:          p.EnumString(5, "air temp valid", DegreesCMDA) == DegreesCMDA,
		WaterTemp:             p.Float64(6, "water temp"),
		WaterTempValid:        p.EnumString(7, "water temp valid", DegreesCMDA) == DegreesCMDA,
		RelativeHum:           p.Float64(8, "relative humidity"),
		AbsoluteHum:           p.Float64(9, "absolute humidity"),
		DewPoint:              p.Float64(10, "dewpoint"),
		DewPointValid:         p.EnumString(11, "dewpoint valid", DegreesCMDA) == DegreesCMDA,
		WindDirectionTrue:     p.Float64(12, "wind direction true"),
		TrueValid:             p.EnumString(13, "wind direction true valid", TrueMDA) == TrueMDA,
		WindDirectionMagnetic: p.Float64(14, "wind direction magnetic"),
		MagneticValid:         p.EnumString(15, "wind direction magnetic valid", MagneticMDA) == MagneticMDA,
		WindSpeedKnots:        p.Float64(16, "windspeed knots"),
		KnotsValid:            p.EnumString(17, "windspeed knots valid", KnotsMDA) == KnotsMDA,
		WindSpeedMeters:       p.Float64(18, "windspeed m/s"),
		MetersValid:           p.EnumString(19, "windspeed m/s valid", MetersSecondMDA) == MetersSecondMDA,
	}, p.Err()
}
