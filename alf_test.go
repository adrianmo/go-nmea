package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestALF(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  ALF
	}{
		{
			name: "good sentence",
			raw:  "$VDALF,1,0,1,220516,B,A,S,SAL,001,1,2,0,My alarm*2c",
			msg: ALF{
				NumFragments:   1,
				FragmentNumber: 0,
				MessageID:      1,
				Time: Time{
					Valid:       true,
					Hour:        22,
					Minute:      05,
					Second:      16,
					Millisecond: 0,
				},
				Category:                 "B",
				Priority:                 "A",
				State:                    "S",
				ManufacturerMnemonicCode: "SAL",
				AlertIdentifier:          1,
				AlertInstance:            1,
				RevisionCounter:          2,
				EscalationCounter:        0,
				Text:                     "My alarm",
			},
		},
		{
			name: "invalid nmea: NumFragments",
			raw:  "$VDALF,x,0,1,220516,B,A,S,SAL,001,1,2,0,My alarm*65",
			err:  "nmea: VDALF invalid number of fragments: x",
		},
		{
			name: "invalid nmea: FragmentNumber",
			raw:  "$VDALF,1,x,1,220516,B,A,S,SAL,001,1,2,0,My alarm*64",
			err:  "nmea: VDALF invalid fragment number: x",
		},
		{
			name: "invalid nmea: MessageID",
			raw:  "$VDALF,1,0,x,220516,B,A,S,SAL,001,1,2,0,My alarm*65",
			err:  "nmea: VDALF invalid message ID: x",
		},
		{
			name: "invalid nmea: Time",
			raw:  "$VDALF,1,0,1,2x0516,B,A,S,SAL,001,1,2,0,My alarm*66",
			err:  "nmea: VDALF invalid time: 2x0516",
		},
		{
			name: "invalid nmea: Category",
			raw:  "$VDALF,1,0,1,220516,x,A,S,SAL,001,1,2,0,My alarm*16",
			err:  "nmea: VDALF invalid alarm category: x",
		},
		{
			name: "invalid nmea: Priority",
			raw:  "$VDALF,1,0,1,220516,B,x,S,SAL,001,1,2,0,My alarm*15",
			err:  "nmea: VDALF invalid alarm priority: x",
		},
		{
			name: "invalid nmea: State",
			raw:  "$VDALF,1,0,1,220516,B,A,x,SAL,001,1,2,0,My alarm*07",
			err:  "nmea: VDALF invalid alarm state: x",
		},
		{
			name: "invalid nmea: AlertIdentifier",
			raw:  "$VDALF,1,0,1,220516,B,A,S,SAL,x01,1,2,0,My alarm*64",
			err:  "nmea: VDALF invalid alert identifier: x01",
		},
		{
			name: "invalid nmea: AlertInstance",
			raw:  "$VDALF,1,0,1,220516,B,A,S,SAL,001,x,2,0,My alarm*65",
			err:  "nmea: VDALF invalid alert instance: x",
		},
		{
			name: "invalid nmea: RevisionCounter",
			raw:  "$VDALF,1,0,1,220516,B,A,S,SAL,001,1,x,0,My alarm*66",
			err:  "nmea: VDALF invalid revision counter: x",
		},
		{
			name: "invalid nmea: EscalationCounter",
			raw:  "$VDALF,1,0,1,220516,B,A,S,SAL,001,1,2,x,My alarm*64",
			err:  "nmea: VDALF invalid escalation counter: x",
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
				alf := m.(ALF)
				alf.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, alf)
			}
		})
	}
}
