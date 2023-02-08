package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPMTK001(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  PMTK001
	}{
		{
			name: "good: Packet Type: 001 PMTK_ACK",
			raw:  "$PMTK001,604,3*32",
			msg: PMTK001{
				Cmd:  604,
				Flag: 3,
			},
		},
		{
			name: "missing flag",
			raw:  "$PMTK001,604*2d",
			err:  "nmea: PMTK001 invalid flag: index out of range",
		},
		{
			name: "missing cmd",
			raw:  "$PMTK001*33",
			err:  "nmea: PMTK001 invalid command: index out of range",
		},
	}
	p := SentenceParser{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := p.Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				mtk := m.(PMTK001) // is used by non-global SentenceParser instance
				mtk.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, mtk)
			}
		})
	}
}

func TestDefaultParserUsesDeprecatedMTK(t *testing.T) {
	m, err := Parse("$PMTK001,604,3*32")
	assert.NoError(t, err)
	assert.IsType(t, MTK{}, m)
}
