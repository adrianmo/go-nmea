package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVPW(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  VPW
	}{
		{
			name: "good sentence",
			raw:  "$IIVPW,4.5,N,6.7,M*52",
			msg: VPW{
				SpeedKnots:     4.5,
				SpeedKnotsUnit: SpeedKnots,
				SpeedMPS:       6.7,
				SpeedMPSUnit:   SpeedMeterPerSecond,
			},
		},
		{
			name: "invalid nmea: SpeedKnotsUnit",
			raw:  "$IIVPW,4.5,x,6.7,M*64",
			err:  "nmea: IIVPW invalid wind speed in knots unit: x",
		},
		{
			name: "invalid nmea: SpeedMPSUnit",
			raw:  "$IIVPW,4.5,N,6.7,x*67",
			err:  "nmea: IIVPW invalid wind speed in meters per second unit: x",
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
				hdt := m.(VPW)
				hdt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hdt)
			}
		})
	}
}
