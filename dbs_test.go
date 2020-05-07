package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			DepthFeet:    MustParseDecimal("1.9"),
			DepthMeters:  MustParseDecimal("0.58"),
			DepthFathoms: MustParseDecimal("0.3"),
		},
	},
	{
		name: "bad validity",
		raw:  "$23DBS,01.9,f,0.58,M,00.3,F*25",
		err:  "nmea: sentence checksum mismatch [21 != 25]",
	},
}

func TestDBS(t *testing.T) {
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
