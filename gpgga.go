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

// NewGPGGA constructor
func NewGPGGA(sentence Sentence) GPGGA {
	s := new(GPGGA)
	s.Sentence = sentence
	return *s
}

// GetSentence getter
func (s GPGGA) GetSentence() Sentence {
	return s.Sentence
}

// Parse parses the GPGGA sentence into this struct.
// e.g: $GPGGA,034225.077,3356.4650,S,15124.5567,E,1,03,9.7,-25.0,M,21.0,M,,0000*58
func (s *GPGGA) parse() error {
	p := newParser(s.Sentence, PrefixGPGGA)

	s.Time = p.Time(0, "time")
	s.Latitude = p.LatLong(1, 2, "latitude")
	s.Longitude = p.LatLong(3, 4, "longitude")
	s.FixQuality = p.String(5, "fix quality")
	if s.FixQuality != Invalid && s.FixQuality != GPS && s.FixQuality != DGPS {
		p.SetErr("fix quality", s.FixQuality)
	}
	s.NumSatellites = p.String(6, "number of satelites")
	s.HDOP = p.String(7, "hdap")
	s.Altitude = p.String(8, "altitude")
	s.Separation = p.String(10, "separation")
	s.DGPSAge = p.String(12, "dgps age")
	s.DGPSId = p.String(13, "dgps id")

	return p.Err()
}
