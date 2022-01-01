package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVWT(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  VWT
	}{
		{ // these examples are from SignalK
			name: "good sentence",
			raw:  "$IIVWT,75,R,1.0,N,0.51,M,1.85,K*6A",
			msg: VWT{
				TrueAngle:        75,
				TrueDirectionBow: Right,
				SpeedKnots:       1,
				SpeedKnotsUnit:   SpeedKnots,
				SpeedMPS:         0.51,
				SpeedMPSUnit:     SpeedMeterPerSecond,
				SpeedKPH:         1.85,
				SpeedKPHUnit:     SpeedKilometerPerHour,
			},
		},
		{
			name: "good sentence, shorter but still valid",
			raw:  "$IIVWT,024,L,018,N,,,,*58",
			msg: VWT{
				TrueAngle:        24,
				TrueDirectionBow: Left,
				SpeedKnots:       18,
				SpeedKnotsUnit:   SpeedKnots,
				SpeedMPS:         0,
				SpeedMPSUnit:     "",
				SpeedKPH:         0,
				SpeedKPHUnit:     "",
			},
		},
		{
			name: "good sentence, handle empty values",
			raw:  "$IIVWT,,,,,,,,*55",
			msg: VWT{
				TrueAngle:        0,
				TrueDirectionBow: "",
				SpeedKnots:       0,
				SpeedKnotsUnit:   "",
				SpeedMPS:         0,
				SpeedMPSUnit:     "",
				SpeedKPH:         0,
				SpeedKPHUnit:     "",
			},
		},
		{
			name: "invalid nmea: DirectionBow",
			raw:  "$IIVWT,75,x,1.0,N,0.51,M,1.85,K*40",
			err:  "nmea: IIVWT invalid true wind direction to bow: x",
		},
		{
			name: "invalid nmea: SpeedKnotsUnit",
			raw:  "$IIVWT,75,R,1.0,x,0.51,M,1.85,K*5c",
			err:  "nmea: IIVWT invalid wind speed in knots unit: x",
		},
		{
			name: "invalid nmea: SpeedMPSUnit",
			raw:  "$IIVWT,75,R,1.0,N,0.51,x,1.85,K*5f",
			err:  "nmea: IIVWT invalid wind speed in meters per second unit: x",
		},
		{
			name: "invalid nmea: SpeedKPHUnit",
			raw:  "$IIVWT,75,R,1.0,N,0.51,M,1.85,x*59",
			err:  "nmea: IIVWT invalid wind speed in kilometers per hour unit: x",
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
				hdt := m.(VWT)
				hdt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hdt)
			}
		})
	}
}
