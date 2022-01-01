package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAAM(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  AAM
	}{
		{
			name: "good sentence",
			raw:  "$GPAAM,A,A,0.10,N,WPTNME*32",
			msg: AAM{
				StatusArrivalCircleEntered: WPStatusArrivalCircleEnteredA,
				StatusPerpendicularPassed:  WPStatusPerpendicularPassedA,
				ArrivalCircleRadius:        0.1,
				ArrivalCircleRadiusUnit:    DistanceUnitNauticalMile,
				DestinationWaypointID:      "WPTNME",
			},
		},
		{
			name: "invalid nmea: StatusArrivalCircleEntered",
			raw:  "$GPAAM,x,A,0.10,N,WPTNME*0B",
			err:  "nmea: GPAAM invalid arrival circle entered status: x",
		},
		{
			name: "invalid nmea: StatusPerpendicularPassed",
			raw:  "$GPAAM,A,x,0.10,N,WPTNME*0B",
			err:  "nmea: GPAAM invalid perpendicularly passed status: x",
		},
		{
			name: "invalid nmea: DistanceUnitNauticalMile",
			raw:  "$GPAAM,A,A,0.10,x,WPTNME*04",
			err:  "nmea: GPAAM invalid arrival circle radius units: x",
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
				hdt := m.(AAM)
				hdt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hdt)
			}
		})
	}
}
