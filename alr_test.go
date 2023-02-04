package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestALR(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  ALR
	}{
		{
			name: "good sentence",
			raw:  "$RAALR,220516,001,A,A,Bilge pump alarm1*4c",
			msg: ALR{
				Time: Time{
					Valid:       true,
					Hour:        22,
					Minute:      05,
					Second:      16,
					Millisecond: 0,
				},
				AlarmIdentifier: 1,
				Condition:       StatusValid,
				State:           StatusValid,
				Description:     "Bilge pump alarm1",
			},
		},
		{
			name: "nmea: Description empty",
			raw:  "$RAALR,220516,001,A,A,*53",
			msg: ALR{
				Time: Time{
					Valid:       true,
					Hour:        22,
					Minute:      05,
					Second:      16,
					Millisecond: 0,
				},
				AlarmIdentifier: 1,
				Condition:       StatusValid,
				State:           StatusValid,
				Description:     "",
			},
		},
		{
			name: "invalid nmea: Time",
			raw:  "$RAALR,2x0516,001,A,A,Bilge pump alarm1*06",
			err:  "nmea: RAALR invalid time: 2x0516",
		},
		{
			name: "invalid nmea: Condition",
			raw:  "$RAALR,220516,001,x,A,Bilge pump alarm1*75",
			err:  "nmea: RAALR invalid alarm condition: x",
		},
		{
			name: "invalid nmea: State",
			raw:  "$RAALR,220516,001,A,x,Bilge pump alarm1*75",
			err:  "nmea: RAALR invalid alarm state: x",
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
				alr := m.(ALR)
				alr.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, alr)
			}
		})
	}
}
