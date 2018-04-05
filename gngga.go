package nmea

const (
	// PrefixGNGGA prefix
	PrefixGNGGA = "GNGGA"
)

// GNGGA is the Time, position, and fix related data of the receiver.
type GNGGA struct {
	Sent
	Time          Time    // Time of fix.
	Latitude      LatLong // Latitude.
	Longitude     LatLong // Longitude.
	FixQuality    string  // Quality of fix.
	NumSatellites int64   // Number of satellites in use.
	HDOP          float64 // Horizontal dilution of precision.
	Altitude      float64 // Altitude.
	Separation    float64 // Geoidal separation
	DGPSAge       string  // Age of differential GPD data.
	DGPSId        string  // DGPS reference station ID.
}

// NewGNGGA constructor
func NewGNGGA(s Sent) (GNGGA, error) {
	p := newParser(s, PrefixGNGGA)
	return GNGGA{
		Sent:          s,
		Time:          p.Time(0, "time"),
		Latitude:      p.LatLong(1, 2, "latitude"),
		Longitude:     p.LatLong(3, 4, "longitude"),
		FixQuality:    p.EnumString(5, "fix quality", Invalid, GPS, DGPS, PPS, RTK, FRTK),
		NumSatellites: p.Int64(6, "number of satelites"),
		HDOP:          p.Float64(7, "hdop"),
		Altitude:      p.Float64(8, "altitude"),
		Separation:    p.Float64(10, "separation"),
		DGPSAge:       p.String(12, "dgps age"),
		DGPSId:        p.String(13, "dgps id"),
	}, p.Err()
}
