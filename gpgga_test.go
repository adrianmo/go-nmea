package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var gpggatests = []struct {
	name string
	raw  string
	err  string
	msg  GPGGA
}{
	{
		name: "good sentence",
		raw:  "$GPGGA,034225.077,3356.4650,S,15124.5567,E,1,03,9.7,-25.0,M,21.0,M,,0000*51",
		msg: GPGGA{
			Time:          Time{true, 3, 42, 25, 77},
			Latitude:      MustParseLatLong("3356.4650 S"),
			Longitude:     MustParseLatLong("15124.5567 E"),
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
		name: "wrong type",
		raw:  "$GPRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*70",
		err:  "nmea: GPGGA invalid prefix: GPRMC",
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

func TestGPGGA(t *testing.T) {
	for _, tt := range gpggatests {
		t.Run(tt.name, func(t *testing.T) {
			sent, err := ParseSentence(tt.raw)
			assert.NoError(t, err)
			gpgga, err := NewGPGGA(sent)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				gpgga.Sent = Sent{}
				assert.Equal(t, tt.msg, gpgga)
			}
		})
	}
}
