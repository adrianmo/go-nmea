# go-nmea

[![CI](https://github.com/adrianmo/go-nmea/actions/workflows/ci.yml/badge.svg)](https://github.com/adrianmo/go-nmea/actions/workflows/ci.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/adrianmo/go-nmea)](https://goreportcard.com/report/github.com/adrianmo/go-nmea) [![Coverage Status](https://coveralls.io/repos/adrianmo/go-nmea/badge.svg?branch=master&service=github)](https://coveralls.io/github/adrianmo/go-nmea?branch=master) [![GoDoc](https://godoc.org/github.com/adrianmo/go-nmea?status.svg)](https://godoc.org/github.com/adrianmo/go-nmea)

This is a NMEA library for the Go programming language (Golang).

## Features

- Parse individual NMEA 0183 sentences
- Support for sentences with NMEA 4.10 "TAG Blocks"
- Register custom parser for unsupported sentence types
- User-friendly MIT license

## Installing

To install go-nmea use `go get`:

```
go get github.com/adrianmo/go-nmea
```

This will then make the `github.com/adrianmo/go-nmea` package available to you.

### Staying up to date

To update go-nmea to the latest version, use `go get -u github.com/adrianmo/go-nmea`.

## Supported sentences

At this moment, this library supports the following sentence types:

| Sentence type                                                                                 | Description                                               |
|-----------------------------------------------------------------------------------------------|-----------------------------------------------------------|
| [AAM](https://gpsd.gitlab.io/gpsd/NMEA.html#_aam_waypoint_arrival_alarm)                      | Waypoint Arrival Alarm                                    |
| [ALA](./ala.go)                                                                               | System Faults and Alarms                                  |
| [APB](https://gpsd.gitlab.io/gpsd/NMEA.html#_apb_autopilot_sentence_b)                        | Autopilot Sentence "B"                                    |
| [BEC](http://www.nmea.de/nmea0183datensaetze.html#bec)                                        | Bearing and distance to waypoint (dead reckoning)         |
| [BOD](https://gpsd.gitlab.io/gpsd/NMEA.html#_bod_bearing_waypoint_to_waypoint)                | Bearing waypoint to waypoint (origin to destination)      |
| [BWC](https://gpsd.gitlab.io/gpsd/NMEA.html#_bwc_bearing_distance_to_waypoint_great_circle)   | Bearing and distance to waypoint (great circle)           |
| [BWR](https://gpsd.gitlab.io/gpsd/NMEA.html#_bwr_bearing_and_distance_to_waypoint_rhumb_line) | Bearing and distance to waypoint (Rhumb Line)             |
| [BWW](https://gpsd.gitlab.io/gpsd/NMEA.html#_bww_bearing_waypoint_to_waypoint)                | Bearing from destination waypoint to origin waypoint      |
| [DBK](https://gpsd.gitlab.io/gpsd/NMEA.html#_dbk_depth_below_keel)                            | Depth Below Keel (obsolete, use DPT instead)              |
| [DBS](https://gpsd.gitlab.io/gpsd/NMEA.html#_dbs_depth_below_surface)                         | Depth Below Surface (obsolete, use DPT instead)           |
| [DBT](https://gpsd.gitlab.io/gpsd/NMEA.html#_dbt_depth_below_transducer)                      | Depth below transducer                                    |
| [DOR](./dor.go)                                                                               | Door Status Detection                                     |
| [DPT](https://gpsd.gitlab.io/gpsd/NMEA.html#_dpt_depth_of_water)                              | Depth of Water                                            |
| [DSC](./dsc.go)                                                                               | Digital Selective Calling Information                     |
| [DSE](./dse.go)                                                                               | Expanded digital selective calling                        |
| [DTM](https://gpsd.gitlab.io/gpsd/NMEA.html#_dtm_datum_reference)                             | Datum Reference                                           |
| [EVE](./eve.go)                                                                               | General Event Message                                     |
| [FIR](./fir.go)                                                                               | Fire Detection event with time and location               |
| [GGA](http://aprs.gids.nl/nmea/#gga)                                                          | GPS Positioning System Fix Data                           |
| [GLL](http://aprs.gids.nl/nmea/#gll)                                                          | Geographic Position, Latitude / Longitude and time        |
| [GNS](https://gpsd.gitlab.io/gpsd/NMEA.html#_gns_fix_data)                                    | Combined GPS fix for GPS, Glonass, Galileo, and BeiDou    |
| [GSA](http://aprs.gids.nl/nmea/#gsa)                                                          | GPS DOP and active satellites                             |
| [GSV](http://aprs.gids.nl/nmea/#gsv)                                                          | GPS Satellites in view                                    |
| [HDG](https://gpsd.gitlab.io/gpsd/NMEA.html#_hdg_heading_deviation_variation)                 | Heading, Deviation & Variation                            |
| [HDM](https://gpsd.gitlab.io/gpsd/NMEA.html#_hdm_heading_magnetic)                            | Heading - Magnetic                                        |
| [HDT](http://aprs.gids.nl/nmea/#hdt)                                                          | Actual vessel heading in degrees True                     |
| [HSC](https://gpsd.gitlab.io/gpsd/NMEA.html#_hsc_heading_steering_command)                    | Heading steering command                                  |
| [MDA](https://gpsd.gitlab.io/gpsd/NMEA.html#_mda_meteorological_composite)                    | Meteorological Composite                                  |
| [MTA](./mta.go)                                                                               | Air Temperature (obsolete, use XDR instead)               |
| [MTW](https://gpsd.gitlab.io/gpsd/NMEA.html#_mtw_mean_temperature_of_water)                   | Mean Temperature of Water                                 |
| [MWD](https://www.tronico.fi/OH6NT/docs/NMEA0183.pdf)                                         | Wind Direction and Speed                                  |
| [MWV](https://gpsd.gitlab.io/gpsd/NMEA.html#_mwv_wind_speed_and_angle)                        | Wind Speed and Angle                                      |
| [OSD](https://gpsd.gitlab.io/gpsd/NMEA.html#_osd_own_ship_data)                               | Own Ship Data                                             |
| [RMB](https://gpsd.gitlab.io/gpsd/NMEA.html#_rmb_recommended_minimum_navigation_information)  | Recommended Minimum Navigation Information                |
| [RMC](http://aprs.gids.nl/nmea/#rmc)                                                          | Recommended Minimum Specific GPS/Transit data             |
| [ROT](https://gpsd.gitlab.io/gpsd/NMEA.html#_rot_rate_of_turn)                                | Rate of turn                                              |
| [RPM](https://gpsd.gitlab.io/gpsd/NMEA.html#_rpm_revolutions)                                 | Engine or Shaft revolutions and pitch                     |
| [RSA](https://gpsd.gitlab.io/gpsd/NMEA.html#_rsa_rudder_sensor_angle)                         | Rudder Sensor Angle                                       |
| [RSD](https://gpsd.gitlab.io/gpsd/NMEA.html#_rsd_radar_system_data)                           | RADAR System Data                                         |
| [RTE](http://aprs.gids.nl/nmea/#rte)                                                          | Route                                                     |
| [THS](http://www.nuovamarea.net/pytheas_9.html)                                               | Actual vessel heading in degrees True and status          |
| [TLL](https://gpsd.gitlab.io/gpsd/NMEA.html#_tll_target_latitude_and_longitude)               | Target latitude and longitude                             |
| [TTM](https://gpsd.gitlab.io/gpsd/NMEA.html#_ttm_tracked_target_message)                      | Tracked Target Message                                    |
| [TXT](https://www.nmea.org/Assets/20160520%20txt%20amendment.pdf)                             | Sentence is for the transmission of text messages         |
| [VBW](https://gpsd.gitlab.io/gpsd/NMEA.html#_vbw_dual_groundwater_speed)                      | Dual Ground/Water Speed                                   |
| [VDM/VDO](https://gpsd.gitlab.io/gpsd/AIVDM.html)                                             | Encapsulated binary payload (commonly used with AIS data) |
| [VDR](https://gpsd.gitlab.io/gpsd/NMEA.html#_vdr_set_and_drift)                               | Set and Drift                                             |
| [VHW](https://www.tronico.fi/OH6NT/docs/NMEA0183.pdf)                                         | Water Speed and Heading                                   |
| [VLW](https://gpsd.gitlab.io/gpsd/NMEA.html#_vlw_distance_traveled_through_water)             | Distance Traveled through Water                           |
| [VPW](https://gpsd.gitlab.io/gpsd/NMEA.html#_vpw_speed_measured_parallel_to_wind)             | Speed Measured Parallel to Wind                           |
| [VTG](http://aprs.gids.nl/nmea/#vtg)                                                          | Track Made Good and Ground Speed                          |
| [VWR](https://gpsd.gitlab.io/gpsd/NMEA.html#_vwr_relative_wind_speed_and_angle)               | Relative Wind Speed and Angle                             |
| [VWT](./vwt.go)                                                                               | True Wind Speed and Angle                                 |
| [WPL](http://aprs.gids.nl/nmea/#wpl)                                                          | Waypoint location                                         |
| [XDR](https://gpsd.gitlab.io/gpsd/NMEA.html#_xdr_transducer_measurement)                      | Transducer Measurement                                    |
| [ZDA](http://aprs.gids.nl/nmea/#zda)                                                          | Date & time data                                          |

| Proprietary sentence type                                   | Description                                                                                     |
|-------------------------------------------------------------|-------------------------------------------------------------------------------------------------|
| [PGRME](http://aprs.gids.nl/nmea/#rme)                      | Estimated Position Error (Garmin proprietary sentence)                                          |
| [PHTRO](#)                                                  | Vessel pitch and roll (Xsens IMU/VRU/AHRS)                                                      |
| [PMTK](https://www.rhydolabz.com/documents/25/PMTK_A11.pdf) | Messages for setting and reading commands for MediaTek gps modules.                             |
| [PRDID](#)                                                  | Vessel pitch, roll and heading (Xsens IMU/VRU/AHRS)                                             |
| [PSONCMS](#)                                                | Quaternion, acceleration, rate of turn, magnetic field, sensor temperature (Xsens IMU/VRU/AHRS) |

If you need to parse a message that contains an unsupported sentence type you can implement and register your own
message parser and get yourself unblocked immediately. Check the example below to know how
to [implement and register a custom message parser](#custom-message-parsing). However, if you think your custom message
parser could be beneficial to other users we encourage you to contribute back to the library by submitting a PR and get
it included in the list of supported sentences.

## Examples

### Built-in message parsing

```go
package main

import (
	"fmt"
	"log"
	"github.com/adrianmo/go-nmea"
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
	"github.com/adrianmo/go-nmea"
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

If you need to parse a message not supported by the library you can implement your own message parsing. The following
example implements a parser for the hypothetical XYZ NMEA sentence type.

```go
package main

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
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

	switch m := s.(type) {
	case XYZType:
		fmt.Printf("Raw sentence: %v\n", m)
		fmt.Printf("Time: %s\n", m.Time)
		fmt.Printf("Label: %s\n", m.Label)
		fmt.Printf("Counter: %d\n", m.Counter)
		fmt.Printf("Value: %f\n", m.Value)
	default:
		panic("Could not parse XYZ sentence")
	}
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

Please feel free to submit issues or fork the repository and send pull requests to update the library and fix bugs,
implement support for new sentence types, refactor code, etc.

## License

Check [LICENSE](LICENSE).
