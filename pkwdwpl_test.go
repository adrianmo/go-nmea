package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pkwdkpltests = []struct {
	name string
	raw  string
	err  string
	msg  PKWDWPL
}{
	{
		name: "good sentence",
		raw:  "$PKWDWPL,150803,A,4237.14,N,07120.83,W,173.8,231.8,190316,1120,test,/'*39",
		msg: PKWDWPL{
			Time:         Time{true, 15, 8, 3, 0},
			Validity:     "A",
			Latitude:     MustParseGPS("4237.14 N"),
			Longitude:    MustParseGPS("07120.83 W"),
			Speed:        173.8,
			Course:       231.8,
			Date:         Date{true, 19, 3, 16},
			Altitude:     1120,
			WaypointName: "test",
			TableSymbol:  "/'",
		},
	},
}

func TestPKWDWPL(t *testing.T) {
	for _, tt := range pkwdkpltests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				pkwdwpl := m.(PKWDWPL)
				pkwdwpl.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, pkwdwpl)
			}
		})
	}
}
