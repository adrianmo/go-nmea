package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPSONCMS(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  PSONCMS
	}{
		{
			name: "good sentence",
			raw:  "$PSONCMS,0.0905,0.4217,0.9020,-0.0196,-1.7685,0.3861,-9.6648,-0.0116,0.0065,-0.0080,0.0581,0.3846,0.7421,33.1*76",
			msg: PSONCMS{
				BaseSentence:      BaseSentence{},
				Quaternion0:       0.0905,
				Quaternion1:       0.4217,
				Quaternion2:       0.9020,
				Quaternion3:       -0.0196,
				AccelerationX:     -1.7685,
				AccelerationY:     0.3861,
				AccelerationZ:     -9.6648,
				RateOfTurnX:       -0.0116,
				RateOfTurnY:       0.0065,
				RateOfTurnZ:       -0.0080,
				MagneticFieldX:    0.0581,
				MagneticFieldY:    0.3846,
				MagneticFieldZ:    0.7421,
				SensorTemperature: 33.1,
			},
		},
		{
			name: "invalid Quaternion0",
			raw:  "$PSONCMS,x,0.4217,0.9020,-0.0196,-1.7685,0.3861,-9.6648,-0.0116,0.0065,-0.0080,0.0581,0.3846,0.7421,33.1*1C",
			err:  "nmea: PSONCMS invalid q0 from quaternions: x",
		},
		{
			name: "invalid Quaternion1",
			raw:  "$PSONCMS,0.0905,x,0.9020,-0.0196,-1.7685,0.3861,-9.6648,-0.0116,0.0065,-0.0080,0.0581,0.3846,0.7421,33.1*10",
			err:  "nmea: PSONCMS invalid q1 from quaternions: x",
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
				hdt := m.(PSONCMS)
				hdt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hdt)
			}
		})
	}
}
