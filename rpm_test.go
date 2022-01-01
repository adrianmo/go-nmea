package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRPM(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  RPM
	}{
		{
			name: "good sentence",
			raw:  "$RCRPM,S,0,74.6,30.0,A*56",
			msg: RPM{
				Source:       SourceShaftRPM,
				EngineNumber: 0,
				SpeedRPM:     74.6,
				PitchPercent: 30,
				Status:       StatusValid,
			},
		},
		{
			name: "invalid nmea: Source",
			raw:  "$RCRPM,x,0,74.6,30.0,A*7D",
			err:  "nmea: RCRPM invalid source: x",
		},
		{
			name: "invalid nmea: Status",
			raw:  "$RCRPM,S,0,74.6,30.0,x*6F",
			err:  "nmea: RCRPM invalid status: x",
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
				hdt := m.(RPM)
				hdt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hdt)
			}
		})
	}
}
