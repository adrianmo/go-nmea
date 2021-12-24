package nmea

const (
	// TypeRMC type for RMC sentences
	TypeRMC = "RMC"
	// ValidRMC character
	ValidRMC = "A"
	// InvalidRMC character
	InvalidRMC = "V"
)

// RMC is the Recommended Minimum Specific GNSS data.
// http://aprs.gids.nl/nmea/#rmc
// https://gpsd.gitlab.io/gpsd/NMEA.html#_rmc_recommended_minimum_navigation_information
//
// Format:          $--RMC,hhmmss.ss,A,ddmm.mm,a,dddmm.mm,a,x.x,x.x,xxxx,x.x,a*hh<CR><LF>
// Format NMEA 2.3: $--RMC,hhmmss.ss,A,ddmm.mm,a,dddmm.mm,a,x.x,x.x,xxxx,x.x,a,m*hh<CR><LF>
// Format NMEA 4.1: $--RMC,hhmmss.ss,A,ddmm.mm,a,dddmm.mm,a,x.x,x.x,xxxx,x.x,a,m,s*hh<CR><LF>
// Example: $GNRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*6E
//          $GNRMC,142754.0,A,4302.539570,N,07920.379823,W,0.0,,070617,0.0,E,A*21
//          $GNRMC,102014.00,A,5550.6082,N,03732.2488,E,000.00000,092.9,300518,,,A,V*3B
type RMC struct {
	BaseSentence
	Time      Time    // Time Stamp
	Validity  string  // validity - A-ok, V-invalid
	Latitude  float64 // Latitude
	Longitude float64 // Longitude
	Speed     float64 // Speed in knots
	Course    float64 // True course
	Date      Date    // Date
	Variation float64 // Magnetic variation
	FFAMode   string  // FAA mode indicator (filled in NMEA 2.3 and later)
	NavStatus string  // Nav Status (NMEA 4.1 and later)
}

// newRMC constructor
func newRMC(s BaseSentence) (RMC, error) {
	p := NewParser(s)
	p.AssertType(TypeRMC)
	m := RMC{
		BaseSentence: s,
		Time:         p.Time(0, "time"),
		Validity:     p.EnumString(1, "validity", ValidRMC, InvalidRMC),
		Latitude:     p.LatLong(2, 3, "latitude"),
		Longitude:    p.LatLong(4, 5, "longitude"),
		Speed:        p.Float64(6, "speed"),
		Course:       p.Float64(7, "course"),
		Date:         p.Date(8, "date"),
		Variation:    p.Float64(9, "variation"),
	}
	if p.EnumString(10, "direction", West, East) == West {
		m.Variation = 0 - m.Variation
	}
	if len(p.Fields) > 11 {
		m.FFAMode = p.String(11, "FAA mode") // not enum because some devices have proprietary "non-nmea" values
	}
	if len(p.Fields) > 12 {
		m.NavStatus = p.EnumString(
			12,
			"navigation status",
			NavStatusAutonomous,
			NavStatusDifferential,
			NavStatusEstimated,
			NavStatusManualInput,
			NavStatusSimulated,
			NavStatusDataNotValid,
			NavStatusDataValid,
		)
	}
	return m, p.Err()
}
