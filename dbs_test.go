package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBS(t *testing.T) {
	var dbstests = []struct {
		name string
		raw  string
		err  string
		msg  DBS
	}{
		{
			name: "good sentence",
			raw:  "$23DBS,01.9,f,0.58,M,00.3,F*21",
			msg: DBS{
				DepthFeet:       1.9,
				DepthFeetUnit:   DistanceUnitFeet,
				DepthMeters:     0.58,
				DepthMeterUnit:  DistanceUnitMetre,
				DepthFathoms:    0.3,
				DepthFathomUnit: DistanceUnitFathom,
			},
		},
		{
			name: "good sentence 2",
			raw:  "$SDDBS,,,0187.5,M,,*1A", // Simrad ITI Trawl System
			msg: DBS{
				DepthFeet:       0,
				DepthFeetUnit:   "",
				DepthMeters:     187.5,
				DepthMeterUnit:  DistanceUnitMetre,
				DepthFathoms:    0,
				DepthFathomUnit: "",
			},
		},
		{
			name: "bad validity",
			raw:  "$23DBS,01.9,f,0.58,M,00.3,F*25",
			err:  "nmea: sentence checksum mismatch [21 != 25]",
		},
	}

	for _, tt := range dbstests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				dbs := m.(DBS)
				dbs.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, dbs)
			}
		})
	}
}
