package nmea

const (
	// PrefixGNGGA prefix
	PrefixGNGGA = "GNGGA"
)

type GNGGA struct {
	Sentence
	Time          Time    // Time of fix.
	Latitude      LatLong // Latitude.
	Longitude     LatLong // Longitude.
	FixQuality    string  // Quality of fix.
	NumSatellites string  // Number of satellites in use.
	HDOP          string  // Horizontal dilution of precision.
	Altitude      string  // Altitude.
	Separation    string  // Geoidal separation
	DGPSAge       string  // Age of differential GPD data.
	DGPSId        string  // DGPS reference station ID.
}

func NewGNGGA(s Sentence) (GNGGA, error) {
	p := newParser(s, PrefixGNGGA)
	return GNGGA{
		Sentence:      s,
		Time:          p.Time(0, "time"),
		Latitude:      p.LatLong(1, 2, "latitude"),
		Longitude:     p.LatLong(3, 4, "longitude"),
		FixQuality:    p.EnumString(5, "fix quality", Invalid, GPS, DGPS),
		NumSatellites: p.String(6, "number of satelites"),
		HDOP:          p.String(7, "hdop"),
		Altitude:      p.String(8, "altitude"),
		Separation:    p.String(10, "separation"),
		DGPSAge:       p.String(12, "dgps age"),
		DGPSId:        p.String(13, "dgps id"),
	}, p.Err()
}

func (s GNGGA) GetSentence() Sentence {
	return s.Sentence
}
