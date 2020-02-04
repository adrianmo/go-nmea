package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var vhw = []struct {
	name string
	raw  string
	err  string
	msg  VHW
}{
	{
		name: "good sentence",
		raw:  "$VWVHW,45.0,T,43.0,M,3.5,N,6.4,K*56",
		msg: VHW{
			TrueHeading:            45.0,
			MagneticHeading:        43.0,
			SpeedThroughWaterKnots: 3.5,
			SpeedThroughWaterKPH:   6.4,
		},
	},
	{
		name: "bad sentence",
		raw:  "$VWVHW,T,45.0,43.0,M,3.5,N,6.4,K*56",
		err:  "nmea: VWVHW invalid true heading: T",
	},
}

func TestVHW(t *testing.T) {
	for _, tt := range vhw {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				vhw := m.(VHW)
				vhw.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, vhw)
			}
		})
	}
}
