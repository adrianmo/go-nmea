package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBWW(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  BWW
	}{
		{
			name: "good sentence",
			raw:  "$GPBWW,097.0,T,103.2,M,POINTB,POINTA*41",
			msg: BWW{
				BearingTrue:           97.0,
				BearingTrueType:       BearingTrue,
				BearingMagnetic:       103.2,
				BearingMagneticType:   BearingMagnetic,
				DestinationWaypointID: "POINTB",
				OriginWaypointID:      "POINTA",
			},
		},
		{
			name: "invalid nmea: BearingTrueValid",
			raw:  "$GPBWW,097.0,M,103.2,M,POINTB,POINTA*58",
			err:  "nmea: GPBWW invalid true bearing type: M",
		},
		{
			name: "invalid nmea: BearingMagneticValid",
			raw:  "$GPBWW,097.0,T,103.2,T,POINTB,POINTA*58",
			err:  "nmea: GPBWW invalid magnetic bearing type: T",
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
				bww := m.(BWW)
				bww.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, bww)
			}
		})
	}
}
