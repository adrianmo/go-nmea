package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDBK(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  DBK
	}{
		{
			name: "good sentence",
			raw:  "$SDDBK,12.3,f,3.7,M,2.0,F*2F",
			msg: DBK{
				DepthFeet:        12.3,
				DepthFeetUnit:    DistanceUnitFeet,
				DepthMeters:      3.7,
				DepthMetersUnit:  DistanceUnitMetre,
				DepthFathoms:     2,
				DepthFathomsUnit: DistanceUnitFathom,
			},
		},
		{
			name: "invalid nmea: DepthFeetUnit",
			raw:  "$SDDBK,12.3,x,3.7,M,2.0,F*31",
			err:  "nmea: SDDBK invalid depth feet unit: x",
		},
		{
			name: "invalid nmea: DepthMeterUnit",
			raw:  "$SDDBK,12.3,f,3.7,x,2.0,F*1A",
			err:  "nmea: SDDBK invalid depth meters unit: x",
		},
		{
			name: "invalid nmea: DepthFathomUnit",
			raw:  "$SDDBK,12.3,f,3.7,M,2.0,x*11",
			err:  "nmea: SDDBK invalid depth fathom unit: x",
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
				dbk := m.(DBK)
				dbk.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, dbk)
			}
		})
	}
}
