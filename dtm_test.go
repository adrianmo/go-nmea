package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDTM(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  DTM
	}{
		{
			name: "good sentence 1",
			raw:  "$GPDTM,W84,,0.0,N,0.0,E,0.0,W84*6F",
			msg: DTM{
				BaseSentence:          BaseSentence{},
				LocalDatumCode:        "W84",
				LocalDatumSubcode:     "",
				LatitudeOffsetMinute:  0,
				LongitudeOffsetMinute: 0,
				AltitudeOffsetMeters:  0,
				DatumName:             "W84",
			},
		},
		{
			name: "good sentence 2",
			raw:  "$GPDTM,W84,X,00.1200,S,12.0000,W,100,W84*27",
			msg: DTM{
				BaseSentence:          BaseSentence{},
				LocalDatumCode:        "W84",
				LocalDatumSubcode:     "X",
				LatitudeOffsetMinute:  -0.12,
				LongitudeOffsetMinute: -12,
				AltitudeOffsetMeters:  100,
				DatumName:             "W84",
			},
		},
		{
			name: "invalid nmea: LatitudeOffsetMinute",
			raw:  "$GPDTM,W84,,x,N,0.0,E,0.0,W84*39",
			err:  "nmea: GPDTM invalid latitude offset minutes: x",
		},
		{
			name: "invalid nmea: LongitudeOffsetMinute",
			raw:  "$GPDTM,W84,,0.0,N,x,E,0.0,W84*39",
			err:  "nmea: GPDTM invalid longitude offset minutes: x",
		},
		{
			name: "invalid nmea: AltitudeOffsetMeters",
			raw:  "$GPDTM,W84,,0.0,N,0.0,E,x,W84*39",
			err:  "nmea: GPDTM invalid altitude offset offset: x",
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
				mm := m.(DTM)
				mm.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, mm)
			}
		})
	}
}
