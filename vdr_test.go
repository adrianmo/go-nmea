package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVDR(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  VDR
	}{
		{
			name: "good sentence",
			raw:  "$IIVDR,10.1,T,12.3,M,1.2,N*3A",
			msg: VDR{
				SetDegreesTrue:         10.1,
				SetDegreesTrueUnit:     BearingTrue,
				SetDegreesMagnetic:     12.3,
				SetDegreesMagneticUnit: BearingMagnetic,
				DriftKnots:             1.2,
				DriftUnit:              SpeedKnots,
			},
		},
		{
			name: "invalid nmea: SetDegreesTrueUnit",
			raw:  "$IIVDR,10.1,x,12.3,M,1.2,N*16",
			err:  "nmea: IIVDR invalid true set unit: x",
		},
		{
			name: "invalid nmea: SetDegreesMagneticUnit",
			raw:  "$IIVDR,10.1,T,12.3,x,1.2,N*0f",
			err:  "nmea: IIVDR invalid magnetic set unit: x",
		},
		{
			name: "invalid nmea: DriftUnit",
			raw:  "$IIVDR,10.1,T,12.3,M,1.2,x*0c",
			err:  "nmea: IIVDR invalid drift unit: x",
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
				vdr := m.(VDR)
				vdr.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, vdr)
			}
		})
	}
}
