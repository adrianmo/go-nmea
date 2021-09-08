package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var mwvtests = []struct {
	name string
	raw  string
	err  string
	msg  MWV
}{
	{
		name: "good sentence",
		raw:  "$WIMWV,12.1,T,10.1,N,A*27",
		msg: MWV{
			WindAngle:     12.1,
			Reference:     "T",
			WindSpeed:     10.1,
			WindSpeedUnit: "N",
			StatusValid:   true,
		},
	},
	{
		name: "invalid data",
		raw:  "$WIMWV,,T,,N,V*32",
		msg: MWV{
			WindAngle:     0,
			Reference:     "T",
			WindSpeed:     0,
			WindSpeedUnit: "N",
			StatusValid:   false,
		},
	},
}

func TestMWV(t *testing.T) {
	for _, tt := range mwvtests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				mwv := m.(MWV)
				mwv.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, mwv)
			}
		})
	}
}
