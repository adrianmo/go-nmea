package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPMTK001(t *testing.T) {
	var mtktests = []struct {
		name string
		raw  string
		err  string
		msg  PMTK001
	}{
		{
			name: "good: Packet Type: 001 PMTK_ACK",
			raw:  "$PMTK001,604,3*" + Checksum("PMTK001,604,3"),
			msg: PMTK001{
				Cmd:  604,
				Flag: 3,
			},
		},
		{
			name: "missing flag",
			raw:  "$PMTK001,604*" + Checksum("PMTK001,604"),
			err:  "nmea: PMTK001 invalid flag: index out of range",
		},
		{
			name: "missing cmd",
			raw:  "$PMTK001*" + Checksum("PMTK001"),
			err:  "nmea: PMTK001 invalid command: index out of range",
		},
	}

	for _, tt := range mtktests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				mtk := m.(PMTK001)
				mtk.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, mtk)
			}
		})
	}
}
