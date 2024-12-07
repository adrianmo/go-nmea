package nmea

import (
	"strings"
)

const (
	// TypePKNDS type for PKLDS sentences
	TypePKNDS = "KNDS"
)

// PKNDS is Kenwood propirtary sentance it is RMC with the addition of NEXTEDGE and status information.
// http://aprs.gids.nl/nmea/#rmc
//
// Format:          $PKNDS,hhmmss.ss,A,ddmm.mm,a,dddmm.mm,a,xxx.x,x,x.x,xxx,Uxxxx,xxx.xx,*hh<CR><LF>
// Example: $PKNDS,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W00,U00001,207,00,*6E
type PKNDS struct {
	BaseSentence
	Time      Time    // Time Stamp
	Validity  string  // validity - A-ok, V-invalid
	Latitude  float64 // Latitude
	Longitude float64 // Longitude
	Speed     float64 // Speed in knots
	Course    float64 // True course
	Date      Date    // Date
	Variation float64 // Magnetic variation
        SentanceVersion	string	// 00 to 15
	UnitID		string	// U00001 to U65519 or U00000001 to U16776415 (U is FIXED)
	Status		string	// 001 to 255
	Extension	string	// 00 to 99
}

// newPKNDS constructor
func newPKNDS(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypePKNDS)
	m := PKNDS{
		BaseSentence: s,
		Time:            p.Time(0, "time"),
		Validity:        p.EnumString(1, "validity", ValidRMC, InvalidRMC),
		Latitude:        p.LatLong(2, 3, "latitude"),
		Longitude:       p.LatLong(4, 5, "longitude"),
		Speed:           p.Float64(6, "speed"),
		Course:          p.Float64(7, "course"),
		Date:            p.Date(8, "date"),
		Variation:       p.Float64(9, "variation"),
		SentanceVersion: p.String(10, "sentance version, range of 00 to 15"),
		UnitID:          p.String(11, "unit ID, NXDN range U00001 to U65519, DMR range of  U00000001 to U16776415"),
		Status:          p.String(12, "subscriber unit status id, range of 001 to 255"),
		Extension:       p.String(13, "reserved for future use, range of 00 to 99"),
	}
        if strings.HasPrefix(m.SentanceVersion, "W") == true {
		m.Variation = 0 - m.Variation
	}
        m.SentanceVersion = strings.TrimPrefix(strings.TrimPrefix(m.SentanceVersion, "W"), "E")
	return m, p.Err()
}
