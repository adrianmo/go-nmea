package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestROT(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  ROT
	}{
		{
			name: "good sentence",
			raw:  "$HEROT,-11.23,A*07",
			msg: ROT{
				RateOfTurn: -11.23,
				Valid:      true,
			},
		},
		{
			name: "invalid RateOfTurn",
			raw:  "$HEROT,x,A*7D",
			err:  "nmea: HEROT invalid rate of turn: x",
		},
		{
			name: "invalid Valid",
			raw:  "$HEROT,-11.23,X*1E",
			err:  "nmea: HEROT invalid status valid: X",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				rot := m.(ROT)
				rot.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, rot)
			}
		})
	}
}
