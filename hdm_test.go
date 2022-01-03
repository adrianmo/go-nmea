package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHDM(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  HDM
	}{
		{
			name: "good sentence",
			raw:  "$HCHDM,093.8,M*2B",
			msg: HDM{
				Heading:       93.8,
				MagneticValid: true,
			},
		},
		{
			name: "invalid Magnetic",
			raw:  "$HCHDM,093.8,X*3E",
			err:  "nmea: HCHDM invalid magnetic: X",
		},
		{
			name: "invalid Heading",
			raw:  "$HCHDM,09X.X,M*20",
			err:  "nmea: HCHDM invalid heading: 09X.X",
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
				hdm := m.(HDM)
				hdm.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hdm)
			}
		})
	}
}
