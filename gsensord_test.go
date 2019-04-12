package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var gsensoredtests = []struct {
	name string
	raw  string
	err  string
	msg  GSensord
}{
	{
		name: "good sentence",
		raw:  "$GSENSORD,0.060,0.060,-0.180",
		msg: GSensord{
			X: 0.06,
			Y: 0.06,
			Z: -0.180,
		},
	},
}

func TestGSensord(t *testing.T) {
	for _, tt := range gsensoredtests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				gsensored := m.(GSensord)
				gsensored.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, gsensored)
			}
		})
	}
}
