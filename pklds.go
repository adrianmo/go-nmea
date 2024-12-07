package nmea

import (
	"strings"
)

const (
	// TypePKLDS type for PKLDS sentences
	TypePKLDS = "KLDS"
)

// PKLDS is Kenwood propirtary sentance it is RMC with the addition of Fleetsync ID and status information.
// http://aprs.gids.nl/nmea/#rmc
//
// Format:          $PKLDS,hhmmss.ss,A,ddmm.mm,a,dddmm.mm,a,x.x,x.x,xxxx,x.x,a*hh<CR><LF>
// Example: $PKLDS,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W00,100,2000,25,00*6E
type PKLDS struct {
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
	Fleet		string	// 100 to 349
	ID		string	// 1000 to 4999
	Status		string	// 10 to 99
	Extension	string	// 00 to 99
}

// newPKLDS constructor
func newPKLDS(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypePKLDS)
	m := PKLDS{
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
		Fleet:		 p.String(11, "fleet, range of 100 to 349"),
		ID:              p.String(12, "subscriber unit id, range of 1000 to 4999"),
		Status:          p.String(13, "subscriber unit status id, range of 10 to 99"),
		Extension:       p.String(14, "reserved for future use, range of 00 to 99"),
	}
        if strings.HasPrefix(m.SentanceVersion, "W") == true {
		m.Variation = 0 - m.Variation
	}
        m.SentanceVersion = strings.TrimPrefix(strings.TrimPrefix(m.SentanceVersion, "W"), "E")
	return m, p.Err()
}
