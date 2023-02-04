package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestACN(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  ACN
	}{
		{
			name: "good sentence",
			raw:  "$RAACN,220516,TCK,002,1,A,C*00",
			msg: ACN{
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
				State:                    "C",
			},
		},
		{
			name: "invalid nmea: Time",
			raw:  "$RAACN,2x0516,TCK,002,1,A,C*4a",
			err:  "nmea: RAACN invalid time: 2x0516",
		},
		{
			name: "invalid nmea: AlertIdentifier",
			raw:  "$RAACN,220516,TCK,x02,1,A,C*48",
			err:  "nmea: RAACN invalid alert identifier: x02",
		},
		{
			name: "invalid nmea: AlertInstance",
			raw:  "$RAACN,220516,TCK,002,x,A,C*49",
			err:  "nmea: RAACN invalid alert instance: x",
		},
		{
			name: "invalid nmea: Command",
			raw:  "$RAACN,220516,TCK,002,1,x,C*39",
			err:  "nmea: RAACN invalid alert command: x",
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
				acn := m.(ACN)
				acn.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, acn)
			}
		})
	}
}
