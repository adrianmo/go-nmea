package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dpttests = []struct {
	name string
	raw  string
	err  string
	msg  DPT
}{
	{
		name: "good sentence",
		raw:  "$SDDPT,0.5,0.5,*7B",
		msg: DPT{
			Depth:  MustParseDecimal("0.5"),
			Offset: MustParseDecimal("0.5"),
		},
	},
	{
		name: "bad validity",
		raw:  "$SDDPT,0.5,0.5,*AA",
		err:  "nmea: sentence checksum mismatch [7B != AA]",
	},
}

func TestDPT(t *testing.T) {
	for _, tt := range dpttests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				dpt := m.(DPT)
				dpt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, dpt)
			}
		})
	}
}
