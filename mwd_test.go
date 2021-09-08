package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var mwdtests = []struct {
	name string
	raw  string
	err  string
	msg  MWD
}{
	{
		name: "good sentence",
		raw:  "$WIMWD,10.1,T,10.1,M,12,N,40,M*5D",
		msg: MWD{
			WindDirectionTrue:     10.1,
			TrueValid:             true,
			WindDirectionMagnetic: 10.1,
			MagneticValid:         true,
			WindSpeedKnots:        12,
			KnotsValid:            true,
			WindSpeedMeters:       40,
			MetersValid:           true,
		},
	},
	{
		name: "empty data",
		raw:  "$WIMWD,,,,,,,,*40",
		msg: MWD{
			WindDirectionTrue:     0,
			TrueValid:             false,
			WindDirectionMagnetic: 0,
			MagneticValid:         false,
			WindSpeedKnots:        0,
			KnotsValid:            false,
			WindSpeedMeters:       0,
			MetersValid:           false,
		},
	},
}

func TestMWD(t *testing.T) {
	for _, tt := range mwdtests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				mwd := m.(MWD)
				mwd.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, mwd)
			}
		})
	}
}
