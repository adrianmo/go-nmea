package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pklidtests = []struct {
	name string
	raw  string
	err  string
	msg  PKLID
}{
	{
		name: "typical sentance",
		raw:  "$PKLID,00,100,2000,15,00,*6D",
		msg:  PKLID{
			SentanceVersion: "00",
			Fleet:           "100",
			ID:              "2000",
			Status:          "15",
			Extension:       "00",
		},
	},
}

func TestPKLID(t *testing.T) {
	for _, tt := range pklidtests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t,err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				pklid := m.(PKLID)
				pklid.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, pklid)
			}
		})
	}
}
	
