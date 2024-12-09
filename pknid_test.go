package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pknidtests = []struct {
	name string
	raw  string
	err  string
	msg  PKNID
}{
	{
		name: "typical sentance",
		raw:  "$PKNID,00,U00001,015,00,*24",
		msg: PKNID{
			SentanceVersion: "00",
			UnitID:          "U00001",
			Status:          "015",
			Extension:       "00",
		},
	},
}

func TestPKNID(t *testing.T) {
	for _, tt := range pknidtests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				pknid := m.(PKNID)
				pknid.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, pknid)
			}
		})
	}
}
