package nmea

/**
$WIMWV

NMEA 0183 standard Wind Speed and Angle, in relation to the vessel’s bow/centerline.

Syntax
$WIMWV,<1>,<2>,<3>,<4>,<5>*hh<CR><LF>

Fields
<1>    Wind angle, 0.0 to 359.9 degrees, in relation to the vessel’s bow/centerline, to the nearest 0.1
           degree. If the data for this field is not valid, the field will be blank.
<2>    Reference:
           R = Relative (apparent wind, as felt when standing on the moving ship)
           T = Theoretical (calculated actual wind, as though the vessel were stationary)
<3>    Wind speed, to the nearest tenth of a unit.  If the data for this field is not valid, the field will be
           blank.
<4>    Wind speed units:
           K = km/hr
           M = m/s
           N = knots
           S = statute miles/hr
           (Most WeatherStations will commonly use "N" (knots))
<5>    Status:
           A = data valid; V = data invalid
*/

const (
	// TypeMWV type for MWV sentences
	TypeMWV = "MWV"
	// RelativeMWV for Valid Relative angle data
	RelativeMWV = "R"
	// TheoreticalMWV for valid Theoretical angle data
	TheoreticalMWV = "T"
	// UnitKMHMWV unit for Kilometer per hour (KM/H)
	UnitKMHMWV = "K" // KM/H
	// UnitMSMWV unit for Meters per second (M/S)
	UnitMSMWV = "M" // M/S
	// UnitKnotsMWV unit for knots
	UnitKnotsMWV = "N" // knots
	// UnitSMilesHMWV unit for Miles per hour (M/H)
	UnitSMilesHMWV = "S"
	// ValidMWV data is valid
	ValidMWV = "A"
	// InvalidMWV data is invalid
	InvalidMWV = "V"
)

// MWV is the Wind Speed and Angle, in relation to the vessel’s bow/centerline.
// https://gpsd.gitlab.io/gpsd/NMEA.html#_mwv_wind_speed_and_angle
//
// Format: $--MWV,x.x,a,x.x,a*hh<CR><LF>
// Example: $WIMWV,12.1,T,10.1,N,A*27
type MWV struct {
	BaseSentence
	WindAngle     float64
	Reference     string
	WindSpeed     float64
	WindSpeedUnit string
	StatusValid   bool
}

func newMWV(s BaseSentence) (MWV, error) {
	p := NewParser(s)
	p.AssertType(TypeMWV)
	return MWV{
		BaseSentence:  s,
		WindAngle:     p.Float64(0, "wind angle"),
		Reference:     p.EnumString(1, "reference", RelativeMWV, TheoreticalMWV),
		WindSpeed:     p.Float64(2, "wind speed"),
		WindSpeedUnit: p.EnumString(3, "wind speed unit", UnitKMHMWV, UnitMSMWV, UnitKnotsMWV, UnitSMilesHMWV),
		StatusValid:   p.EnumString(4, "status", ValidMWV, InvalidMWV) == ValidMWV,
	}, p.Err()
}
