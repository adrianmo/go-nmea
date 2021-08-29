# go-nmea

[![Build Status](https://travis-ci.com/adrianmo/go-nmea.svg?branch=master)](https://travis-ci.com/adrianmo/go-nmea) [![Go Report Card](https://goreportcard.com/badge/github.com/storskegg/go-nmea)](https://goreportcard.com/report/github.com/storskegg/go-nmea) [![Coverage Status](https://coveralls.io/repos/adrianmo/go-nmea/badge.svg?branch=master&service=github)](https://coveralls.io/github/adrianmo/go-nmea?branch=master) [![GoDoc](https://godoc.org/github.com/storskegg/go-nmea?status.svg)](https://godoc.org/github.com/storskegg/go-nmea)

This is a NMEA library for the Go programming language (Golang).

## Features

- Parse individual NMEA 0183 sentences
- Support for sentences with NMEA 4.10 "TAG Blocks"
- Register custom parser for unsupported sentence types
- User-friendly MIT license

## Installing

To install go-nmea use `go get`:

```
go get github.com/storskegg/go-nmea
```

This will then make the `github.com/storskegg/go-nmea` package available to you.

### Staying up to date

To update go-nmea to the latest version, use `go get -u github.com/storskegg/go-nmea`.

## Supported sentences

At this moment, this library supports the following sentence types:

| Sentence type                                                                       | Description                                                         |
| ----------------------------------------------------------------------------------- | ------------------------------------------------------------------- |
| [RMC](http://aprs.gids.nl/nmea/#rmc)                                                | Recommended Minimum Specific GPS/Transit data                       |
| [PMTK](https://www.rhydolabz.com/documents/25/PMTK_A11.pdf)                         | Messages for setting and reading commands for MediaTek gps modules. |
| [GGA](http://aprs.gids.nl/nmea/#gga)                                                | GPS Positioning System Fix Data                                     |
| [GSA](http://aprs.gids.nl/nmea/#gsa)                                                | GPS DOP and active satellites                                       |
| [GSV](http://aprs.gids.nl/nmea/#gsv)                                                | GPS Satellites in view                                              |
| [GLL](http://aprs.gids.nl/nmea/#gll)                                                | Geographic Position, Latitude / Longitude and time                  |
| [VTG](http://aprs.gids.nl/nmea/#vtg)                                                | Track Made Good and Ground Speed                                    |
| [ZDA](http://aprs.gids.nl/nmea/#zda)                                                | Date & time data                                                    |
| [HDT](http://aprs.gids.nl/nmea/#hdt)                                                | Actual vessel heading in degrees True                               |
| [GNS](https://www.trimble.com/oem_receiverhelp/v4.44/en/NMEA-0183messages_GNS.html) | Combined GPS fix for GPS, Glonass, Galileo, and BeiDou              |
| [PGRME](http://aprs.gids.nl/nmea/#rme)                                              | Estimated Position Error (Garmin proprietary sentence)              |
| [THS](http://www.nuovamarea.net/pytheas_9.html)                                     | Actual vessel heading in degrees True and status                    |
| [VDM/VDO](http://catb.org/gpsd/AIVDM.html)                                          | Encapsulated binary payload                                         |
| [WPL](http://aprs.gids.nl/nmea/#wpl)                                                | Waypoint location                                                   |
| [RTE](http://aprs.gids.nl/nmea/#rte)                                                | Route                                                               |
| [VHW](https://www.tronico.fi/OH6NT/docs/NMEA0183.pdf)                               | Water Speed and Heading                                             |
| [DPT](https://gpsd.gitlab.io/gpsd/NMEA.html#_dpt_depth_of_water)                    | Depth of Water                                                      |
| [DBS](https://gpsd.gitlab.io/gpsd/NMEA.html#_dbs_depth_below_surface)               | Depth Below Surface                                                 |
| [DBT](https://gpsd.gitlab.io/gpsd/NMEA.html#_dbt_depth_below_transducer)            | Depth below transducer                                              |

If you need to parse a message that contains an unsupported sentence type you can implement and register your own message parser and get yourself unblocked immediately. Check the example below to know how to [implement and register a custom message parser](#custom-message-parsing). However, if you think your custom message parser could be beneficial to other users we encourage you to contribute back to the library by submitting a PR and get it included in the list of supported sentences.

## Examples

### Built-in message parsing

```go
package main

import (
	"fmt"
	"log"
	"github.com/storskegg/go-nmea"
)

func main() {
	sentence := "$GPRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*70"
	s, err := nmea.Parse(sentence)
	if err != nil {
		log.Fatal(err)
	}
	if s.DataType() == nmea.TypeRMC {
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
Speed: 173.800000
Course: 231.800000
Date: 13/06/94
Variation: -4.200000
```

### TAG Blocks

NMEA 4.10 TAG Block values can be accessed via the message's `TagBlock` struct:

```go
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
```

Output (locale/time zone dependent):

```
$  go run main/main.go
 
TAG Block timestamp: 2019-03-24 14:22:19 +1300 NZDT
TAG Block source:    Satelite_1
```

### Custom message parsing

If you need to parse a message not supported by the library you can implement your own message parsing.
The following example implements a parser for the hypothetical XYZ NMEA sentence type.

```go
package main

import (
	"fmt"

	"github.com/storskegg/go-nmea"
)

// A type to hold the parsed record
type XYZType struct {
	nmea.BaseSentence
	Time    nmea.Time
	Counter int64
	Label   string
	Value   float64
}

func main() {
	// Do this once it will error if you register the same type multiple times
	err := nmea.RegisterParser("XYZ", func(s nmea.BaseSentence) (nmea.Sentence, error) {
		// This example uses the package builtin parsing helpers
		// you can implement your own parsing logic also
		p := nmea.NewParser(s)
		return XYZType{
			BaseSentence: s,
			Time:         p.Time(0, "time"),
			Label:        p.String(1, "label"),
			Counter:      p.Int64(2, "counter"),
			Value:        p.Float64(3, "value"),
		}, p.Err()
	})

	if err != nil {
		panic(err)
	}

	sentence := "$00XYZ,220516,A,23,5133.82,W*42"
	s, err := nmea.Parse(sentence)
	if err != nil {
		panic(err)
	}

	m, ok := s.(XYZType)
	if !ok {
		panic("Could not parse type XYZ")
	}

	fmt.Printf("Raw sentence: %v\n", m)
	fmt.Printf("Time: %s\n", m.Time)
	fmt.Printf("Label: %s\n", m.Label)
	fmt.Printf("Counter: %d\n", m.Counter)
	fmt.Printf("Value: %f\n", m.Value)
}
```

Output:

```
$ go run main/main.go

Raw sentence: $AAXYZ,220516,A,23,5133.82,W*42
Time: 22:05:16.0000
Label: A
Counter: 23
Value: 5133.820000
```

## Contributing

Please feel free to submit issues or fork the repository and send pull requests to update the library and fix bugs, implement support for new sentence types, refactor code, etc.

## License

Check [LICENSE](LICENSE).
