package nmea

import (
	"fmt"
	"strconv"
)

const (
	PrefixGPRMC = "GPRMC"
	ValidRMC    = "A"
	InvalidRMC  = "V"
)

// GPRMC is the Recommended Minimum Specific GNSS data.
// http://aprs.gids.nl/nmea/#rmc
type GPRMC struct {
	Sentence
	// Time.
	Time string
	// Status
	Status string
	// Latitude.
	Latitude LatLong
	// Longitude.
	Longitude LatLong
	// Speed over ground.
	Speed float64
	// Course over ground.
	Course float64
	// Date.
	Date string
	// Magnetic variation.
	Variation float64
}

// Parse parses the GPRMC sentence into this struct.
// e.g: $GPRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*70
func (g *GPRMC) Parse(s string) error {
	err := g.Sentence.Parse(s)
	if err != nil {
		// Sentence decode error.
		return err
	}
	if g.SType != PrefixGPRMC {
		return fmt.Errorf("%s is not a %s", g.SType, PrefixGPRMC)
	}
	g.Time = g.Fields[0]
	g.Status = g.Fields[1]
	if g.Status != ValidRMC && g.Status != InvalidRMC {
		return fmt.Errorf("GPRMC decode, invalid status '%s'", g.Status)
	}
	g.Latitude, err = NewLatLong(fmt.Sprintf("%s %s", g.Fields[2], g.Fields[3]))
	if err != nil {
		return fmt.Errorf("GPRMC decode latitude error: %s", err)
	}
	g.Longitude, err = NewLatLong(fmt.Sprintf("%s %s", g.Fields[4], g.Fields[5]))
	if err != nil {
		return fmt.Errorf("GPRMC decode longitude error: %s", err)
	}
	g.Speed, err = strconv.ParseFloat(g.Fields[6], 64)
	if err != nil {
		return fmt.Errorf("GPRMC decode speed error for: %s", s)
	}
	g.Course, err = strconv.ParseFloat(g.Fields[7], 64)
	if err != nil {
		return fmt.Errorf("GPRMC decode course error for: %s", s)
	}
	g.Date = g.Fields[8]
	g.Variation, err = strconv.ParseFloat(g.Fields[9], 64)
	if err != nil {
		return fmt.Errorf("GPRMC decode variation error for: %s", s)
	}
	if g.Fields[10] == "W" {
		g.Variation = 0 - g.Variation
	} else if g.Fields[10] != "E" {
		return fmt.Errorf("GPRMC decode variation error for: %s", s)
	}
	return nil
}
