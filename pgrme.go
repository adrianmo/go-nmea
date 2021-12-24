package nmea

const (
	// TypePGRME type for PGRME sentences
	TypePGRME = "GRME"
	// ErrorUnit must be meters (M)
	ErrorUnit = "M"
)

// PGRME is Estimated Position Error (Garmin proprietary sentence)
// http://aprs.gids.nl/nmea/#rme
// https://gpsd.gitlab.io/gpsd/NMEA.html#_pgrme_garmin_estimated_error
//
// Format: $PGRME,hhh,M,vvv,M,ttt,M*hh<CR><LF>
// Example: $PGRME,3.3,M,4.9,M,6.0,M*25
type PGRME struct {
	BaseSentence
	Horizontal float64 // Estimated horizontal position error (HPE) in metres
	Vertical   float64 // Estimated vertical position error (VPE) in metres
	Spherical  float64 // Overall spherical equivalent position error in meters
}

// newPGRME constructor
func newPGRME(s BaseSentence) (PGRME, error) {
	p := NewParser(s)
	p.AssertType(TypePGRME)

	horizontal := p.Float64(0, "horizontal error")
	_ = p.EnumString(1, "horizontal error unit", ErrorUnit)

	vertial := p.Float64(2, "vertical error")
	_ = p.EnumString(3, "vertical error unit", ErrorUnit)

	spherical := p.Float64(4, "spherical error")
	_ = p.EnumString(5, "spherical error unit", ErrorUnit)

	return PGRME{
		BaseSentence: s,
		Horizontal:   horizontal,
		Vertical:     vertial,
		Spherical:    spherical,
	}, p.Err()
}
