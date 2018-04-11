package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var gpgsatests = []struct {
	name string
	raw  string
	err  string
	msg  GPGSA
}{
	{
		name: "good sentence",
		raw:  "$GPGSA,A,3,22,19,18,27,14,03,,,,,,,3.1,2.0,2.4*36",
		msg: GPGSA{
			Mode:    "A",
			FixType: "3",
			SV:      []string{"22", "19", "18", "27", "14", "03"},
			PDOP:    3.1,
			HDOP:    2,
			VDOP:    2.4,
		},
	},
	{
		name: "bad mode",
		raw:  "$GPGSA,F,3,22,19,18,27,14,03,,,,,,,3.1,2.0,2.4*31",
		err:  "nmea: GPGSA invalid selection mode: F",
	},
	{
		name: "bad fix",
		raw:  "$GPGSA,A,6,22,19,18,27,14,03,,,,,,,3.1,2.0,2.4*33",
		err:  "nmea: GPGSA invalid fix type: 6",
	},
}

func TestGPGSA(t *testing.T) {
	for _, tt := range gpgsatests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				gpgsa := m.(GPGSA)
				gpgsa.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, gpgsa)
			}
		})
	}
}
