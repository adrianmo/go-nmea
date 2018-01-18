package nmea

const (
	// PrefixGPGGA prefix
	PrefixGPGGA = "GPGGA"
	// Invalid fix quality.
	Invalid = "0"
	// GPS fix quality
	GPS = "1"
	// DGPS fix quality
	DGPS = "2"
)

// GPGGA represents fix data.
// http://aprs.gids.nl/nmea/#gga
type GPGGA struct {
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

// NewGPGGA parses the GPGGA sentence into this struct.
// e.g: $GPGGA,034225.077,3356.4650,S,15124.5567,E,1,03,9.7,-25.0,M,21.0,M,,0000*58
func NewGPGGA(sentence Sentence) (GPGGA, error) {
	p := newParser(sentence, PrefixGPGGA)
	return GPGGA{
		Sentence:      sentence,
		Time:          p.Time(0, "time"),
		Latitude:      p.LatLong(1, 2, "latitude"),
		Longitude:     p.LatLong(3, 4, "longitude"),
		FixQuality:    p.EnumString(5, "fix quality", Invalid, GPS, DGPS),
		NumSatellites: p.String(6, "number of satelites"),
		HDOP:          p.String(7, "hdap"),
		Altitude:      p.String(8, "altitude"),
		Separation:    p.String(10, "separation"),
		DGPSAge:       p.String(12, "dgps age"),
		DGPSId:        p.String(13, "dgps id"),
	}, p.Err()
}

// GetSentence getter
func (s GPGGA) GetSentence() Sentence {
	return s.Sentence
}
