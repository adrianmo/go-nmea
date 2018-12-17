package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var gngnstests = []struct {
	name string
	raw  string
	err  string
	msg  GNGNS
}{
	{
		name: "good sentence A",
		raw:  "$GNGNS,014035.00,4332.69262,S,17235.48549,E,RR,13,0.9,25.63,11.24,,*70",
		msg: GNGNS{
			Time:       Time{true, 1, 40, 35, 0},
			Latitude:   MustParseGPS("4332.69262 S"),
			Longitude:  MustParseGPS("17235.48549 E"),
			Mode:       []string{"R", "R"},
			SVs:        13,
			HDOP:       0.9,
			Altitude:   25.63,
			Separation: 11.24,
			Age:        0,
			Station:    0,
		},
	},
	{
		name: "good sentence B",
		raw:  "$GNGNS,094821.0,4849.931307,N,00216.053323,E,AA,14,0.6,161.5,48.0,,*6D",
		msg: GNGNS{
			Time:       Time{true, 9, 48, 21, 0},
			Latitude:   MustParseGPS("4849.931307 N"),
			Longitude:  MustParseGPS("00216.053323 E"),
			Mode:       []string{"A", "A"},
			SVs:        14,
			HDOP:       0.6,
			Altitude:   161.5,
			Separation: 48.0,
			Age:        0,
			Station:    0,
		},
	},
	{
		name: "good sentence B",
		raw:  "$GNGNS,094821.0,4849.931307,N,00216.053323,E,AAN,14,0.6,161.5,48.0,,*23",
		msg: GNGNS{
			Time:       Time{true, 9, 48, 21, 0},
			Latitude:   MustParseGPS("4849.931307 N"),
			Longitude:  MustParseGPS("00216.053323 E"),
			Mode:       []string{"A", "A", "N"},
			SVs:        14,
			HDOP:       0.6,
			Altitude:   161.5,
			Separation: 48.0,
			Age:        0,
			Station:    0,
		},
	},
	{
		name: "bad sentence",
		raw:  "$GNGNS,094821.0,4849.931307,N,00216.053323,E,AAX,14,0.6,161.5,48.0,,*35",
		err:  "nmea: GNGNS invalid mode: AAX",
	},
}
var gngnsmodetests = []struct {
	name     string
	raw      string
	modes    []string
	expected bool
}{
	{
		name:     "matched mode string A",
		raw:      "$GNGNS,014035.00,4332.69262,S,17235.48549,E,AAA,13,0.9,25.63,11.24,,*31",
		modes:    []string{AutonomousGNGNS, DifferentialGNGNS},
		expected: true,
	},
	{
		name:     "matched mode string B",
		raw:      "$GNGNS,014035.00,4332.69262,S,17235.48549,E,AAD,13,0.9,25.63,11.24,,*34",
		modes:    []string{DifferentialGNGNS},
		expected: true,
	},
	{
		name:     "unmatched mode string",
		raw:      "$GNGNS,014035.00,4332.69262,S,17235.48549,E,AAD,13,0.9,25.63,11.24,,*34",
		modes:    []string{EstimatedGNGNS},
		expected: false,
	},
}

func TestGNGNS(t *testing.T) {
	for _, tt := range gngnstests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				gngns := m.(GNGNS)
				gngns.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, gngns)
			}
		})
	}
	for _, tt := range gngnsmodetests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			assert.NoError(t, err)
			gngns := m.(GNGNS)
			hasMode := gngns.IsMode(tt.modes...)
			assert.Equal(t, tt.expected, hasMode)

		})
	}
}
