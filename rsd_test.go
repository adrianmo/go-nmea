package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRSD(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  RSD
	}{
		{
			name: "good sentence",
			raw:  "$RARSD,0.00,,2.50,005.0,0.00,,4.50,355.0,,,3.0,N,H*51",
			msg: RSD{
				BaseSentence:           BaseSentence{},
				Origin1Range:           0,
				Origin1Bearing:         0,
				VariableRangeMarker1:   2.5,
				BearingLine1:           5,
				Origin2Range:           0,
				Origin2Bearing:         0,
				VariableRangeMarker2:   4.5,
				BearingLine2:           355,
				CursorRangeFromOwnShip: 0,
				CursorBearingDegrees:   0,
				RangeScale:             3,
				RangeUnit:              "N",
				DisplayRotation:        "H",
			},
		},
		{
			name: "good sentence 2",
			raw:  "$RARSD,,,,,,,,,0.808,326.9,0.750,N,N*58",
			msg: RSD{
				BaseSentence:           BaseSentence{},
				Origin1Range:           0,
				Origin1Bearing:         0,
				VariableRangeMarker1:   0,
				BearingLine1:           0,
				Origin2Range:           0,
				Origin2Bearing:         0,
				VariableRangeMarker2:   0,
				BearingLine2:           0,
				CursorRangeFromOwnShip: 0.808,
				CursorBearingDegrees:   326.9,
				RangeScale:             0.75,
				RangeUnit:              "N",
				DisplayRotation:        "N",
			},
		},
		{
			name: "invalid nmea: Origin1Range",
			raw:  "$RARSD,x,,2.50,005.0,0.00,,4.50,355.0,,,3.0,N,H*37",
			err:  "nmea: RARSD invalid origin 1 range: x",
		},
		{
			name: "invalid nmea: Origin1Bearing",
			raw:  "$RARSD,,x,2.50,005.0,0.00,,4.50,355.0,,,3.0,N,H*37",
			err:  "nmea: RARSD invalid origin 1 bearing: x",
		},
		{
			name: "invalid nmea: VariableRangeMarker1",
			raw:  "$RARSD,,,x2.50,005.0,0.00,,4.50,355.0,,,3.0,N,H*37",
			err:  "nmea: RARSD invalid variable range marker 1: x2.50",
		},
		{
			name: "invalid nmea: BearingLine1",
			raw:  "$RARSD,,,2.50,x005.0,0.00,,4.50,355.0,,,3.0,N,H*37",
			err:  "nmea: RARSD invalid bearing line 1: x005.0",
		},
		{
			name: "invalid nmea: Origin2Range",
			raw:  "$RARSD,,,2.50,005.0,x0.00,,4.50,355.0,,,3.0,N,H*37",
			err:  "nmea: RARSD invalid origin 2 range: x0.00",
		},
		{
			name: "invalid nmea: Origin2Bearing",
			raw:  "$RARSD,,,2.50,005.0,0.00,x,4.50,355.0,,,3.0,N,H*37",
			err:  "nmea: RARSD invalid origin 2 bearing: x",
		},
		{
			name: "invalid nmea: VariableRangeMarker2",
			raw:  "$RARSD,,,2.50,005.0,0.00,,x4.50,355.0,,,3.0,N,H*37",
			err:  "nmea: RARSD invalid variable range marker 2: x4.50",
		},
		{
			name: "invalid nmea: BearingLine2",
			raw:  "$RARSD,,,2.50,005.0,0.00,,4.50,x355.0,,,3.0,N,H*37",
			err:  "nmea: RARSD invalid bearing line 2: x355.0",
		},
		{
			name: "invalid nmea: CursorRangeFromOwnShip",
			raw:  "$RARSD,,,2.50,005.0,0.00,,4.50,355.0,x,,3.0,N,H*37",
			err:  "nmea: RARSD invalid cursor range from own ship: x",
		},
		{
			name: "invalid nmea: CursorBearingDegrees",
			raw:  "$RARSD,,,2.50,005.0,0.00,,4.50,355.0,,x,3.0,N,H*37",
			err:  "nmea: RARSD invalid cursor bearing: x",
		},
		{
			name: "invalid nmea: RangeUnit",
			raw:  "$RARSD,,,2.50,005.0,0.00,,4.50,355.0,,,3.0,X,H*59",
			err:  "nmea: RARSD invalid range units: X",
		},
		{
			name: "invalid nmea: RangeUnit",
			raw:  "$RARSD,,,2.50,005.0,0.00,,4.50,355.0,,,3.0,X,H*59",
			err:  "nmea: RARSD invalid range units: X",
		},
		{
			name: "invalid nmea: DisplayRotation",
			raw:  "$RARSD,,,2.50,005.0,0.00,,4.50,355.0,,,3.0,N,X*5f",
			err:  "nmea: RARSD invalid display rotation: X",
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
				mm := m.(RSD)
				mm.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, mm)
			}
		})
	}
}
