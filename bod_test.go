package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBOD(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  BOD
	}{
		{
			name: "good sentence with both WPs",
			raw:  "$GPBOD,097.0,T,103.2,M,POINTB,POINTA*4A",
			msg: BOD{
				BearingTrue:           97.0,
				BearingTrueType:       BearingTrue,
				BearingMagnetic:       103.2,
				BearingMagneticType:   BearingMagnetic,
				DestinationWaypointID: "POINTB",
				OriginWaypointID:      "POINTA",
			},
		},
		{
			name: "good sentence onyl destination",
			raw:  "$GPBOD,099.3,T,105.6,M,POINTB*64",
			msg: BOD{
				BearingTrue:           99.3,
				BearingTrueType:       BearingTrue,
				BearingMagnetic:       105.6,
				BearingMagneticType:   BearingMagnetic,
				DestinationWaypointID: "POINTB",
				OriginWaypointID:      "",
			},
		},
		{
			name: "invalid nmea: BearingTrueValid",
			raw:  "$GPBOD,097.0,M,103.2,M,POINTB,POINTA*53",
			err:  "nmea: GPBOD invalid true bearing type: M",
		},
		{
			name: "invalid nmea: BearingMagneticValid",
			raw:  "$GPBOD,097.0,T,103.2,T,POINTB,POINTA*53",
			err:  "nmea: GPBOD invalid magnetic bearing type: T",
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
				bod := m.(BOD)
				bod.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, bod)
			}
		})
	}
}
