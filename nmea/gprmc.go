package nmea

import (
	"fmt"
	"strconv"
)

const (
	// PrefixGPRMC prefix of GPRMC sentence type
	PrefixGPRMC = "GPRMC"
	// ValidRMC character
	ValidRMC = "A"
	// InvalidRMC character
	InvalidRMC = "V"
)

// GPRMC is the Recommended Minimum Specific GNSS data.
// http://aprs.gids.nl/nmea/#rmc
type GPRMC struct {
	Sentence
	Time      string  // Time Stamp
	Validity  string  // validity - A-ok, V-invalid
	Latitude  LatLong // Latitude
	Longitude LatLong // Longitude
	Speed     float64 // Speed in knots
	Course    float64 // True course
	Date      string  // Date
	Variation float64 // Magnetic variation
}

func (m *GPRMC) parse() error {
	var err error

	if m.Type != PrefixGPRMC {
		return fmt.Errorf("%s is not a %s", m.Type, PrefixGPRMC)
	}
	m.Time = m.Fields[0]
	m.Validity = m.Fields[1]

	if m.Validity != ValidRMC && m.Validity != InvalidRMC {
		return fmt.Errorf("GPRMC decode, invalid validity '%s'", m.Validity)
	}

	m.Latitude, err = NewLatLong(fmt.Sprintf("%s %s", m.Fields[2], m.Fields[3]))
	if err != nil {
		return fmt.Errorf("GPRMC decode latitude error: %s", err)
	}
	m.Longitude, err = NewLatLong(fmt.Sprintf("%s %s", m.Fields[4], m.Fields[5]))
	if err != nil {
		return fmt.Errorf("GPRMC decode longitude error: %s", err)
	}
	m.Speed, err = strconv.ParseFloat(m.Fields[6], 64)
	if err != nil {
		return fmt.Errorf("GPRMC decode speed error: %s", m.Fields[6])
	}
	m.Course, err = strconv.ParseFloat(m.Fields[7], 64)
	if err != nil {
		return fmt.Errorf("GPRMC decode course error: %s", m.Fields[7])
	}
	m.Date = m.Fields[8]

	if m.Fields[9] != "" {
		m.Variation, err = strconv.ParseFloat(m.Fields[9], 64)
		if err != nil {
			return fmt.Errorf("GPRMC decode variation error: %s", m.Fields[9])
		}
		if m.Fields[10] == "W" {
			m.Variation = 0 - m.Variation
		} else if m.Fields[10] != "E" {
			return fmt.Errorf("GPRMC decode variation error for: %s", m.Fields[10])
		}
	}

	return nil
}
