package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPHTRO(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  PHTRO
	}{
		{
			name: "good sentence",
			raw:  "$PHTRO,10.37,P,177.62,T*65",
			msg: PHTRO{
				Pitch: 10.37,
				Bow:   PHTROBowDown,
				Roll:  177.62,
				Port:  PHTROPortUP,
			},
		},
		{
			name: "invalid Pitch",
			raw:  "$PHTRO,x,P,177.62,T*36",
			err:  "nmea: PHTRO invalid pitch: x",
		},
		{
			name: "invalid Bow",
			raw:  "$PHTRO,10.37,x,177.62,T*4D",
			err:  "nmea: PHTRO invalid bow: x",
		},
		{
			name: "invalid Roll",
			raw:  "$PHTRO,10.37,P,x,T*06",
			err:  "nmea: PHTRO invalid roll: x",
		},
		{
			name: "invalid Port",
			raw:  "$PHTRO,10.37,P,177.62,x*49",
			err:  "nmea: PHTRO invalid port: x",
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
				hdt := m.(PHTRO)
				hdt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hdt)
			}
		})
	}
}
