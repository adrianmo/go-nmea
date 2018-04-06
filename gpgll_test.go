package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var gpglltests = []struct {
	name string
	raw  string
	err  string
	msg  GPGLL
}{
	{
		name: "good sentence",
		raw:  "$GPGLL,3926.7952,N,12000.5947,W,022732,A,A*58",
		msg: GPGLL{
			Latitude:  MustParseLatLong("3926.7952 N"),
			Longitude: MustParseLatLong("12000.5947 W"),
			Time: Time{
				Valid:       true,
				Hour:        2,
				Minute:      27,
				Second:      32,
				Millisecond: 0,
			},
			Validity: "A",
		},
	},
	{
		name: "bad validity",
		raw:  "$GPGLL,3926.7952,N,12000.5947,W,022732,D,A*5D",
		err:  "nmea: GPGLL invalid validity: D",
	},
	{
		name: "wrong sentence",
		raw:  "$GPXTE,A,A,4.07,L,N*6D",
		err:  "nmea: GPGLL invalid prefix: GPXTE",
	},
}

func TestGPGLL(t *testing.T) {
	for _, tt := range gpglltests {
		t.Run(tt.name, func(t *testing.T) {
			sent, err := ParseSentence(tt.raw)
			assert.NoError(t, err)
			gpgll, err := NewGPGLL(sent)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				gpgll.Sent = Sent{}
				assert.Equal(t, tt.msg, gpgll)
			}
		})
	}
}
