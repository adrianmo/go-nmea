package nmea

import (
	"fmt"
	"strconv"
)

const (
	// PrefixGPGSV prefix
	PrefixGPGSV = "GPGSV"
)

// GPGSV represents the GPS Satellites in view
// http://aprs.gids.nl/nmea/#gpgsv
type GPGSV struct {
	Sentence
	TotalMessages   int64 // Total number of messages of this type in this cycle
	MessageNumber   int64 // Message number
	NumberSVsInView int64 // Total number of SVs in view

	Info []GPGSVInfo // visible satellite info (0-4 of these)
}

// GPGSVInfo represents information about a visible satellite
type GPGSVInfo struct {
	SVPRNNumber int64 // SV PRN number, pseudo-random noise or gold code
	Elevation   int64 // Elevation in degrees, 90 maximum
	Azimuth     int64 // Azimuth, degrees from true north, 000 to 359
	SNR         int64 // SNR, 00-99 dB (null when not tracking)
}

// NewGPGSV constructor
func NewGPGSV(sentence Sentence) GPGSV {
	return GPGSV{Sentence: sentence}
}

// GetSentence getter
func (s GPGSV) GetSentence() Sentence {
	return s.Sentence
}

func (s *GPGSV) parse() error {
	if s.Type != PrefixGPGSV {
		return fmt.Errorf("%s is not a %s", s.Type, PrefixGPGSV)
	}
	var err error
	if s.Fields[0] != "" {
		s.TotalMessages, err = strconv.ParseInt(s.Fields[0], 10, 64)
		if err != nil {
			return fmt.Errorf("GPGSV decode total number of messages error: %s", s.Fields[0])
		}
	}

	if s.Fields[1] != "" {
		s.MessageNumber, err = strconv.ParseInt(s.Fields[1], 10, 64)
		if err != nil {
			return fmt.Errorf("GPGSV decode message number error: %s", s.Fields[1])
		}
	}

	if s.Fields[2] != "" {
		s.NumberSVsInView, err = strconv.ParseInt(s.Fields[2], 10, 64)
		if err != nil {
			return fmt.Errorf("GPGSV decode number of SVs in view error: %s", s.Fields[2])
		}
	}

	s.Info = nil
	for i := 0; i < 4; i++ {
		if 5*i+4 > len(s.Fields) {
			break
		}
		info := GPGSVInfo{}
		field := s.Fields[3+i*4]
		if s.Fields[3+i*4] != "" {
			info.SVPRNNumber, err = strconv.ParseInt(field, 10, 64)
			if err != nil {
				return fmt.Errorf("GPGSV decode SV prn number error: %s", field)
			}
		}

		field = s.Fields[4+i*4]
		if field != "" {
			info.Elevation, err = strconv.ParseInt(field, 10, 64)
			if err != nil {
				return fmt.Errorf("GPGSV decode elevation error: %s", field)
			}
		}

		field = s.Fields[5+i*4]
		if field != "" {
			info.Azimuth, err = strconv.ParseInt(field, 10, 64)
			if err != nil {
				return fmt.Errorf("GPGSV decode azimuth error: %s", field)
			}
		}

		field = s.Fields[6+i*4]
		if field != "" {
			info.SNR, err = strconv.ParseInt(field, 10, 64)
			if err != nil {
				return fmt.Errorf("GPGSV decode SNR error: %s", field)
			}
		}
		s.Info = append(s.Info, info)
	}

	return nil
}
