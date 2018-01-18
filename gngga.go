package nmea

const (
	// PrefixGNGGA prefix
	PrefixGNGGA = "GNGGA"
)

type GNGGA struct {
	Sentence
	// Time of fix.
	Time Time
	// Latitude.
	Latitude LatLong
	// Longitude.
	Longitude LatLong
	// Quality of fix.
	FixQuality string
	// Number of satellites in use.
	NumSatellites string
	// Horizontal dilution of precision.
	HDOP string
	// Altitude.
	Altitude string
	// Geoidal separation
	Separation string
	// Age of differential GPD data.
	DGPSAge string
	// DGPS reference station ID.
	DGPSId string
}

func NewGNGGA(sentence Sentence) (GNGGA, error) {
	p := newParser(sentence, PrefixGNGGA)
	return GNGGA{
		Sentence:      sentence,
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
