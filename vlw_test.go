package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVLW(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  VLW
	}{
		{
			name: "good sentence 1",
			raw:  "$IIVLW,10.1,N,3.2,N*7C",
			msg: VLW{
				TotalInWater:           10.1,
				TotalInWaterUnit:       "N",
				SinceResetInWater:      3.2,
				SinceResetInWaterUnit:  "N",
				TotalOnGround:          0,
				TotalOnGroundUnit:      "",
				SinceResetOnGround:     0,
				SinceResetOnGroundUnit: "",
			},
		},
		{
			name: "good sentence 2",
			raw:  "$IIVLW,10.1,N,3.2,N,1,N,0.1,N*62",
			msg: VLW{
				TotalInWater:           10.1,
				TotalInWaterUnit:       "N",
				SinceResetInWater:      3.2,
				SinceResetInWaterUnit:  "N",
				TotalOnGround:          1,
				TotalOnGroundUnit:      "N",
				SinceResetOnGround:     0.1,
				SinceResetOnGroundUnit: "N",
			},
		},
		{
			name: "invalid nmea: TotalInWaterUnit",
			raw:  "$IIVLW,10.1,x,3.2,N,1,N,0.1,N*54",
			err:  "nmea: IIVLW invalid total cumulative water distance unit: x",
		},
		{
			name: "invalid nmea: SinceResetInWaterUnit",
			raw:  "$IIVLW,10.1,N,3.2,x,1,N,0.1,N*54",
			err:  "nmea: IIVLW invalid water distance since reset unit: x",
		},
		{
			name: "invalid nmea: TotalOnGroundUnit",
			raw:  "$IIVLW,10.1,N,3.2,N,1,x,0.1,N*54",
			err:  "nmea: IIVLW invalid total cumulative ground distance unit: x",
		},
		{
			name: "invalid nmea: SinceResetOnGroundUnit",
			raw:  "$IIVLW,10.1,N,3.2,N,1,N,0.1,x*54",
			err:  "nmea: IIVLW invalid ground distance since reset unit: x",
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
				hdt := m.(VLW)
				hdt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hdt)
			}
		})
	}
}
