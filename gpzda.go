package nmea

import (
	"fmt"
	"strconv"
)

const (
	// PrefixGPZDA prefix
	PrefixGPZDA = "GPZDA"
)

// GPZDA represents date & time data.
// http://aprs.gids.nl/nmea/#zda
type GPZDA struct {
	Sentence
	Time  Time
	Day   int64
	Month int64
	Year  int64
	// Local time zone offset from GMT, hours
	OffsetHours int64
	// Local time zone offset from GMT, minutes
	OffsetMinutes int64
}

// NewGPZDA constructor
func NewGPZDA(sentence Sentence) GPZDA {
	s := new(GPZDA)
	s.Sentence = sentence
	return *s
}

// GetSentence getter
func (s GPZDA) GetSentence() Sentence {
	return s.Sentence
}

// Parse parses the GPZDA sentence into this struct.
func (s *GPZDA) parse() error {
	var err error

	if s.Type != PrefixGPZDA {
		return fmt.Errorf("%s is not a %s", s.Type, PrefixGPZDA)
	}

	s.Time = ParseTime(s.Fields[0])
	s.Day, err = strconv.ParseInt(s.Fields[1], 10, 64)
	if err != nil {
		return fmt.Errorf("GPZDA decode day error: %s", s.Fields[1])
	}

	s.Month, err = strconv.ParseInt(s.Fields[2], 10, 64)
	if err != nil {
		return fmt.Errorf("GPZDA decode month error: %s", s.Fields[2])
	}

	s.Year, err = strconv.ParseInt(s.Fields[3], 10, 64)
	if err != nil {
		return fmt.Errorf("GPZDA decode year error: %s", s.Fields[3])
	}

	s.OffsetHours, err = strconv.ParseInt(s.Fields[4], 10, 64)
	if err != nil {
		return fmt.Errorf("GPZDA decode offset (hours) error: %s", s.Fields[4])
	}

	s.OffsetMinutes, err = strconv.ParseInt(s.Fields[5], 10, 64)
	if err != nil {
		return fmt.Errorf("GPZDA decode offset (minutes) error: %s", s.Fields[5])
	}

	return nil
}
