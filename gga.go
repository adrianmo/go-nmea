package nmea

const (
	// TypeGGA type for GGA sentences
	TypeGGA = "GGA"
	// Invalid fix quality.
	Invalid = "0"
	// GPS fix quality
	GPS = "1"
	// DGPS fix quality
	DGPS = "2"
	// PPS fix
	PPS = "3"
	// RTK real time kinematic fix
	RTK = "4"
	// FRTK float RTK fix
	FRTK = "5"
	// EST estimated fix.
	EST = "6"
)

// GGA is the Time, position, and fix related data of the receiver.
type GGA struct {
	BaseSentence
	Time          Time    // Time of fix.
	Latitude      float64 // Latitude.
	Longitude     float64 // Longitude.
	FixQuality    string  // Quality of fix.
	NumSatellites int64   // Number of satellites in use.
	HDOP          float64 // Horizontal dilution of precision.
	Altitude      float64 // Altitude.
	Separation    float64 // Geoidal separation
	DGPSAge       string  // Age of differential GPD data.
	DGPSId        string  // DGPS reference station ID.
}

// newGGA constructor
func newGGA(s BaseSentence) (GGA, error) {
	p := NewParser(s)
	p.AssertType(TypeGGA)
	return GGA{
		BaseSentence:  s,
		Time:          p.Time(0, "time"),
		Latitude:      p.LatLong(1, 2, "latitude"),
		Longitude:     p.LatLong(3, 4, "longitude"),
		FixQuality:    p.EnumString(5, "fix quality", Invalid, GPS, DGPS, PPS, RTK, FRTK, EST),
		NumSatellites: p.Int64(6, "number of satellites"),
		HDOP:          p.Float64(7, "hdop"),
		Altitude:      p.Float64(8, "altitude"),
		Separation:    p.Float64(10, "separation"),
		DGPSAge:       p.String(12, "dgps age"),
		DGPSId:        p.String(13, "dgps id"),
	}, p.Err()
}
