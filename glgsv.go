package nmea

import (
	"fmt"
	"strconv"
)

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
	if s.Type != PrefixGLGSV {
		return fmt.Errorf("%s is not a %s", s.Type, PrefixGLGSV)
	}
	var err error
	if s.Fields[0] != "" {
		s.TotalMessages, err = strconv.ParseInt(s.Fields[0], 10, 64)
		if err != nil {
			return fmt.Errorf("GLGSV decode total number of messages error: %s", s.Fields[0])
		}
	}

	if s.Fields[1] != "" {
		s.MessageNumber, err = strconv.ParseInt(s.Fields[1], 10, 64)
		if err != nil {
			return fmt.Errorf("GLGSV decode message number error: %s", s.Fields[1])
		}
	}

	if s.Fields[2] != "" {
		s.NumberSVsInView, err = strconv.ParseInt(s.Fields[2], 10, 64)
		if err != nil {
			return fmt.Errorf("GLGSV decode number of SVs in view error: %s", s.Fields[2])
		}
	}

	s.Info = nil
	for i := 0; i < 4; i++ {
		if 5*i+4 > len(s.Fields) {
			break
		}
		info := GLGSVInfo{}
		field := s.Fields[3+i*4]
		if s.Fields[3+i*4] != "" {
			info.SVPRNNumber, err = strconv.ParseInt(field, 10, 64)
			if err != nil {
				return fmt.Errorf("GLGSV decode SV prn number error: %s", field)
			}
		}

		field = s.Fields[4+i*4]
		if field != "" {
			info.Elevation, err = strconv.ParseInt(field, 10, 64)
			if err != nil {
				return fmt.Errorf("GLGSV decode elevation error: %s", field)
			}
		}

		field = s.Fields[5+i*4]
		if field != "" {
			info.Azimuth, err = strconv.ParseInt(field, 10, 64)
			if err != nil {
				return fmt.Errorf("GLGSV decode azimuth error: %s", field)
			}
		}

		field = s.Fields[6+i*4]
		if field != "" {
			info.SNR, err = strconv.ParseInt(field, 10, 64)
			if err != nil {
				return fmt.Errorf("GLGSV decode SNR error: %s", field)
			}
		}
		s.Info = append(s.Info, info)
	}

	return nil
}
