package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEVE(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  EVE
	}{
		{
			name: "good sentence",
			raw:  "$FREVE,000001,DZ00513,Fire Alarm On: TEST DZ201 Name*0A",
			msg: EVE{
				Time: Time{
					Valid:       true,
					Hour:        0,
					Minute:      0,
					Second:      1,
					Millisecond: 0,
				},
				TagCode: "DZ00513",
				Message: "Fire Alarm On: TEST DZ201 Name",
			},
		},
		{
			name: "invalid nmea: Time",
			raw:  "$FREVE,0x0001,DZ00513,Fire Alarm On: TEST DZ201 Name*42",
			err:  "nmea: FREVE invalid time: 0x0001",
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
				hdt := m.(EVE)
				hdt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hdt)
			}
		})
	}
}
