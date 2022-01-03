package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHDG(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  HDG
	}{
		{
			name: "good sentence",
			raw:  "$HCHDG,98.3,0.1,E,12.6,W*56",
			msg: HDG{
				Heading:            98.3,
				Deviation:          0.1,
				DeviationDirection: East,
				Variation:          12.6,
				VariationDirection: West,
			},
		},
		{
			name: "invalid Heading",
			raw:  "$HCHDG,X,0.1,E,12.6,W*12",
			err:  "nmea: HCHDG invalid heading: X",
		},
		{
			name: "invalid Deviation",
			raw:  "$HCHDG,98.3,x.1,E,12.6,W*1E",
			err:  "nmea: HCHDG invalid deviation: x.1",
		},
		{
			name: "invalid DeviationDirection",
			raw:  "$HCHDG,98.3,0.1,X,12.6,W*4B",
			err:  "nmea: HCHDG invalid deviation direction: X",
		},
		{
			name: "invalid Variation",
			raw:  "$HCHDG,98.3,0.1,E,x.1,W*2A",
			err:  "nmea: HCHDG invalid variation: x.1",
		},
		{
			name: "invalid VariationDirection",
			raw:  "$HCHDG,98.3,0.1,E,12.6,X*59",
			err:  "nmea: HCHDG invalid variation direction: X",
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
				hdg := m.(HDG)
				hdg.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hdg)
			}
		})
	}
}
