package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFIR(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  FIR
	}{
		{
			name: "good sentence",
			raw:  "$FRFIR,E,103000,FD,PT,000,007,A,V,Fire Alarm : TEST PT7 Name TEST DZ2 Name*7A",
			msg: FIR{
				Type: TypeEventOrAlarmFIR,
				Time: Time{
					Valid:       true,
					Hour:        10,
					Minute:      30,
					Second:      0,
					Millisecond: 0,
				},
				SystemIndicator:           "FD",
				DivisionIndicator1:        "PT",
				DivisionIndicator2:        0,
				FireDetectorNumberOrCount: 7,
				Condition:                 ConditionActivationFIR,
				AlarmAckState:             AlarmStateNotAcknowledgedFIR,
				Message:                   "Fire Alarm : TEST PT7 Name TEST DZ2 Name",
			},
		},
		{
			name: "invalid nmea: Type",
			raw:  "$FRFIR,x,103000,FD,PT,000,007,A,V,Fire Alarm : TEST PT7 Name TEST DZ2 Name*47",
			err:  "nmea: FRFIR invalid message type: x",
		},
		{
			name: "invalid nmea: Time",
			raw:  "$FRFIR,E,1x3000,FD,PT,000,007,A,V,Fire Alarm : TEST PT7 Name TEST DZ2 Name*32",
			err:  "nmea: FRFIR invalid time: 1x3000",
		},
		{
			name: "invalid nmea: Condition",
			raw:  "$FRFIR,E,103000,FD,PT,000,007,_,V,Fire Alarm : TEST PT7 Name TEST DZ2 Name*64",
			err:  "nmea: FRFIR invalid condition: _",
		},
		{
			name: "invalid nmea: AlarmAckState",
			raw:  "$FRFIR,E,103000,FD,PT,000,007,A,_,Fire Alarm : TEST PT7 Name TEST DZ2 Name*73",
			err:  "nmea: FRFIR invalid alarm acknowledgement state: _",
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
				hdt := m.(FIR)
				hdt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hdt)
			}
		})
	}
}
