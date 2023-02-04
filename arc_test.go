package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestARC(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  ARC
	}{
		{
			name: "good sentence",
			raw:  "$RAARC,220516,TCK,002,1,A*73",
			msg: ARC{
				Time: Time{
					Valid:       true,
					Hour:        22,
					Minute:      05,
					Second:      16,
					Millisecond: 0,
				},
				ManufacturerMnemonicCode: "TCK",
				AlertIdentifier:          2,
				AlertInstance:            1,
				Command:                  AlertCommandAcknowledge,
			},
		},
		{
			name: "invalid nmea: Time",
			raw:  "$RAARC,2x0516,TCK,002,1,A*39",
			err:  "nmea: RAARC invalid time: 2x0516",
		},
		{
			name: "invalid nmea: AlertIdentifier",
			raw:  "$RAARC,220516,TCK,x02,1,A*3b",
			err:  "nmea: RAARC invalid alert identifier: x02",
		},
		{
			name: "invalid nmea: AlertInstance",
			raw:  "$RAARC,220516,TCK,002,x,A*3a",
			err:  "nmea: RAARC invalid alert instance: x",
		},
		{
			name: "invalid nmea: Command",
			raw:  "$RAARC,220516,TCK,002,1,x*4a",
			err:  "nmea: RAARC invalid refused alert command: x",
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
				arc := m.(ARC)
				arc.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, arc)
			}
		})
	}
}
