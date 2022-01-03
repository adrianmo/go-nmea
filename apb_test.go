package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAPB(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  APB
	}{
		{
			name: "good sentence",
			raw:  "$GPAPB,A,A,0.10,R,N,V,V,011,M,DEST,011,M,011,M*3C",
			msg: APB{
				StatusGeneralWarning:       "A",
				StatusLockWarning:          "A",
				CrossTrackErrorMagnitude:   0.1,
				DirectionToSteer:           "R",
				CrossTrackUnits:            "N",
				StatusArrivalCircleEntered: "V",
				StatusPerpendicularPassed:  "V",
				BearingOriginToDest:        11,
				BearingOriginToDestType:    "M",
				DestinationWaypointID:      "DEST",
				BearingPresentToDest:       11,
				BearingPresentToDestType:   "M",
				Heading:                    11,
				HeadingType:                "M",
				FFAMode:                    "",
			},
		},
		{
			name: "good sentence b with FAA mode",
			raw:  "$ECAPB,A,A,0.0,L,M,V,V,175.2,T,Antechamber_Bay,175.2,T,175.2,T,V*32",
			msg: APB{
				StatusGeneralWarning:       "A",
				StatusLockWarning:          "A",
				CrossTrackErrorMagnitude:   0,
				DirectionToSteer:           "L",
				CrossTrackUnits:            "M",
				StatusArrivalCircleEntered: "V",
				StatusPerpendicularPassed:  "V",
				BearingOriginToDest:        175.2,
				BearingOriginToDestType:    "T",
				DestinationWaypointID:      "Antechamber_Bay",
				BearingPresentToDest:       175.2,
				BearingPresentToDestType:   "T",
				Heading:                    175.2,
				HeadingType:                "T",
				FFAMode:                    "V",
			},
		},
		{
			name: "invalid nmea: CrossTrackErrorMagnitude",
			raw:  "$ECAPB,A,A,x.0,L,M,V,V,175.2,T,Antechamber_Bay,175.2,T,175.2,T,V*7A",
			err:  "nmea: ECAPB invalid cross track error magnitude: x.0",
		},
		{
			name: "invalid nmea: BearingOriginToDest",
			raw:  "$ECAPB,A,A,0.0,L,M,V,V,175.x,T,Antechamber_Bay,175.2,T,175.2,T,V*78",
			err:  "nmea: ECAPB invalid origin bearing to destination: 175.x",
		},
		{
			name: "invalid nmea: BearingPresentToDest",
			raw:  "$ECAPB,A,A,0.0,L,M,V,V,175.2,T,Antechamber_Bay,175.x,T,175.2,T,V*78",
			err:  "nmea: ECAPB invalid present bearing to destination: 175.x",
		},
		{
			name: "invalid nmea: Heading",
			raw:  "$ECAPB,A,A,0.0,L,M,V,V,175.2,T,Antechamber_Bay,175.2,T,175.x,T,V*78",
			err:  "nmea: ECAPB invalid heading: 175.x",
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
				apb := m.(APB)
				apb.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, apb)
			}
		})
	}
}
