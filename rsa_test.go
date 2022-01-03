package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRSA(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  RSA
	}{
		{
			name: "good sentence 1",
			raw:  "$IIRSA,10.5,A,0.4,A*70",
			msg: RSA{
				StarboardRudderAngle:       10.5,
				StarboardRudderAngleStatus: StatusValid,
				PortRudderAngle:            0.4,
				PortRudderAngleStatus:      StatusValid,
			},
		},
		{
			name: "good sentence 2",
			raw:  "$IIRSA,10.5,A,,V*4D",
			msg: RSA{
				StarboardRudderAngle:       10.5,
				StarboardRudderAngleStatus: StatusValid,
				PortRudderAngle:            0,
				PortRudderAngleStatus:      StatusInvalid,
			},
		},
		{
			name: "invalid nmea: StarboardRudderAngleStatus",
			raw:  "$IIRSA,10.5,x,,V*74",
			err:  "nmea: IIRSA invalid starboard rudder angle status: x",
		},
		{
			name: "invalid nmea: PortRudderAngleStatus",
			raw:  "$IIRSA,10.5,A,,x*63",
			err:  "nmea: IIRSA invalid port rudder angle status: x",
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
				rsa := m.(RSA)
				rsa.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, rsa)
			}
		})
	}
}
