package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pknshtests = []struct {
	name string
	raw  string
	err  string
	msg  PKNSH
}{
	{
		name: "good sentence",
		raw:  "$PKNSH,3926.7952,N,12000.5947,W,022732,A,U00001*63",
		msg: PKNSH{
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
			UnitID:   "U00001",
		},
	},
	{
		name: "bad validity",
		raw:  "$PKNSH,3926.7952,N,12000.5947,W,022732,D,U00001*66",
		err:  "nmea: PKNSH invalid validity: D",
	},
}

func TestPKNSH(t *testing.T) {
	for _, tt := range pknshtests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				pknsh := m.(PKNSH)
				pknsh.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, pknsh)
			}
		})
	}
}
