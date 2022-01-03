package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTXT(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  TXT
	}{
		{
			name: "good sentence",
			raw:  "$GNTXT,01,01,02,u-blox AG - www.u-blox.com*4E",
			msg: TXT{
				TotalNumber: 1,
				Number:      1,
				ID:          2,
				Message:     "u-blox AG - www.u-blox.com",
			},
		},
		{
			name: "invalid TotalNumber",
			raw:  "$GNTXT,x,01,02,u-blox AG - www.u-blox.com*37",
			err:  "nmea: GNTXT invalid total number of sentences: x",
		},
		{
			name: "invalid Number",
			raw:  "$GNTXT,01,X,02,u-blox AG - www.u-blox.com*17",
			err:  "nmea: GNTXT invalid sentence number: X",
		},
		{
			name: "invalid ID",
			raw:  "$GNTXT,01,01,X,u-blox AG - www.u-blox.com*14",
			err:  "nmea: GNTXT invalid sentence identifier: X",
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
				txt := m.(TXT)
				txt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, txt)
			}
		})
	}
}
