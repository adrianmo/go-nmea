package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pgrmetests = []struct {
	name string
	raw  string
	err  string
	msg  PGRME
}{
	{
		name: "good sentence",
		raw:  "$PGRME,3.3,M,4.9,M,6.0,M*25",
		msg: PGRME{
			Horizontal: 3.3,
			Vertical:   4.9,
			Spherical:  6,
		},
	},
	{
		name: "invalid horizontal error",
		raw:  "$PGRME,A,M,4.9,M,6.0,M*4A",
		err:  "nmea: PGRME invalid horizontal error: A",
	},
	{
		name: "invalid vertical error",
		raw:  "$PGRME,3.3,M,A,M,6.0,M*47",
		err:  "nmea: PGRME invalid vertical error: A",
	},
	{
		name: "invalid vertical error unit",
		raw:  "$PGRME,3.3,M,4.9,A,6.0,M*29",
		err:  "nmea: PGRME invalid vertical error unit: A",
	},
	{
		name: "invalid spherical error",
		raw:  "$PGRME,3.3,M,4.9,M,A,M*4C",
		err:  "nmea: PGRME invalid spherical error: A",
	},
	{
		name: "invalid spherical error unit",
		raw:  "$PGRME,3.3,M,4.9,M,6.0,A*29",
		err:  "nmea: PGRME invalid spherical error unit: A",
	},
}

func TestPGRME(t *testing.T) {
	for _, tt := range pgrmetests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				gprme := m.(PGRME)
				gprme.Sent = Sent{}
				assert.Equal(t, tt.msg, gprme)
			}
		})
	}
}
