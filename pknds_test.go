package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pkndstests = []struct {
	name string
	raw  string
	err  string
	msg  PKNDS
}{
	{
		name: "good sentence West",
		raw:  "$PKNDS,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W00,U00001,207,00,*28",
		msg: PKNDS{
			Time:            Time{true, 22, 05, 16, 0},
			Validity:        "A",
			Latitude:        MustParseGPS("5133.82 N"),
			Longitude:       MustParseGPS("00042.24 W"),
			Speed:           173.8,
			Course:          231.8,
			Date:            Date{true, 13, 06, 94},
			Variation:       -4.2,
			SentanceVersion: "00",
			UnitID:          "U00001",
			Status:          "207",
			Extension:       "00",
		},
	},
	{
		name: "good sentence East",
		raw:  "$PKNDS,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,E00,U00001,207,00,*3A",
		msg: PKNDS{
			Time:            Time{true, 22, 05, 16, 0},
			Validity:        "A",
			Latitude:        MustParseGPS("5133.82 N"),
			Longitude:       MustParseGPS("00042.24 W"),
			Speed:           173.8,
			Course:          231.8,
			Date:            Date{true, 13, 06, 94},
			Variation:       4.2,
			SentanceVersion: "00",
			UnitID:          "U00001",
			Status:          "207",
			Extension:       "00",
		},
	},

	{
		name: "bad sentence",
		raw:  "$PKNDS,220516,D,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,E00,U00001,207,00,*3F",
		err:  "nmea: PKNDS invalid validity: D",
	},
}

func TestPKNDS(t *testing.T) {
	for _, tt := range pkndstests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				pknds := m.(PKNDS)
				pknds.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, pknds)
			}
		})
	}
}
