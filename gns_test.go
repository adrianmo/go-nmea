package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var gnstests = []struct {
	name string
	raw  string
	err  string
	msg  GNS
}{
	{
		name: "good sentence A",
		raw:  "$GNGNS,014035.00,4332.69262,S,17235.48549,E,RR,13,0.9,25.63,11.24,,*70",
		msg: GNS{
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
			NavStatus:  "",
		},
	},
	{
		name: "good sentence B",
		raw:  "$GNGNS,094821.0,4849.931307,N,00216.053323,E,AA,14,0.6,161.5,48.0,,*6D",
		msg: GNS{
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
			NavStatus:  "",
		},
	},
	{
		name: "good sentence C",
		raw:  "$GNGNS,094821.0,4849.931307,N,00216.053323,E,AAN,14,0.6,161.5,48.0,,*23",
		msg: GNS{
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
			NavStatus:  "",
		},
	},
	{
		name: "good sentence D with nav status",
		raw:  "$GPGNS,224749.00,3333.4268304,N,11153.3538273,W,D,19,0.6,406.110,-26.294,6.0,0138,S*6A",
		msg: GNS{
			Time:       Time{Valid: true, Hour: 22, Minute: 47, Second: 49, Millisecond: 0},
			Latitude:   33.55711384000001,
			Longitude:  -111.88923045499999,
			Mode:       []string{"D"},
			SVs:        19,
			HDOP:       0.6,
			Altitude:   406.11,
			Separation: -26.294,
			Age:        6,
			Station:    138,
			NavStatus:  "S",
		},
	},
	{
		name: "bad sentence",
		raw:  "$GNGNS,094821.0,4849.931307,N,00216.053323,E,AAX,14,0.6,161.5,48.0,,*35",
		err:  "nmea: GNGNS invalid mode: AAX",
	},
}

func TestGNS(t *testing.T) {
	for _, tt := range gnstests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				gns := m.(GNS)
				gns.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, gns)
			}
		})
	}
}
