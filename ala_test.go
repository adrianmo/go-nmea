package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestALA(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  ALA
	}{
		{
			name: "good sentence",
			raw:  "$FRALA,143955,FR,OT,00,901,N,V,Syst Fault : AutroSafe comm. OK*4F",
			msg: ALA{
				Time: Time{
					Valid:       true,
					Hour:        14,
					Minute:      39,
					Second:      55,
					Millisecond: 0,
				},
				SystemIndicator:    "FR",
				SubSystemIndicator: "OT",
				InstanceNumber:     0,
				Type:               901,
				Condition:          "N",
				AlarmAckState:      "V",
				Message:            "Syst Fault : AutroSafe comm. OK",
			},
		},
		{
			name: "invalid nmea: Time",
			raw:  "$FRALA,1x3955,FR,OT,00,901,N,V,Syst Fault : AutroSafe comm. OK*03",
			err:  "nmea: FRALA invalid time: 1x3955",
		},
		{
			name: "invalid nmea: InstanceNumber",
			raw:  "$FRALA,143955,FR,OT,x0,901,N,V,Syst Fault : AutroSafe comm. OK*07",
			err:  "nmea: FRALA invalid instance number: x0",
		},
		{
			name: "invalid nmea: Type",
			raw:  "$FRALA,143955,FR,OT,00,9x1,N,V,Syst Fault : AutroSafe comm. OK*07",
			err:  "nmea: FRALA invalid type: 9x1",
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
				ala := m.(ALA)
				ala.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, ala)
			}
		})
	}
}
