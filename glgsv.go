package nmea

const (
	// PrefixGLGSV prefix
	PrefixGLGSV = "GLGSV"
)

// GLGSV represents the GPS Satellites in view
// http://aprs.gids.nl/nmea/#glgsv
type GLGSV struct {
	Sentence
	TotalMessages   int64 // Total number of messages of this type in this cycle
	MessageNumber   int64 // Message number
	NumberSVsInView int64 // Total number of SVs in view

	Info []GLGSVInfo // visible satellite info (0-4 of these)
}

// GLGSVInfo represents information about a visible satellite
type GLGSVInfo struct {
	SVPRNNumber int64 // SV PRN number, pseudo-random noise or gold code
	Elevation   int64 // Elevation in degrees, 90 maximum
	Azimuth     int64 // Azimuth, degrees from true north, 000 to 359
	SNR         int64 // SNR, 00-99 dB (null when not tracking)
}

// NewGLGSV constructor
func NewGLGSV(sentence Sentence) GLGSV {
	return GLGSV{Sentence: sentence}
}

// GetSentence getter
func (s GLGSV) GetSentence() Sentence {
	return s.Sentence
}

func (s *GLGSV) parse() error {
	p := newParser(s.Sentence, PrefixGLGSV)

	s.TotalMessages = p.Int64(0, "total number of messages")
	s.MessageNumber = p.Int64(1, "message number")
	s.NumberSVsInView = p.Int64(2, "number of SVs in view")

	s.Info = nil
	for i := 0; i < 4; i++ {
		if 5*i+4 > len(s.Fields) {
			break
		}
		s.Info = append(s.Info, GLGSVInfo{
			SVPRNNumber: p.Int64(3+i*4, "SV prn number"),
			Elevation:   p.Int64(4+i*4, "elevation"),
			Azimuth:     p.Int64(5+i*4, "azimuth"),
			SNR:         p.Int64(6+i*4, "SNR"),
		})
	}

	return p.Err()
}
