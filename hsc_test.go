package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHSC(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  HSC
	}{
		{
			name: "good sentence",
			raw:  "$FTHSC,40.12,T,39.11,M*5E",
			msg: HSC{
				TrueHeading:         40.12,
				TrueHeadingType:     HeadingTrue,
				MagneticHeading:     39.11,
				MagneticHeadingType: HeadingMagnetic,
			},
		},
		{
			name: "invalid nmea: TrueHeading",
			raw:  "$FTHSC,40.1x,T,39.11,M*14",
			err:  "nmea: FTHSC invalid true heading: 40.1x",
		},
		{
			name: "invalid nmea: TrueHeadingType",
			raw:  "$FTHSC,40.12,x,39.11,M*72",
			err:  "nmea: FTHSC invalid true heading type: x",
		},
		{
			name: "invalid nmea: MagneticHeading",
			raw:  "$FTHSC,40.12,T,x,M*02",
			err:  "nmea: FTHSC invalid magnetic heading: x",
		},
		{
			name: "invalid nmea: MagneticHeadingType",
			raw:  "$FTHSC,40.12,T,39.11,x*6b",
			err:  "nmea: FTHSC invalid magnetic heading type: x",
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
				hsc := m.(HSC)
				hsc.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hsc)
			}
		})
	}
}
