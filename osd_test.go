package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOSD(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  OSD
	}{
		{
			name: "good sentence",
			raw:  "$RAOSD,179.0,A,179.0,M,00.0,M,,,N*76",
			msg: OSD{
				BaseSentence:     BaseSentence{},
				Heading:          179,
				HeadingStatus:    "A",
				VesselTrueCourse: 179,
				CourseReference:  "M",
				VesselSpeed:      0,
				SpeedReference:   "M",
				VesselSetTrue:    0,
				VesselDrift:      0,
				SpeedUnits:       "N",
			},
		},
		{
			name: "invalid nmea: Heading",
			raw:  "$RAOSD,x179.0,A,179.0,M,00.0,M,,,N*0e",
			err:  "nmea: RAOSD invalid heading: x179.0",
		},
		{
			name: "invalid nmea: HeadingStatus",
			raw:  "$RAOSD,179.0,xA,179.0,M,00.0,M,,,N*0e",
			err:  "nmea: RAOSD invalid heading status: xA",
		},
		{
			name: "invalid nmea: VesselTrueCourse",
			raw:  "$RAOSD,179.0,A,x179.0,M,00.0,M,,,N*0e",
			err:  "nmea: RAOSD invalid vessel course true: x179.0",
		},
		{
			name: "invalid nmea: CourseReference",
			raw:  "$RAOSD,179.0,A,179.0,xM,00.0,M,,,N*0e",
			err:  "nmea: RAOSD invalid course reference: xM",
		},
		{
			name: "invalid nmea: VesselSpeed",
			raw:  "$RAOSD,179.0,A,179.0,M,x00.0,M,,,N*0e",
			err:  "nmea: RAOSD invalid vessel speed: x00.0",
		},
		{
			name: "invalid nmea: SpeedReference",
			raw:  "$RAOSD,179.0,A,179.0,M,00.0,xM,,,N*0e",
			err:  "nmea: RAOSD invalid speed reference: xM",
		},
		{
			name: "invalid nmea: VesselSetTrue",
			raw:  "$RAOSD,179.0,A,179.0,M,00.0,M,x,,N*0e",
			err:  "nmea: RAOSD invalid vessel set: x",
		},
		{
			name: "invalid nmea: VesselDrift",
			raw:  "$RAOSD,179.0,A,179.0,M,00.0,M,,x,N*0e",
			err:  "nmea: RAOSD invalid vessel drift: x",
		},
		{
			name: "invalid nmea: SpeedUnits",
			raw:  "$RAOSD,179.0,A,179.0,M,00.0,M,,,xN*0e",
			err:  "nmea: RAOSD invalid speed units: xN",
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
				mm := m.(OSD)
				mm.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, mm)
			}
		})
	}
}
