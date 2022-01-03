package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDOR(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  DOR
	}{
		{
			name: "good sentence",
			raw:  "$FRDOR,E,233042,FD,FP,000,010,C,C,Door Closed : TEST FPA Name*4D",
			msg: DOR{
				Type: TypeSingleDoorDOR,
				Time: Time{
					Valid:       true,
					Hour:        23,
					Minute:      30,
					Second:      42,
					Millisecond: 0,
				},
				SystemIndicator:    "FD",
				DivisionIndicator1: "FP",
				DivisionIndicator2: 0,
				DoorNumberOrCount:  10,
				DoorStatus:         DoorStatusClosedDOR,
				SwitchSetting:      SwitchSettingSeaModeDOR,
				Message:            "Door Closed : TEST FPA Name",
			},
		},
		{
			name: "invalid nmea: Type",
			raw:  "$FRDOR,x,233042,FD,FP,000,010,C,C,Door Closed : TEST FPA Name*70",
			err:  "nmea: FRDOR invalid message type: x",
		},
		{
			name: "invalid nmea: Time",
			raw:  "$FRDOR,E,2x3042,FD,FP,000,010,C,C,Door Closed : TEST FPA Name*06",
			err:  "nmea: FRDOR invalid time: 2x3042",
		},
		{
			name: "invalid nmea: DoorStatus",
			raw:  "$FRDOR,E,233042,FD,FP,000,010,_,C,Door Closed : TEST FPA Name*51",
			err:  "nmea: FRDOR invalid door state: _",
		},
		{
			name: "invalid nmea: SwitchSetting",
			raw:  "$FRDOR,E,233042,FD,FP,000,010,C,_,Door Closed : TEST FPA Name*51",
			err:  "nmea: FRDOR invalid switch setting mode: _",
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
				dor := m.(DOR)
				dor.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, dor)
			}
		})
	}
}
