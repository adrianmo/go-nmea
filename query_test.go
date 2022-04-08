package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuery(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  Query
	}{
		{
			name: "good sentence",
			raw:  "$CCGPQ,GGA*2B",
			msg: Query{
				BaseSentence:        BaseSentence{},
				DestinationTalkerID: "GP",
				RequestedSentence:   "GGA",
			},
		},
		{
			name: "invalid nmea: RequestedSentence",
			raw:  "$CCGPQ*46",
			err:  "nmea: CCQ invalid requested sentence: index out of range",
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
				rmb := m.(Query)
				rmb.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, rmb)
			}
		})
	}
}
