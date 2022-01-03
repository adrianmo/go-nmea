package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBWC(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  BWC
	}{
		{
			name: "good sentence",
			raw:  "$GPBWC,220516,5130.02,N,00046.34,W,213.8,T,218.0,M,0004.6,N,EGLM*21",
			msg: BWC{
				Time: Time{
					Valid:       true,
					Hour:        22,
					Minute:      5,
					Second:      16,
					Millisecond: 0,
				},
				Latitude:                  51.50033333333334,
				Longitude:                 -0.7723333333333334,
				BearingTrue:               213.8,
				BearingTrueType:           BearingTrue,
				BearingMagnetic:           218,
				BearingMagneticType:       BearingMagnetic,
				DistanceNauticalMiles:     4.6,
				DistanceNauticalMilesUnit: DistanceUnitNauticalMile,
				DestinationWaypointID:     "EGLM",
				FFAMode:                   "",
			},
		},
		{
			name: "good sentence no waypoint",
			raw:  "$GPBWC,081837,,,,,,T,,M,,N,*13",
			msg: BWC{
				Time:                      Time{Valid: true, Hour: 8, Minute: 18, Second: 37, Millisecond: 0},
				Latitude:                  0,
				Longitude:                 0,
				BearingTrue:               0,
				BearingTrueType:           BearingTrue,
				BearingMagnetic:           0,
				BearingMagneticType:       BearingMagnetic,
				DistanceNauticalMiles:     0,
				DistanceNauticalMilesUnit: DistanceUnitNauticalMile,
				DestinationWaypointID:     "",
				FFAMode:                   "",
			},
		},
		{
			name: "good sentence with FAAMode",
			raw:  "$GPBWC,220516,5130.02,N,00046.34,W,213.8,T,218.0,M,0004.6,N,EGLM,D*49",
			msg: BWC{
				Time: Time{
					Valid:       true,
					Hour:        22,
					Minute:      5,
					Second:      16,
					Millisecond: 0,
				},
				Latitude:                  51.50033333333334,
				Longitude:                 -0.7723333333333334,
				BearingTrue:               213.8,
				BearingTrueType:           BearingTrue,
				BearingMagnetic:           218,
				BearingMagneticType:       BearingMagnetic,
				DistanceNauticalMiles:     4.6,
				DistanceNauticalMilesUnit: DistanceUnitNauticalMile,
				DestinationWaypointID:     "EGLM",
				FFAMode:                   FAAModeDifferential,
			},
		},
		{
			name: "invalid nmea: Time",
			raw:  "$GPBWC,2x0516,5130.02,N,00046.34,W,213.8,T,218.0,M,0004.6,N,EGLM*6B",
			err:  "nmea: GPBWC invalid time: 2x0516",
		},
		{
			name: "invalid nmea: BearingTrueValid",
			raw:  "$GPBWC,220516,5130.02,N,00046.34,W,213.8,M,218.0,M,0004.6,N,EGLM*38",
			err:  "nmea: GPBWC invalid true bearing type: M",
		},
		{
			name: "invalid nmea: BearingMagneticValid",
			raw:  "$GPBWC,220516,5130.02,N,00046.34,W,213.8,T,218.0,T,0004.6,N,EGLM*38",
			err:  "nmea: GPBWC invalid magnetic bearing type: T",
		},
		{
			name: "invalid nmea: DistanceNauticalMilesValid",
			raw:  "$GPBWC,220516,5130.02,N,00046.34,W,213.8,T,218.0,M,0004.6,K,EGLM*24",
			err:  "nmea: GPBWC invalid is distance to waypoint nautical miles unit: K",
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
				bwc := m.(BWC)
				bwc.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, bwc)
			}
		})
	}
}
