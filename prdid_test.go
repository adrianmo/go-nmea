package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPRDID(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  PRDID
	}{
		{
			name: "good sentence",
			raw:  "$PRDID,-10.37,2.34,230.34*62",
			msg: PRDID{
				Pitch:   -10.37,
				Roll:    2.34,
				Heading: 230.34,
			},
		},
		{
			name: "invalid Pitch",
			raw:  "$PRDID,x.37,2.34,230.34*36",
			err:  "nmea: PRDID invalid pitch: x.37",
		},
		{
			name: "invalid Roll",
			raw:  "$PRDID,-10.37,x.34,230.34*28",
			err:  "nmea: PRDID invalid roll: x.34",
		},
		{
			name: "invalid Heading",
			raw:  "$PRDID,-10.37,2.34,x.34*2B",
			err:  "nmea: PRDID invalid heading: x.34",
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
				hdt := m.(PRDID)
				hdt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hdt)
			}
		})
	}
}
