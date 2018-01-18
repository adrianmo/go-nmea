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
	s := new(GNGGA)
	s.Sentence = sentence
	return *s, s.parse()
}

func (s GNGGA) GetSentence() Sentence {
	return s.Sentence
}

func (s *GNGGA) parse() error {
	p := newParser(s.Sentence, PrefixGNGGA)
	if err := p.Err(); err != nil {
		return err
	}

	s.Time = p.Time(0, "time")
	s.Latitude = p.LatLong(1, 2, "latitude")
	s.Longitude = p.LatLong(3, 4, "longitude")

	s.FixQuality = p.String(5, "fix quality")
	if s.FixQuality != Invalid && s.FixQuality != GPS && s.FixQuality != DGPS {
		p.SetErr("fix quality", s.FixQuality)
	}
	s.NumSatellites = p.String(6, "number of satelites")
	s.HDOP = p.String(7, "hdop")
	s.Altitude = p.String(8, "altitude")
	s.Separation = p.String(10, "separation")
	s.DGPSAge = p.String(12, "dgps age")
	s.DGPSId = p.String(13, "dgps id")
	return p.Err()
}
