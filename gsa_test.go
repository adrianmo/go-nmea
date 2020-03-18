package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var gsatests = []struct {
	name string
	raw  string
	err  string
	msg  GSA
}{
	{
		name: "good sentence",
		raw:  "$GPGSA,A,3,22,19,18,27,14,03,,,,,,,3.1,2.0,2.4*36",
		msg: GSA{
			Mode:    "A",
			FixType: "3",
			SV:      []string{"22", "19", "18", "27", "14", "03"},
			PDOP:    3.1,
			HDOP:    2,
			VDOP:    2.4,
			GSID:    "1",
		},
	},
	{
		name: "good sentence of GNGSA",
		raw:  "$GNGSA,A,3,01,03,07,17,22,28,30,,,,,,1.8,0.9,1.5,1*3F",
		msg: GSA{
			Mode:    "A",
			FixType: "3",
			SV:      []string{"01", "03", "07", "17", "22", "28", "30"},
			PDOP:    1.8,
			HDOP:    0.9,
			VDOP:    1.5,
			GSID:    "1",
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

func TestGSA(t *testing.T) {
	for _, tt := range gsatests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				gsa := m.(GSA)
				gsa.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, gsa)
			}
		})
	}
}
