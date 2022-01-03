package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBEC(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  BEC
	}{
		{
			name: "good sentence",
			raw:  "$GPBEC,220516,5130.02,N,00046.34,W,213.8,T,218.0,M,0004.6,N,EGLM*33",
			msg: BEC{
				Time: Time{
					Valid:       true,
					Hour:        22,
					Minute:      5,
					Second:      16,
					Millisecond: 0,
				},
				Latitude:                   51.50033333333334,
				Longitude:                  -0.7723333333333334,
				BearingTrue:                213.8,
				BearingTrueValid:           true,
				BearingMagnetic:            218,
				BearingMagneticValid:       true,
				DistanceNauticalMiles:      4.6,
				DistanceNauticalMilesValid: true,
				DestinationWaypointID:      "EGLM",
			},
		},
		{
			name: "invalid nmea: Time",
			raw:  "$GPBEC,2x0516,5130.02,N,00046.34,W,213.8,T,218.0,M,0004.6,N,EGLM*79",
			err:  "nmea: GPBEC invalid time: 2x0516",
		},
		{
			name: "invalid nmea: BearingTrueValid",
			raw:  "$GPBEC,220516,5130.02,N,00046.34,W,213.8,M,218.0,M,0004.6,N,EGLM*2A",
			err:  "nmea: GPBEC invalid true bearing unit valid: M",
		},
		{
			name: "invalid nmea: BearingMagneticValid",
			raw:  "$GPBEC,220516,5130.02,N,00046.34,W,213.8,T,218.0,T,0004.6,N,EGLM*2A",
			err:  "nmea: GPBEC invalid magnetic bearing unit valid: T",
		},
		{
			name: "invalid nmea: DistanceNauticalMilesValid",
			raw:  "$GPBEC,220516,5130.02,N,00046.34,W,213.8,T,218.0,M,0004.6,K,EGLM*36",
			err:  "nmea: GPBEC invalid is distance to waypoint nautical miles valid: K",
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
				bec := m.(BEC)
				bec.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, bec)
			}
		})
	}
}
