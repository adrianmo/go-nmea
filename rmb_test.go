package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRMB(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  RMB
	}{
		{
			name: "good sentence",
			raw:  "$GPRMB,A,0.66,L,003,004,4917.24,N,12309.57,W,001.3,052.5,000.5,V*20",
			msg: RMB{
				DataStatus:                      DataStatusWarningClearRMB,
				CrossTrackErrorNauticalMiles:    0.66,
				DirectionToSteer:                Left,
				OriginWaypointID:                "003",
				DestinationWaypointID:           "004",
				DestinationLatitude:             49.28733333333333,
				DestinationLongitude:            -123.1595,
				RangeToDestinationNauticalMiles: 1.3,
				TrueBearingToDestination:        52.5,
				VelocityToDestinationKnots:      0.5,
				ArrivalStatus:                   WPStatusArrivalCircleEnteredV,
				FFAMode:                         "",
			},
		},
		{
			name: "good sentence with FAAMode",
			raw:  "$GPRMB,A,0.66,L,003,004,4917.24,N,12309.57,W,001.3,052.5,000.5,V,D*48",
			msg: RMB{
				DataStatus:                      DataStatusWarningClearRMB,
				CrossTrackErrorNauticalMiles:    0.66,
				DirectionToSteer:                Left,
				OriginWaypointID:                "003",
				DestinationWaypointID:           "004",
				DestinationLatitude:             49.28733333333333,
				DestinationLongitude:            -123.1595,
				RangeToDestinationNauticalMiles: 1.3,
				TrueBearingToDestination:        52.5,
				VelocityToDestinationKnots:      0.5,
				ArrivalStatus:                   WPStatusArrivalCircleEnteredV,
				FFAMode:                         FAAModeDifferential,
			},
		},
		{
			name: "invalid nmea: DataStatus",
			raw:  "$GPRMB,x,0.66,L,003,004,4917.24,N,12309.57,W,001.3,052.5,000.5,V,D*71",
			err:  "nmea: GPRMB invalid data status: x",
		},
		{
			name: "invalid nmea: CrossTrackErrorNauticalMiles",
			raw:  "$GPRMB,A,x.66,L,003,004,4917.24,N,12309.57,W,001.3,052.5,000.5,V,D*00",
			err:  "nmea: GPRMB invalid cross track error: x.66",
		},
		{
			name: "invalid nmea: DirectionToSteer",
			raw:  "$GPRMB,A,0.66,x,003,004,4917.24,N,12309.57,W,001.3,052.5,000.5,V,D*7C",
			err:  "nmea: GPRMB invalid direction to steer: x",
		},
		{
			name: "invalid nmea: DestinationLatitude",
			raw:  "$GPRMB,A,0.66,L,003,004,4x17.24,N,12309.57,W,001.3,052.5,000.5,V,D*09",
			err:  "nmea: GPRMB invalid latitude: cannot parse [4x17.24 N], unknown format",
		},
		{
			name: "invalid nmea: DestinationLongitude",
			raw:  "$GPRMB,A,0.66,L,003,004,4917.24,N,12x09.57,W,001.3,052.5,000.5,V,D*03",
			err:  "nmea: GPRMB invalid latitude: cannot parse [12x09.57 W], unknown format",
		},
		{
			name: "invalid nmea: RangeToDestinationNauticalMiles",
			raw:  "$GPRMB,A,0.66,L,003,004,4917.24,N,12309.57,W,x01.3,052.5,000.5,V,D*00",
			err:  "nmea: GPRMB invalid range to destination: x01.3",
		},
		{
			name: "invalid nmea: TrueBearingToDestination",
			raw:  "$GPRMB,A,0.66,L,003,004,4917.24,N,12309.57,W,001.3,052.x,000.5,V,D*05",
			err:  "nmea: GPRMB invalid true bearing to destination: 052.x",
		},
		{
			name: "invalid nmea: VelocityToDestinationKnots",
			raw:  "$GPRMB,A,0.66,L,003,004,4917.24,N,12309.57,W,001.3,052.5,000.x,V,D*05",
			err:  "nmea: GPRMB invalid velocity to destination: 000.x",
		},
		{
			name: "invalid nmea: ArrivalStatus",
			raw:  "$GPRMB,A,0.66,L,003,004,4917.24,N,12309.57,W,001.3,052.5,000.5,x,D*66",
			err:  "nmea: GPRMB invalid arrival status: x",
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
				hdt := m.(RMB)
				hdt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hdt)
			}
		})
	}
}
