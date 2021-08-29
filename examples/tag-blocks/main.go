package main

import (
	"fmt"
	"log"
	"time"

	"github.com/storskegg/go-nmea"
)

func main() {
	sentence := "\\s:Satelite_1,c:1553390539*62\\!AIVDM,1,1,,A,13M@ah0025QdPDTCOl`K6`nV00Sv,0*52"
	s, err := nmea.Parse(sentence)
	if err != nil {
		log.Fatal(err)
	}
	parsed := s.(nmea.VDMVDO)
	fmt.Printf("TAG Block timestamp: %v\n", time.Unix(parsed.TagBlock.Time, 0))
	fmt.Printf("TAG Block source:    %v\n", parsed.TagBlock.Source)
}
