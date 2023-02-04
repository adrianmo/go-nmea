package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHBT(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  HBT
	}{
		{
			name: "good sentence",
			raw:  "$HCHBT,1.5,A,1*23",
			msg: HBT{
				Interval:        1.5,
				OperationStatus: StatusValid,
				MessageID:       1,
			},
		},
		{
			name: "invalid interval and status",
			raw:  "$HCHBT,,V,1*1E",
			msg: HBT{
				Interval:        0,
				OperationStatus: StatusInvalid,
				MessageID:       1,
			},
		},
		{
			name: "invalid interval",
			raw:  "$HCHBT,x.5,A,1*6A",
			err:  "nmea: HCHBT invalid interval: x.5",
		},
		{
			name: "invalid operation status",
			raw:  "$HCHBT,1.5,X,1*3A",
			err:  "nmea: HCHBT invalid operation status: X",
		},
		{
			name: "invalid sequence identification",
			raw:  "$HCHBT,1.5,A,x*6A",
			err:  "nmea: HCHBT invalid message ID: x",
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
				hbt := m.(HBT)
				hbt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hbt)
			}
		})
	}
}
