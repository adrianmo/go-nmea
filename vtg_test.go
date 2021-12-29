package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var vtgtests = []struct {
	name string
	raw  string
	err  string
	msg  VTG
}{
	{
		name: "good sentence",
		raw:  "$GPVTG,45.5,T,67.5,M,30.45,N,56.40,K*4B",
		msg: VTG{
			TrueTrack:        45.5,
			MagneticTrack:    67.5,
			GroundSpeedKnots: 30.45,
			GroundSpeedKPH:   56.4,
			FFAMode:          "",
		},
	},
	{
		name: "good sentence with FAA mode",
		raw:  "$GPVTG,220.86,T,,M,2.550,N,4.724,K,A*34",
		msg: VTG{
			TrueTrack:        220.86,
			MagneticTrack:    0,
			GroundSpeedKnots: 2.55,
			GroundSpeedKPH:   4.724,
			FFAMode:          "A",
		},
	},
	{
		name: "bad true track",
		raw:  "$GPVTG,T,45.5,67.5,M,30.45,N,56.40,K*4B",
		err:  "nmea: GPVTG invalid true track: T",
	},
}

func TestVTG(t *testing.T) {
	for _, tt := range vtgtests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				vtg := m.(VTG)
				vtg.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, vtg)
			}
		})
	}
}
