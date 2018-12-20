package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var hdttests = []struct {
	name string
	raw  string
	err  string
	msg  HDT
}{
	{
		name: "good sentence",
		raw:  "$GPHDT,123.456,T*32",
		msg: HDT{
			Heading: 123.456,
			True:    true,
		},
	},
	{
		name: "invalid True",
		raw:  "$GPHDT,123.456,X*3E",
		err:  "nmea: GPHDT invalid true: X",
	},
	{
		name: "invalid Heading",
		raw:  "$GPHDT,XXX,T*43",
		err:  "nmea: GPHDT invalid heading: XXX",
	},
}

func TestHDT(t *testing.T) {
	for _, tt := range hdttests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				hdt := m.(HDT)
				hdt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hdt)
			}
		})
	}
}
