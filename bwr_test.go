package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBWR(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  BWR
	}{
		{
			name: "good sentence",
			raw:  "$GPBWR,220516,5130.02,N,00046.34,W,213.8,T,218.0,M,0004.6,N,EGLM*30",
			msg: BWR{
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
			raw:  "$GPBWR,081837,,,,,,T,,M,,N,*02",
			msg: BWR{
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
			raw:  "$GPBWR,220516,5130.02,N,00046.34,W,213.8,T,218.0,M,0004.6,N,EGLM,D*58",
			msg: BWR{
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
			raw:  "$GPBWR,2x0516,5130.02,N,00046.34,W,213.8,T,218.0,M,0004.6,N,EGLM*7A",
			err:  "nmea: GPBWR invalid time: 2x0516",
		},
		{
			name: "invalid nmea: BearingTrueType",
			raw:  "$GPBWR,220516,5130.02,N,00046.34,W,213.8,M,218.0,M,0004.6,N,EGLM*29",
			err:  "nmea: GPBWR invalid true bearing type: M",
		},
		{
			name: "invalid nmea: BearingMagneticType",
			raw:  "$GPBWR,220516,5130.02,N,00046.34,W,213.8,T,218.0,T,0004.6,N,EGLM*29",
			err:  "nmea: GPBWR invalid magnetic bearing type: T",
		},
		{
			name: "invalid nmea: DistanceNauticalMilesUnit",
			raw:  "$GPBWR,220516,5130.02,N,00046.34,W,213.8,T,218.0,M,0004.6,K,EGLM*35",
			err:  "nmea: GPBWR invalid is distance to waypoint nautical miles unit: K",
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
				hdt := m.(BWR)
				hdt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hdt)
			}
		})
	}
}
