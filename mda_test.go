package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var mdatests = []struct {
	name string
	raw  string
	err  string
	msg  MDA
}{
	{
		name: "good sentence",
		raw:  "$WIMDA,3.02,I,1.01,B,23.4,C,,,40.2,,12.1,C,19.3,T,20.1,M,13.1,N,1.1,M*62",
		msg: MDA{
			PressureInch:          3.02,
			InchesValid:           true,
			PressureBar:           1.01,
			BarsValid:             true,
			AirTemp:               23.4,
			AirTempValid:          true,
			WaterTemp:             0,
			WaterTempValid:        false,
			RelativeHum:           40.2,
			AbsoluteHum:           0,
			DewPoint:              12.1,
			DewPointValid:         true,
			WindDirectionTrue:     19.3,
			TrueValid:             true,
			WindDirectionMagnetic: 20.1,
			MagneticValid:         true,
			WindSpeedKnots:        13.1,
			KnotsValid:            true,
			WindSpeedMeters:       1.1,
			MetersValid:           true,
		},
	},
}

func TestMDA(t *testing.T) {
	for _, tt := range mdatests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				mda := m.(MDA)
				mda.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, mda)
			}
		})
	}
}
