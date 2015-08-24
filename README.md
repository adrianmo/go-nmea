# go-nmea [![Build Status](https://travis-ci.org/adrianmo/go-nmea.svg?branch=master)](https://travis-ci.org/adrianmo/go-nmea) [![Coverage Status](https://coveralls.io/repos/adrianmo/go-nmea/badge.svg?branch=master&service=github)](https://coveralls.io/github/adrianmo/go-nmea?branch=master) [![GoDoc](https://godoc.org/github.com/adrianmo/go-nmea?status.svg)](https://godoc.org/github.com/adrianmo/go-nmea)

This is a NMEA library for the Go programming language (http://golang.org).

## Installing

### Using `go get`

    go get github.com/adrianmo/go-nmea

After this command *go-nmea* is ready to use. Its source will be in:

    $GOPATH/src/github.com/adrianmo/go-nmea

## Supported sentences

At this moment, this library supports the following sentence types:

- [GPRMC](http://aprs.gids.nl/nmea/#rmc) - Recommended minimum specific GPS/Transit data
- [GPGGA](http://aprs.gids.nl/nmea/#gga) - Global Positioning System Fix Data
- [GPGSA](http://aprs.gids.nl/nmea/#gsa) - GPS DOP and active satellites
- [GPGLL](http://aprs.gids.nl/nmea/#gll) - Geographic Position, Latitude / Longitude and time

I will implement new types whenever I find some time. Also feel free to implement it yourself and send a pull-request to include it to the library.

## Example

```go
package main

import (
	"fmt"
	"github.com/adrianmo/go-nmea"
)

func main() {
	m, err := nmea.Parse("$GPRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*70")
	if err == nil {
		fmt.Printf("%+v\n", m)
	}
}
```
