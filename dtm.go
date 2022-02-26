package nmea

const (
	// TypeDTM type of DTM sentence for Datum Reference
	TypeDTM = "DTM"
)

// DTM - Datum Reference
// https://gpsd.gitlab.io/gpsd/NMEA.html#_dtm_datum_reference
//
// Format: $--DTM,ref,x,llll,c,llll,c,aaa,ref*hh<CR><LF>
// Example: $GPDTM,W84,,0.0,N,0.0,E,0.0,W84*6F
// Example: $GPDTM,W84,,00.0000,N,00.0000,W,,W84*53
type DTM struct {
	BaseSentence
	LocalDatumCode    string // Local datum code (W84,W72,S85,P90,999)
	LocalDatumSubcode string // Local datum subcode. May be blank.

	LatitudeOffsetMinute  float64 // Latitude offset (minutes) (negative if south)
	LongitudeOffsetMinute float64 // Longitude offset (minutes) (negative if west)

	AltitudeOffsetMeters float64 // Altitude offset in meters
	DatumName            string  // Reference datum name. What’s usually seen here is "W84", the standard WGS84 datum used by GPS.
}

// newDTM constructor
func newDTM(s BaseSentence) (DTM, error) {
	p := NewParser(s)
	p.AssertType(TypeDTM)
	m := DTM{
		BaseSentence:      s,
		LocalDatumCode:    p.String(0, "local datum code"),
		LocalDatumSubcode: p.String(1, "local datum subcode"),

		LatitudeOffsetMinute:  p.Float64(2, "latitude offset minutes"),
		LongitudeOffsetMinute: p.Float64(4, "longitude offset minutes"),

		AltitudeOffsetMeters: p.Float64(6, "altitude offset offset"),
		DatumName:            p.String(7, "datum name"),
	}
	if p.String(3, "latitude offset direction") == South {
		m.LatitudeOffsetMinute *= -1
	}
	if p.String(5, "longitude offset direction") == West {
		m.LongitudeOffsetMinute *= -1
	}
	return m, p.Err()
}
