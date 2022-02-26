package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTLL(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  TLL
	}{
		{
			name: "good sentence",
			raw:  "$RATLL,,3647.422,N,01432.592,E,,,,*58",
			msg: TLL{
				BaseSentence:    BaseSentence{},
				TargetNumber:    0,
				TargetLatitude:  36.790366666666664,
				TargetLongitude: 14.543200000000002,
				TargetName:      "",
				TimeUTC: Time{
					Valid:       false,
					Hour:        0,
					Minute:      0,
					Second:      0,
					Millisecond: 0,
				},
				TargetStatus:    "",
				ReferenceTarget: "",
			},
		},
		{
			name: "good sentence 2",
			raw:  "$RATLL,1,3646.54266,N,00235.37778,W,test,020915,L,R*78",
			msg: TLL{
				BaseSentence:    BaseSentence{},
				TargetNumber:    1,
				TargetLatitude:  36.775711,
				TargetLongitude: -2.5896296666666667,
				TargetName:      "test",
				TimeUTC:         Time{Valid: true, Hour: 2, Minute: 9, Second: 15, Millisecond: 0},
				TargetStatus:    "L",
				ReferenceTarget: "R",
			},
		},
		{
			name: "invalid nmea: TargetNumber",
			raw:  "$RATLL,x,3647.422,N,01432.592,E,,,,*20",
			err:  "nmea: RATLL invalid target number: x",
		},
		{
			name: "invalid nmea: TargetLatitude",
			raw:  "$RATLL,1,x3647.422,N,01432.592,E,,,,*11",
			err:  "nmea: RATLL invalid latitude: cannot parse [x3647.422 N], unknown format",
		},
		{
			name: "invalid nmea: TargetLongitude",
			raw:  "$RATLL,1,3647.422,N,x01432.592,E,,,,*11",
			err:  "nmea: RATLL invalid longitude: cannot parse [x01432.592 E], unknown format",
		},
		{
			name: "invalid nmea: TimeUTC",
			raw:  "$RATLL,1,3646.54266,N,00235.37778,W,test,x020915,L,R*00",
			err:  "nmea: RATLL invalid UTC time: x020915",
		},
		{
			name: "invalid nmea: TargetStatus",
			raw:  "$RATLL,1,3646.54266,N,00235.37778,W,test,020915,xL,R*00",
			err:  "nmea: RATLL invalid target status: xL",
		},
		{
			name: "invalid nmea: ReferenceTarget",
			raw:  "$RATLL,1,3646.54266,N,00235.37778,W,test,020915,L,xR*00",
			err:  "nmea: RATLL invalid reference target: xR",
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
				mm := m.(TLL)
				mm.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, mm)
			}
		})
	}
}
