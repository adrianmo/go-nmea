package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var wpltests = []struct {
	name string
	raw  string
	err  string
	msg  WPL
}{
	{
		name: "good sentence",
		raw:  "$IIWPL,5503.4530,N,01037.2742,E,411*6F",
		msg: WPL{
			Latitude:  MustParseLatLong("5503.4530 N"),
			Longitude: MustParseLatLong("01037.2742 E"),
			Ident:     "411",
		},
	},
	{
		name: "bad latitude",
		raw:  "$IIWPL,A,N,01037.2742,E,411*01",
		err:  "nmea: IIWPL invalid latitude: cannot parse [A N], unknown format",
	},
	{
		name: "bad longitude",
		raw:  "$IIWPL,5503.4530,N,A,E,411*36",
		err:  "nmea: IIWPL invalid longitude: cannot parse [A E], unknown format",
	},
	{
		name: "good sentence",
		raw:  "$IIWPL,3356.4650,S,15124.5567,E,411*73",
		msg: WPL{
			Latitude:  MustParseLatLong("3356.4650 S"),
			Longitude: MustParseLatLong("15124.5567 E"),
			Ident:     "411",
		},
	},
	{
		name: "bad latitude",
		raw:  "$IIWPL,A,S,15124.5567,E,411*18",
		err:  "nmea: IIWPL invalid latitude: cannot parse [A S], unknown format",
	},
	{
		name: "bad longitude",
		raw:  "$IIWPL,3356.4650,S,A,E,411*2E",
		err:  "nmea: IIWPL invalid longitude: cannot parse [A E], unknown format",
	},
}

func TestWPL(t *testing.T) {
	for _, tt := range wpltests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				wpl := m.(WPL)
				wpl.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, wpl)
			}
		})
	}
}
