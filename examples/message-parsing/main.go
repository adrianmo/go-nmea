package main

import (
	"fmt"
	"log"
	"os"

	"github.com/storskegg/go-nmea"
)

func main() {
	sentence := "$GPRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*70"
	s, err := nmea.Parse(sentence)
	if err != nil {
		log.Fatal(err)
	}
	reportMessage(s)

	sentence = "$GNGSA,A,3,26,22,,,,,,,,,,,2.99,1.43,2.63,1*06"
	s, err = nmea.Parse(sentence)
	if err != nil {
		log.Fatal(err)
	}
	reportMessage(s)
}

func reportMessage(s nmea.Sentence) {
	switch s.DataType() {
	case nmea.TypeRMC:
		m := s.(nmea.RMC)
		fmt.Printf("Raw sentence: %v\n", m)
		fmt.Printf("Time: %s\n", m.Time)
		fmt.Printf("Validity: %s\n", m.Validity)
		fmt.Printf("Latitude GPS: %s\n", nmea.FormatGPS(m.Latitude))
		fmt.Printf("Latitude DMS: %s\n", nmea.FormatDMS(m.Latitude))
		fmt.Printf("Longitude GPS: %s\n", nmea.FormatGPS(m.Longitude))
		fmt.Printf("Longitude DMS: %s\n", nmea.FormatDMS(m.Longitude))
		fmt.Printf("Speed: %f\n", m.Speed)
		fmt.Printf("Course: %f\n", m.Course)
		fmt.Printf("Date: %s\n", m.Date)
		fmt.Printf("Variation: %f\n", m.Variation)
	default:
		fmt.Fprintln(os.Stderr, "Unknown Message Type")
	}
}
