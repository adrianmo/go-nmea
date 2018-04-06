package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var gnggatests = []struct {
	name string
	raw  string
	err  string
	msg  GNGGA
}{
	{
		name: "good sentence",
		raw:  "$GNGGA,203415.000,6325.6138,N,01021.4290,E,1,8,2.42,72.5,M,41.5,M,,*7C",
		msg: GNGGA{
			Time: Time{
				Valid:       true,
				Hour:        20,
				Minute:      34,
				Second:      15,
				Millisecond: 0,
			},
			Latitude:      MustParseLatLong("6325.6138 N"),
			Longitude:     MustParseLatLong("01021.4290 E"),
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
		name: "bad type",
		raw:  "$GPRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*70",
		err:  "nmea: GNGGA invalid prefix: GPRMC",
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
}

func TestGNGGA(t *testing.T) {
	for _, tt := range gnggatests {
		t.Run(tt.name, func(t *testing.T) {
			sent, err := ParseSentence(tt.raw)
			assert.NoError(t, err)
			gngga, err := NewGNGGA(sent)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				gngga.Sent = Sent{}
				assert.Equal(t, tt.msg, gngga)
			}
		})
	}
}
