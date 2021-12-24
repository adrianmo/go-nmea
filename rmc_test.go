package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var rmctests = []struct {
	name string
	raw  string
	err  string
	msg  RMC
}{
	{
		name: "good sentence A",
		raw:  "$GNRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*6E",
		msg: RMC{
			Time:      Time{true, 22, 05, 16, 0},
			Validity:  "A",
			Latitude:  MustParseGPS("5133.82 N"),
			Longitude: MustParseGPS("00042.24 W"),
			Speed:     173.8,
			Course:    231.8,
			Date:      Date{true, 13, 06, 94},
			Variation: -4.2,
			FFAMode:   "",
			NavStatus: "",
		},
	},
	{
		name: "good sentence B",
		raw:  "$GNRMC,142754.0,A,4302.539570,N,07920.379823,W,0.0,,070617,0.0,E,A*21",
		msg: RMC{
			Time:      Time{true, 14, 27, 54, 0},
			Validity:  "A",
			Latitude:  MustParseGPS("4302.539570 N"),
			Longitude: MustParseGPS("07920.379823 W"),
			Speed:     0,
			Course:    0,
			Date:      Date{true, 7, 6, 17},
			Variation: 0,
			FFAMode:   FAAModeAutonomous,
			NavStatus: "",
		},
	},
	{
		name: "good sentence C",
		raw:  "$GNRMC,100538.00,A,5546.27711,N,03736.91144,E,0.061,,260318,,,A*60",
		msg: RMC{
			Time:      Time{true, 10, 5, 38, 0},
			Validity:  "A",
			Latitude:  MustParseGPS("5546.27711 N"),
			Longitude: MustParseGPS("03736.91144 E"),
			Speed:     0.061,
			Course:    0,
			Date:      Date{true, 26, 3, 18},
			Variation: 0,
			FFAMode:   FAAModeAutonomous,
			NavStatus: "",
		},
	},
	{
		name: "bad sentence",
		raw:  "$GNRMC,220516,D,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*6B",
		err:  "nmea: GNRMC invalid validity: D",
	},
	{
		name: "good sentence D",
		raw:  "$GPRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*70",
		msg: RMC{
			Time:      Time{true, 22, 5, 16, 0},
			Validity:  "A",
			Latitude:  MustParseGPS("5133.82 N"),
			Longitude: MustParseGPS("00042.24 W"),
			Speed:     173.8,
			Course:    231.8,
			Date:      Date{true, 13, 6, 94},
			Variation: -4.2,
			FFAMode:   "",
			NavStatus: "",
		},
	},
	{
		name: "good sentence E",
		raw:  "$GPRMC,142754.0,A,4302.539570,N,07920.379823,W,0.0,,070617,0.0,E,A*3F",
		msg: RMC{
			Time:      Time{true, 14, 27, 54, 0},
			Validity:  "A",
			Latitude:  MustParseGPS("4302.539570 N"),
			Longitude: MustParseGPS("07920.379823 W"),
			Speed:     0,
			Course:    0,
			Date:      Date{true, 7, 6, 17},
			Variation: 0,
			FFAMode:   FAAModeAutonomous,
			NavStatus: "",
		},
	},
	{
		name: "good sentence F with nav status",
		raw:  "$GNRMC,102014.00,A,5550.6082,N,03732.2488,E,000.00000,092.9,300518,,,A,V*3B",
		msg: RMC{
			Time:      Time{Valid: true, Hour: 10, Minute: 20, Second: 14, Millisecond: 0},
			Validity:  "A",
			Latitude:  55.843469999999996,
			Longitude: 37.537479999999995,
			Speed:     0,
			Course:    92.9,
			Date:      Date{Valid: true, DD: 30, MM: 5, YY: 18},
			Variation: 0,
			FFAMode:   FAAModeAutonomous,
			NavStatus: NavStatusDataValid,
		},
	},
	{
		name: "bad validity",
		raw:  "$GPRMC,220516,D,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*75",
		err:  "nmea: GPRMC invalid validity: D",
	},
}

func TestRMC(t *testing.T) {
	for _, tt := range rmctests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				rmc := m.(RMC)
				rmc.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, rmc)
			}
		})
	}
}
