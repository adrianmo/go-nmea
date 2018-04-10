# go-nmea [![Build Status](https://travis-ci.org/adrianmo/go-nmea.svg?branch=master)](https://travis-ci.org/adrianmo/go-nmea) [![Go Report Card](https://goreportcard.com/badge/github.com/adrianmo/go-nmea)](https://goreportcard.com/report/github.com/adrianmo/go-nmea) [![Coverage Status](https://coveralls.io/repos/adrianmo/go-nmea/badge.svg?branch=master&service=github)](https://coveralls.io/github/adrianmo/go-nmea?branch=master) [![GoDoc](https://godoc.org/github.com/adrianmo/go-nmea?status.svg)](https://godoc.org/github.com/adrianmo/go-nmea)

This is a NMEA library for the Go programming language (http://golang.org).

## Installing

### Using `go get`

    go get github.com/adrianmo/go-nmea

After this command *go-nmea* is ready to use. Its source will be in:

    $GOPATH/src/github.com/adrianmo/go-nmea

## Supported sentences

At this moment, this library supports the following sentence types:

- [GPRMC](http://aprs.gids.nl/nmea/#rmc) - Recommended Minimum Specific GPS/Transit data
- [GNRMC](http://aprs.gids.nl/nmea/#rmc) - Recommended Minimum Specific GNSS data
- [GPGGA](http://aprs.gids.nl/nmea/#gga) - GPS Positioning System Fix Data
- [GNGGA](http://aprs.gids.nl/nmea/#gga) - GNSS Positioning System Fix Data
- [GPGSA](http://aprs.gids.nl/nmea/#gsa) - GPS DOP and active satellites
- [GPGSV](http://aprs.gids.nl/nmea/#gsv) - GPS Satellites in view
- [GLGSV](http://aprs.gids.nl/nmea/#gsv) - GLONASS Satellites in view
- [GPGLL](http://aprs.gids.nl/nmea/#gll) - Geographic Position, Latitude / Longitude and time
- [GPVTG](http://aprs.gids.nl/nmea/#vtg) - Track Made Good and Ground Speed
- [GPZDA](http://aprs.gids.nl/nmea/#zda) - Date & time data
- [PGRME](http://aprs.gids.nl/nmea/#rme) - Estimated Position Error (Garmin proprietary sentence)


## Example

```go
package main

import (
	"fmt"
	"github.com/adrianmo/go-nmea"
)

func main() {
	sentence := "$GPRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*70"
	m, err := nmea.Parse(sentence)
	if err == nil {
		s := m.(nmea.GPRMC)
		fmt.Printf("Raw sentence: %v\n", s)
		fmt.Printf("Time: %v\n", s.Time.String())
		fmt.Printf("Validity: %v\n", s.Validity)
		fmt.Printf("Latitude GPS: %v\n", s.Latitude.PrintGPS())
		fmt.Printf("Latitude DMS: %v\n", s.Latitude.PrintDMS())
		fmt.Printf("Longitude GPS: %v\n", s.Longitude.PrintGPS())
		fmt.Printf("Longitude DMS: %v\n", s.Longitude.PrintDMS())
		fmt.Printf("Speed: %v\n", s.Speed)
		fmt.Printf("Course: %v\n", s.Course)
		fmt.Printf("Date: %v\n", s.Date.String())
		fmt.Printf("Variation: %v\n", s.Variation)
	} else {
		panic(err)
	}
}
```

Output:

```
$ go run main/main.go

Raw sentence: $GPRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*70
Time: 22:05:16.0000
Validity: A
Latitude GPS: 5133.8200
Latitude DMS: 51° 33' 49.200000"
Longitude GPS: 042.2400
Longitude DMS: 0° 42' 14.400000"
Speed: 173.8
Course: 231.8
Date: 13/06/94
Variation: -4.2
```

## Contributions

Please, feel free to implement support for new sentences, fix bugs, refactor code, etc. and send a pull-request to update the library.
