package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pkldstests = []struct {
	name string
	raw  string
	err  string
	msg  PKLDS
}{
	{
		name: "good sentence West",
		raw:  "$PKLDS,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W00,100,2000,15,00,*60",
		msg: PKLDS{
			Time:      Time{true, 22, 05, 16, 0},
			Validity:  "A",
			Latitude:  MustParseGPS("5133.82 N"),
			Longitude: MustParseGPS("00042.24 W"),
			Speed:     173.8,
			Course:    231.8,
			Date:      Date{true, 13, 06, 94},
			Variation: -4.2,
                        SentanceVersion: "00",
			Fleet:     "100",
			UnitID:    "2000",
                        Status:    "15",
                        Extension: "00",
		},
	},
        {
                name: "good sentence East",
                raw:  "$PKLDS,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,E00,100,2000,15,00,*72",
                msg: PKLDS{
                        Time:      Time{true, 22, 05, 16, 0},
                        Validity:  "A",
                        Latitude:  MustParseGPS("5133.82 N"),
                        Longitude: MustParseGPS("00042.24 W"),
                        Speed:     173.8,
                        Course:    231.8,
                        Date:      Date{true, 13, 06, 94},
                        Variation: 4.2,
                        SentanceVersion: "00",
                        Fleet:     "100",
                        UnitID:    "2000",
                        Status:    "15",
                        Extension: "00",
                },
        },

	{
		name: "bad sentence",
		raw:  "$PKLDS,220516,D,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W00,100,2000,15,00,*65",
		err:  "nmea: PKLDS invalid validity: D",
	},
}

func TestPKLDS(t *testing.T) {
	for _, tt := range pkldstests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				pklds := m.(PKLDS)
				pklds.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, pklds)
			}
		})
	}
}
