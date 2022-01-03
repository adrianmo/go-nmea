package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVWR(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  VWR
	}{
		{ // these examples are from SignalK
			name: "good sentence",
			raw:  "$IIVWR,75,R,1.0,N,0.51,M,1.85,K*6C",
			msg: VWR{
				MeasuredAngle:        75,
				MeasuredDirectionBow: Right,
				SpeedKnots:           1,
				SpeedKnotsUnit:       SpeedKnots,
				SpeedMPS:             0.51,
				SpeedMPSUnit:         SpeedMeterPerSecond,
				SpeedKPH:             1.85,
				SpeedKPHUnit:         SpeedKilometerPerHour,
			},
		},
		{
			name: "good sentence, shorter but still valid",
			raw:  "$IIVWR,024,L,018,N,,,,*5e",
			msg: VWR{
				MeasuredAngle:        24,
				MeasuredDirectionBow: Left,
				SpeedKnots:           18,
				SpeedKnotsUnit:       SpeedKnots,
				SpeedMPS:             0,
				SpeedMPSUnit:         "",
				SpeedKPH:             0,
				SpeedKPHUnit:         "",
			},
		},
		{
			name: "good sentence, handle empty values",
			raw:  "$IIVWR,,,,,,,,*53",
			msg: VWR{
				MeasuredAngle:        0,
				MeasuredDirectionBow: "",
				SpeedKnots:           0,
				SpeedKnotsUnit:       "",
				SpeedMPS:             0,
				SpeedMPSUnit:         "",
				SpeedKPH:             0,
				SpeedKPHUnit:         "",
			},
		},
		{
			name: "invalid nmea: DirectionBow",
			raw:  "$IIVWR,75,x,1.0,N,0.51,M,1.85,K*46",
			err:  "nmea: IIVWR invalid measured wind direction to bow: x",
		},
		{
			name: "invalid nmea: SpeedKnotsUnit",
			raw:  "$IIVWR,75,R,1.0,x,0.51,M,1.85,K*5a",
			err:  "nmea: IIVWR invalid wind speed in knots unit: x",
		},
		{
			name: "invalid nmea: SpeedMPSUnit",
			raw:  "$IIVWR,75,R,1.0,N,0.51,x,1.85,K*59",
			err:  "nmea: IIVWR invalid wind speed in meters per second unit: x",
		},
		{
			name: "invalid nmea: SpeedKPHUnit",
			raw:  "$IIVWR,75,R,1.0,N,0.51,M,1.85,x*5f",
			err:  "nmea: IIVWR invalid wind speed in kilometers per hour unit: x",
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
				vwr := m.(VWR)
				vwr.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, vwr)
			}
		})
	}
}
