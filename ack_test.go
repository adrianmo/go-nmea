package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestACK(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  ACK
	}{
		{
			name: "good sentence",
			raw:  "$VRACK,001*50",
			msg: ACK{
				AlertIdentifier: 1,
			},
		},
		{
			name: "invalid nmea: AlertIdentifier",
			raw:  "$VRACK,x*19",
			err:  "nmea: VRACK invalid alert identifier: x",
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
				ack := m.(ACK)
				ack.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, ack)
			}
		})
	}
}
