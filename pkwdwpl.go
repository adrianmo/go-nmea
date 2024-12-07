package nmea

const (
	// TypePKWDWPL type for PKLDS sentences
	TypePKWDWPL = "KWDWPL"
)

// PKWDWPL is Kenwood Waypoint Location
// https://github.com/wb2osz/direwolf/blob/master/waypoint.c
//
// Format:  $PKWDWPL,hhmmss,v,ddmm.mm,ns,dddmm.mm,ew,speed,course,ddmmyy,alt,wname,ts*hh<CR><LF>
// Example: $PKWDWPL,204714,V,4237.1400,N,07120.8300,W,,,200316,,test|5,/'*61
type PKWDWPL struct {
	BaseSentence
	Time         Time    // Time Stamp
	Validity     string  // validity - A-ok, V-invalid
	Latitude     float64 // Latitude
	Longitude    float64 // Longitude
	Speed        float64 // Speed in knots
	Course       float64 // True course
	Date         Date    // Date
	Altitude     float64 // Magnetic variation
	WaypointName string  // 00 to 15
	TableSymbol  string  // U00001 to U65519 or U00000001 to U16776415 (U is FIXED)
}

// newPKWDWPL constructor
func newPKWDWPL(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypePKWDWPL)
	m := PKWDWPL{
		BaseSentence: s,
		Time:         p.Time(0, "time"),
		Validity:     p.EnumString(1, "validity", ValidRMC, InvalidRMC),
		Latitude:     p.LatLong(2, 3, "latitude"),
		Longitude:    p.LatLong(4, 5, "longitude"),
		Speed:        p.Float64(6, "speed"),
		Course:       p.Float64(7, "course"),
		Date:         p.Date(8, "date"),
		Altitude:     p.Float64(9, "altitude"),
		WaypointName: p.String(10, "waypoint name, Object name/Sendin Station"),
		TableSymbol:  p.String(11, "table and symbol as per APRS spec"),
	}
	return m, p.Err()
}
