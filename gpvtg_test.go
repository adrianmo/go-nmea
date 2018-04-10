package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var gpvtgtests = []struct {
	name string
	raw  string
	err  string
	msg  GPVTG
}{
	{
		name: "good sentence",
		raw:  "$GPVTG,45.5,T,67.5,M,30.45,N,56.40,K*4B",
		msg: GPVTG{
			TrueTrack:        45.5,
			MagneticTrack:    67.5,
			GroundSpeedKnots: 30.45,
			GroundSpeedKPH:   56.4,
		},
	},
	{
		name: "bad true track",
		raw:  "$GPVTG,T,45.5,67.5,M,30.45,N,56.40,K*4B",
		err:  "nmea: GPVTG invalid true track: T",
	},
}

func TestGPVTG(t *testing.T) {
	for _, tt := range gpvtgtests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				gpvtg := m.(GPVTG)
				gpvtg.Sent = Sent{}
				assert.Equal(t, tt.msg, gpvtg)
			}
		})
	}
}
