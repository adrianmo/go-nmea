package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var glltests = []struct {
	name string
	raw  string
	err  string
	msg  GLL
}{
	{
		name: "good sentence",
		raw:  "$GPGLL,3926.7952,N,12000.5947,W,022732,A,A*58",
		msg: GLL{
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
			FFAMode:  FAAModeAutonomous,
		},
	},
	{
		name: "good sentence without FAA mode",
		raw:  "$IIGLL,5924.462,N,01030.048,E,062216,A*38",
		msg: GLL{
			Latitude:  MustParseLatLong("5924.462 N"),
			Longitude: MustParseLatLong("01030.048 E"),
			Time: Time{
				Valid:       true,
				Hour:        6,
				Minute:      22,
				Second:      16,
				Millisecond: 0,
			},
			Validity: "A",
			FFAMode:  "",
		},
	},
	{
		name: "bad validity",
		raw:  "$GPGLL,3926.7952,N,12000.5947,W,022732,D,A*5D",
		err:  "nmea: GPGLL invalid validity: D",
	},
}

func TestGLL(t *testing.T) {
	for _, tt := range glltests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				gll := m.(GLL)
				gll.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, gll)
			}
		})
	}
}
