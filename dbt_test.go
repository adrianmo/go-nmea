package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dbttests = []struct {
	name string
	raw  string
	err  string
	msg  DBT
}{
	{
		name: "good sentence",
		raw:  "$IIDBT,032.93,f,010.04,M,005.42,F*2C",
		msg: DBT{
			DepthFeet:    MustParseDecimal("32.93"),
			DepthMeters:  MustParseDecimal("10.04"),
			DepthFathoms: MustParseDecimal("5.42"),
		},
	},
	{
		name: "bad validity",
		raw:  "$IIDBT,032.93,f,010.04,M,005.42,F*22",
		err:  "nmea: sentence checksum mismatch [2C != 22]",
	},
}

func TestDBT(t *testing.T) {
	for _, tt := range dbttests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				dbt := m.(DBT)
				dbt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, dbt)
			}
		})
	}
}
