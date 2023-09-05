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

Sentence with link is supported by this library. NMEA0183 sentences list is based
on [IEC 61162-1:2016 (Edition 5.0 2016-08)](https://webstore.iec.ch/publication/25754) table of contents.

| Sentence           | Description                                                         | References                                                                                     |
|--------------------|---------------------------------------------------------------------|------------------------------------------------------------------------------------------------|
| [AAM](./aam.go)    | Waypoint arrival alarm                                              | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_aam_waypoint_arrival_alarm)                      |
| ABK                | AIS addressed and binary broadcast acknowledgement                  |                                                                                                |
| [ABM](./abm.go)    | AIS addressed binary and safety related message                     |                                                                                                |
| ACA                | AIS channel assignment message                                      |                                                                                                |
| [ACK](./ack.go)    | Acknowledge alarm                                                   |                                                                                                |
| [ACN](./acn.go)    | Alert command                                                       |                                                                                                |
| ACS                | AIS channel management information source                           |                                                                                                |
| AIR                | AIS interrogation request                                           |                                                                                                |
| AKD                | Acknowledge detail alarm condition                                  |                                                                                                |
| [ALA](./ala.go)    | Report detailed alarm condition                                     |                                                                                                |
| [ALC](./alc.go)    | Cyclic alert list                                                   |                                                                                                |
| [ALF](./alf.go)    | Alert sentence                                                      |                                                                                                |
| [ALR](./alr.go)    | Set alarm state                                                     |                                                                                                |
| [APB](./apb.go)    | Heading/track controller (autopilot) sentence B                     | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_apb_autopilot_sentence_b)                        |
| [ARC](./arc.go)    | Alert command refused                                               |                                                                                                |
| [AZT](./azt.go)    | Azimuth Thruster message                                            |                                                                                                |
| [BBM](./bbm.go)    | AIS broadcast binary message                                        |                                                                                                |
| [BEC](./bec.go)    | Bearing and distance to waypoint, Dead reckoning                    | [1](http://www.nmea.de/nmea0183datensaetze.html#bec)                                           |
| [BOD](./bod.go)    | Bearing origin to destination                                       | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_bod_bearing_waypoint_to_waypoint)                |
| [BWC](./bwc.go)    | Bearing and distance to waypoint, Great circle                      | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_bwc_bearing_distance_to_waypoint_great_circle)   |
| [BWR](./bwr.go)    | Bearing and distance to waypoint, Rhumb line                        | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_bwr_bearing_and_distance_to_waypoint_rhumb_line) |
| [BWW](./bww.go)    | Bearing waypoint to waypoint                                        | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_bww_bearing_waypoint_to_waypoint)                |
| CUR                | Water current layer, Multi-layer water current data                 |                                                                                                |
| [DBK](./dbk.go)    | Depth Below Keel (obsolete, use DPT instead)                        | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_dbk_depth_below_keel)                            |
| [DBS](./dbs.go)    | Depth below transducer                                              | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_dbs_depth_below_surface)                         |
| [DBT](./dbt.go)    | Depth below transducer                                              | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_dbt_depth_below_transducer)                      |
| DDC                | Display dimming control                                             |                                                                                                |
| [DOR](./dor.go)    | Door status detection                                               |                                                                                                |
| [DPT](./dpt.go)    | Depth                                                               | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_dpt_depth_of_water)                              |
| [DSC](./dsc.go)    | Digital selective calling information                               |                                                                                                |
| [DSE](./dse.go)    | Expanded digital selective calling                                  |                                                                                                |
| [DTM](./dtm.go)    | Datum reference                                                     | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_dtm_datum_reference)                             |
| EPV                | Command or report equipment property value                          |                                                                                                |
| ETL                | Engine telegraph operation status                                   |                                                                                                |
| [EVE](./eve.go)    | General event message                                               |                                                                                                |
| [FIR](./fir.go)    | Fire detection                                                      |                                                                                                |
| FSI                | Frequency set information                                           |                                                                                                |
| GBS                | GNSS satellite fault detection                                      |                                                                                                |
| GEN                | Generic binary information                                          |                                                                                                |
| GFA                | GNSS fix accuracy and integrity                                     |                                                                                                |
| [GGA](./gga.go)    | Global positioning system (GPS) fix data                            | [1](http://aprs.gids.nl/nmea/#gga)                                                             |
| [GLL](./gll.go)    | Geographic position, Latitude/longitude                             | [1](http://aprs.gids.nl/nmea/#gll)                                                             |
| [GNS](./gns.go)    | GNSS fix data                                                       | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_gns_fix_data)                                    |
| GRS                | GNSS range residuals                                                |                                                                                                |
| [GSA](./gsa.go)    | GNSS DOP and active satellites                                      | [1](http://aprs.gids.nl/nmea/#gsa)                                                             |
| GST                | GNSS pseudorange noise statistics                                   |                                                                                                |
| [GSV](./gsv.go)    | GNSS satellites in view                                             | [1](http://aprs.gids.nl/nmea/#gsv)                                                             |
| [HBT](./hbt.go)    | Heartbeat supervision sentence                                      |                                                                                                |
| HCR                | Heading correction report                                           |                                                                                                |
| [HDG](./hdg.go)    | Heading, deviation and variation                                    | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_hdg_heading_deviation_variation)                 |
| HDM                | Heading - Magnetic                                                  | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_hdm_heading_magnetic)                            |
| [HDT](./hdt.go)    | Heading true                                                        | [gpsd](http://aprs.gids.nl/nmea/#hdt)                                                          |
| HMR                | Heading monitor receive                                             |                                                                                                |
| HMS                | Heading monitor set                                                 |                                                                                                |
| HRM                | heel angle, roll period and roll amplitude measurement device       |                                                                                                |
| [HSC](./hsc.go)    | Heading steering command                                            | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_hsc_heading_steering_command)                    |
| HSS                | Hull stress surveillance systems                                    |                                                                                                |
| HTC                | Heading/track control command                                       |                                                                                                |
| HTD                | Heading /track control data                                         |                                                                                                |
| LR1                | AIS long-range reply sentence 1                                     |                                                                                                |
| LR2                | AIS long-range reply sentence 2                                     |                                                                                                |
| LR3                | AIS long-range reply sentence 3                                     |                                                                                                |
| LRF                | AIS long-range function                                             |                                                                                                |
| LRI                | AIS long-range interrogation                                        |                                                                                                |
| [MCP](./mcp.go)    | Micropilot Joystick controller message                              |                                                                                                |
| [MDA](./mda.go)    | Meteorological Composite                                            | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_mda_meteorological_composite)                    |
| [MTA](./mta.go)    | Air Temperature (obsolete, use XDR instead)                         |                                                                                                |
| MOB                | Man over board notification                                         |                                                                                                |
| MSK                | MSK receiver interface                                              |                                                                                                |
| MSS                | MSK receiver signal status                                          |                                                                                                |
| [MTW](./mtw.go)    | Water temperature                                                   | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_mtw_mean_temperature_of_water)                   |
| [MWD](./mwd.go)    | Wind direction and speed                                            | [1](https://www.tronico.fi/OH6NT/docs/NMEA0183.pdf)                                            |
| [MWV](./mwv.go)    | Wind speed and angle                                                | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_mwv_wind_speed_and_angle)                        |
| NAK                | Negative acknowledgement                                            |                                                                                                |
| NRM                | NAVTEX receiver mask                                                |                                                                                                |
| NRX                | NAVTEX received message                                             |                                                                                                |
| NSR                | Navigation status report                                            |                                                                                                |
| [OSD](./osd.go)    | Own ship data                                                       | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_osd_own_ship_data)                               |
| POS                | Device position and ship dimensions report or configuration command |                                                                                                |
| PRC                | Propulsion remote control status                                    |                                                                                                |
| RLM                | Return link message                                                 |                                                                                                |
| RMA                | Recommended minimum specific LORAN-C data                           |                                                                                                |
| [RMB](./rmb.go)    | Recommended minimum navigation information                          | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_rmb_recommended_minimum_navigation_information)  |
| [RMC](./rmc.go)    | Recommended minimum specific GNSS data                              | [1](http://aprs.gids.nl/nmea/#rmc)                                                             |
| ROR                | Rudder order status                                                 |                                                                                                |
| [ROT](./rot.go)    | Rate of turn                                                        | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_rot_rate_of_turn)                                |
| RRT                | Report route transfer                                               |                                                                                                |
| [RPM](./rpm.go)    | Revolutions                                                         | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_rpm_revolutions)                                 |
| [RSA](./rsa.go)    | Rudder sensor angle                                                 | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_rsa_rudder_sensor_angle)                         |
| [RSD](./rsd.go)    | Radar system data                                                   | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_rsd_radar_system_data)                           |
| [RTE](./rte.go)    | Routes                                                              | [1](http://aprs.gids.nl/nmea/#rte)                                                             |
| SFI                | Scanning frequency information                                      |                                                                                                |
| SMI                | SafetyNET Message, All Ships/NavArea                                |                                                                                                |
| SM2                | SafetyNET Message, Coastal Warning Area                             |                                                                                                |
| SM3                | SafetyNET Message, Circular Area address                            |                                                                                                |
| SM4                | SafetyNET Message, Rectangular Area Address                         |                                                                                                |
| SMB                | IMO SafetyNET Message Body                                          |                                                                                                |
| SPW                | Security password sentence                                          |                                                                                                |
| SSD                | AIS ship static data                                                |                                                                                                |
| STN                | Multiple data ID                                                    |                                                                                                |
| [THS](./ths.go)    | True heading and status                                             | [1](http://www.nuovamarea.net/pytheas_9.html)                                                  |
| [TLB](./tlb.go)    | Target label                                                        |                                                                                                |
| [TLL](./tll.go)    | Target latitude and longitude                                       | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_tll_target_latitude_and_longitude)               |
| TRC                | Thruster control data                                               |                                                                                                |
| TRL                | AIS transmitter-non-functioning log                                 |                                                                                                |
| TRD                | Thruster response data                                              |                                                                                                |
| [TTD](./ttd.go)    | Tracked target data                                                 |                                                                                                |
| [TTM](./ttm.go)    | Tracked target message                                              | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_ttm_tracked_target_message)                      |
| TUT                | Transmission of multi-language text                                 |                                                                                                |
| [TXT](./txt.go)    | Text transmission                                                   | [NMEA](https://www.nmea.org/Assets/20160520%20txt%20amendment.pdf)                             |
| UID                | User identification code transmission                               |                                                                                                |
| [VBW](./vbw.go)    | Dual ground/water speed                                             | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_vbw_dual_groundwater_speed)                      |
| [VDM](./vdmvdo.go) | AIS VHF data-link message                                           | [gpsd](https://gpsd.gitlab.io/gpsd/AIVDM.html)                                                 |
| [VDO](./vdmvdo.go) | AIS VHF data-link own-vessel report                                 | [gpsd](https://gpsd.gitlab.io/gpsd/AIVDM.html)                                                 |
| [VDR](./vdr.go)    | Set and drift                                                       | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_vdr_set_and_drift)                               |
| VER                | Version                                                             |                                                                                                |
| [VHW](./vhw.go)    | Water speed and heading                                             | [1](https://www.tronico.fi/OH6NT/docs/NMEA0183.pdf)                                            |
| [VLW](./vlw.go)    | Dual ground/water distance                                          | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_vlw_distance_traveled_through_water)             |
| [VPW](./vpw.go)    | Speed measured parallel to wind                                     | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_vpw_speed_measured_parallel_to_wind)             |
| [VSD](./vsd.go)    | AIS voyage static data                                              |                                                                                                |
| [VTG](./vtg.go)    | Course over ground and ground speed                                 | [1](http://aprs.gids.nl/nmea/#vtg)                                                             |
| VWR                | Relative Wind Speed and Angle                                       | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_vwr_relative_wind_speed_and_angle)               |
| VWT                | True Wind Speed and Angle                                           |                                                                                                |
| WAT                | Water level detection                                               |                                                                                                |
| WCV                | Waypoint closure velocity                                           |                                                                                                |
| WNC                | Distance waypoint to waypoint                                       |                                                                                                |
| [WPL](./wpl.go)    | Waypoint location                                                   | [1](http://aprs.gids.nl/nmea/#wpl)                                                             |
| [XDR](./xdr.go)    | Transducer measurements                                             | [gpsd](https://gpsd.gitlab.io/gpsd/NMEA.html#_xdr_transducer_measurement)                      |
| [XTE](./xte.go)    | Cross-track error, measured                                         |                                                                                                |
| XTR                | Cross-track error, dead reckoning                                   |                                                                                                |
| [ZDA](./zda.go)    | Time and date                                                       | [1](http://aprs.gids.nl/nmea/#zda)                                                             |
| ZDL                | Time and distance to variable point                                 |                                                                                                |
| ZFO                | UTC and time from origin waypoint                                   |                                                                                                |
| ZTG                | UTC and time to destination waypoint                                |                                                                                                |

| Proprietary sentence type | Description                                                                                     | References                                                                                         |
|---------------------------|-------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------|
| [PNG](./pgn.go)           | Transfer NMEA2000 frame as NMEA0183 sentence (ShipModul MiniPlex-3)                             | [1](https://opencpn.org/wiki/dokuwiki/lib/exe/fetch.php?media=opencpn:software:mxpgn_sentence.pdf) |
| [PCDIN](./pcdin.go)       | Transfer NMEA2000 frame as NMEA0183 sentence (SeaSmart.Net Protocol)                            | [1](http://www.seasmart.net/pdf/SeaSmart_HTTP_Protocol_RevG_043012.pdf)                            |
| [PGRME](./pgrme.go)       | Estimated Position Error (Garmin proprietary sentence)                                          | [1](http://aprs.gids.nl/nmea/#rme)                                                                 |
| [PHTRO](./phtro.go)       | Vessel pitch and roll (Xsens IMU/VRU/AHRS)                                                      |                                                                                                    |
| [PMTK001](./pmtk.go)      | Acknowledgement of previously sent command/packet                                               | [1](https://www.rhydolabz.com/documents/25/PMTK_A11.pdf)                                           |
| [PRDID](./prdid.go)       | Vessel pitch, roll and heading (Xsens IMU/VRU/AHRS)                                             |                                                                                                    |
| [PSKPDPT](./pskpdpt.go)   | Depth of Water for multiple transducer installation                                             |                                                                                                    |
| [PSONCMS](./psoncms.go)   | Quaternion, acceleration, rate of turn, magnetic field, sensor temperature (Xsens IMU/VRU/AHRS) |                                                                                                    |

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

### Customize sentence parser

Parser logic can be customized by creating `nmea.SentenceParser` instance and by providing callback implementations.

```go
p := nmea.SentenceParser{
    CustomParsers: nil,
    ParsePrefix:   nil,
    CheckCRC:      nil,
    OnTagBlock:    nil,
}
s, err := p.Parse("$GPRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*70")
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

### Message parsing with optional values

Some messages have optional fields. By default, omitted numeric values are set to 0. In situations where you need finer
control to distinguish between an undefined value and an actual 0, you can register types overriding existing sentences,
using `nmea.Int64` and `nmea.Float64` instead of `int64` and `float64`. The matching parsing methods
are `(*Parser).NullInt64` and `(*Parser).NullFloat64`. Both `nmea.Int64` and `nmea.Float64` contains a numeric
field `Value` which is defined only if the field `Valid` is `true`.

See below example for a modified VTG sentence parser:

```go
package main

import (
	"fmt"

	"github.com/adrianmo/go-nmea"
)

// VTG represents track & speed data.
// http://aprs.gids.nl/nmea/#vtg
type VTG struct {
	nmea.BaseSentence
	TrueTrack        nmea.Float64
	MagneticTrack    nmea.Float64
	GroundSpeedKnots nmea.Float64
	GroundSpeedKPH   nmea.Float64
}

func main() {
	nmea.MustRegisterParser("VTG", func(s nmea.BaseSentence) (nmea.Sentence, error) {
		p := nmea.NewParser(s)
		return VTG{
			BaseSentence:     s,
			TrueTrack:        p.NullFloat64(0, "true track"),
			MagneticTrack:    p.NullFloat64(2, "magnetic track"),
			GroundSpeedKnots: p.NullFloat64(4, "ground speed (knots)"),
			GroundSpeedKPH:   p.NullFloat64(6, "ground speed (km/h)"),
		}, p.Err()
	})

	sentence := "$GPVTG,140.88,T,,M,8.04,N,14.89,K,D*05"
	s, err := nmea.Parse(sentence)
	if err != nil {
		panic(err)
	}

	m, ok := s.(VTG)
	if !ok {
		panic("Could not parse VTG sentence")
	}
	fmt.Printf("Raw sentence: %v\n", m)
	fmt.Printf("TrueTrack: %v\n", m.TrueTrack)
	fmt.Printf("MagneticTrack: %v\n", m.MagneticTrack)
	fmt.Printf("GroundSpeedKnots: %v\n", m.GroundSpeedKnots)
	fmt.Printf("GroundSpeedKPH: %v\n", m.GroundSpeedKPH)
}
```

Output:

```
$ go run main/main.go

Raw sentence: $GPVTG,140.88,T,,M,8.04,N,14.89,K,D*05
TrueTrack: {140.88 true}
MagneticTrack: {0 false}
GroundSpeedKnots: {8.04 true}
GroundSpeedKPH: {14.89 true}
```

## Contributing

Please feel free to submit issues or fork the repository and send pull requests to update the library and fix bugs,
implement support for new sentence types, refactor code, etc.

## License

Check [LICENSE](LICENSE).
