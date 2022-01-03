package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXTE(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  XTE
	}{
		{
			name: "good sentence",
			raw:  "$GPXTE,V,V,10.1,L,N*6E",
			msg: XTE{
				StatusGeneralWarning:     "V",
				StatusLockWarning:        "V",
				CrossTrackErrorMagnitude: 10.1,
				DirectionToSteer:         "L",
				CrossTrackUnits:          "N",
				FFAMode:                  "",
			},
		},
		{
			name: "good sentence with FAAMode",
			raw:  "$GPXTE,V,V,,,N,S*43",
			msg: XTE{
				StatusGeneralWarning:     "V",
				StatusLockWarning:        "V",
				CrossTrackErrorMagnitude: 0,
				DirectionToSteer:         "",
				CrossTrackUnits:          "N",
				FFAMode:                  "S",
			},
		},
		{
			name: "invalid nmea: StatusGeneralWarning",
			raw:  "$GPXTE,x,V,,,N,S*6d",
			err:  "nmea: GPXTE invalid general warning: x",
		},
		{
			name: "invalid nmea: StatusLockWarning",
			raw:  "$GPXTE,V,x,,,N,S*6d",
			err:  "nmea: GPXTE invalid lock warning: x",
		},
		{
			name: "invalid nmea: DirectionToSteer",
			raw:  "$GPXTE,V,V,,x,N,S*3b",
			err:  "nmea: GPXTE invalid direction to steer: x",
		},
		{
			name: "invalid nmea: CrossTrackUnits",
			raw:  "$GPXTE,V,V,,,x,S*75",
			err:  "nmea: GPXTE invalid cross track units: x",
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
				xte := m.(XTE)
				xte.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, xte)
			}
		})
	}
}
