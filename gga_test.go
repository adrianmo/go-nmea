package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var ggatests = []struct {
	name string
	raw  string
	err  string
	msg  GGA
}{
	{
		name: "good sentence",
		raw:  "$GNGGA,203415.000,6325.6138,N,01021.4290,E,1,8,2.42,72.5,M,41.5,M,,*7C",
		msg: GGA{
			Time: Time{
				Valid:       true,
				Hour:        20,
				Minute:      34,
				Second:      15,
				Millisecond: 0,
			},
			Latitude:      MustParseLatLong("6325.6138 N"),
			Longitude:     MustParseLatLong("01021.4290 E"),
			LatDirection:  "N",
			LonDirection:  "E",
			FixQuality:    "1",
			NumSatellites: 8,
			HDOP:          2.42,
			Altitude:      72.5,
			Separation:    41.5,
			DGPSAge:       "",
			DGPSId:        "",
		},
	},
	{
		name: "bad latitude",
		raw:  "$GNGGA,034225.077,A,S,15124.5567,E,1,03,9.7,-25.0,M,21.0,M,,0000*24",
		err:  "nmea: GNGGA invalid latitude: cannot parse [A S], unknown format",
	},
	{
		name: "bad longitude",
		raw:  "$GNGGA,034225.077,3356.4650,S,A,E,1,03,9.7,-25.0,M,21.0,M,,0000*12",
		err:  "nmea: GNGGA invalid longitude: cannot parse [A E], unknown format",
	},
	{
		name: "bad fix quality",
		raw:  "$GNGGA,034225.077,3356.4650,S,15124.5567,E,12,03,9.7,-25.0,M,21.0,M,,0000*7D",
		err:  "nmea: GNGGA invalid fix quality: 12",
	},
	{
		name: "good sentence",
		raw:  "$GPGGA,034225.077,3356.4650,S,15124.5567,E,1,03,9.7,-25.0,M,21.0,M,,0000*51",
		msg: GGA{
			Time:          Time{true, 3, 42, 25, 77},
			Latitude:      MustParseLatLong("3356.4650 S"),
			Longitude:     MustParseLatLong("15124.5567 E"),
			LatDirection:  "S",
			LonDirection:  "E",
			FixQuality:    GPS,
			NumSatellites: 03,
			HDOP:          9.7,
			Altitude:      -25.0,
			Separation:    21.0,
			DGPSAge:       "",
			DGPSId:        "0000",
		},
	},
	{
		name: "bad latitude",
		raw:  "$GPGGA,034225.077,A,S,15124.5567,E,1,03,9.7,-25.0,M,21.0,M,,0000*3A",
		err:  "nmea: GPGGA invalid latitude: cannot parse [A S], unknown format",
	},
	{
		name: "bad longitude",
		raw:  "$GPGGA,034225.077,3356.4650,S,A,E,1,03,9.7,-25.0,M,21.0,M,,0000*0C",
		err:  "nmea: GPGGA invalid longitude: cannot parse [A E], unknown format",
	},
	{
		name: "bad fix quality",
		raw:  "$GPGGA,034225.077,3356.4650,S,15124.5567,E,12,03,9.7,-25.0,M,21.0,M,,0000*63",
		err:  "nmea: GPGGA invalid fix quality: 12",
	},
}

func TestGGA(t *testing.T) {
	for _, tt := range ggatests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				gga := m.(GGA)
				gga.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, gga)
			}
		})
	}
}
