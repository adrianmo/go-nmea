package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMTA(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  MTA
	}{
		{
			name: "good sentence",
			raw:  "$IIMTA,13.3,C*04",
			msg: MTA{
				Temperature: 13.3,
				Unit:        TemperatureCelsius,
			},
		},
		{
			name: "invalid nmea: Temperature",
			raw:  "$IIMTA,x.x,C*35",
			err:  "nmea: IIMTA invalid temperature: x.x",
		},
		{
			name: "invalid nmea: Unit",
			raw:  "$IIMTA,13.3,F*01",
			err:  "nmea: IIMTA invalid temperature unit: F",
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
				mta := m.(MTA)
				mta.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, mta)
			}
		})
	}
}
