package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pklshtests = []struct {
	name string
	raw  string
	err  string
	msg  PKLSH
}{
	{
		name: "good sentence",
		raw:  "$PKLSH,3926.7952,N,12000.5947,W,022732,A,100,2000*1A",
		msg: PKLSH{
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
			Fleet:    "100",
			UnitID:   "2000",
		},
	},
	{
		name: "bad validity",
		raw:  "$PKLSH,3926.7952,N,12000.5947,W,022732,D,100,2000*1F",
		err:  "nmea: PKLSH invalid validity: D",
	},
}

func TestPKLSH(t *testing.T) {
	for _, tt := range pklshtests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				pklsh := m.(PKLSH)
				pklsh.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, pklsh)
			}
		})
	}
}
