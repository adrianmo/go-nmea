package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMTW(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  MTW
	}{
		{
			name: "good sentence",
			raw:  "$INMTW,17.9,C*1B",
			msg: MTW{
				Temperature:  17.9,
				CelsiusValid: true,
			},
		},
		{
			name: "invalid Temperature",
			raw:  "$INMTW,x.9,C*65",
			err:  "nmea: INMTW invalid temperature: x.9",
		},
		{
			name: "invalid CelsiusValid",
			raw:  "$INMTW,17.9,x*20",
			err:  "nmea: INMTW invalid unit of measurement celsius: x",
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
				hdt := m.(MTW)
				hdt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hdt)
			}
		})
	}
}
