package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var gpzdatests = []struct {
	name string
	raw  string
	err  string
	msg  GPZDA
}{
	{
		name: "good sentence",
		raw:  "$GPZDA,172809.456,12,07,1996,00,00*57",
		msg: GPZDA{
			Time: Time{
				Valid:       true,
				Hour:        17,
				Minute:      28,
				Second:      9,
				Millisecond: 456,
			},
			Day:           12,
			Month:         7,
			Year:          1996,
			OffsetHours:   0,
			OffsetMinutes: 0,
		},
	},
	{
		name: "invalid day",
		raw:  "$GPZDA,220516,D,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*76",
		err:  "nmea: GPZDA invalid day: D",
	},
	{
		name: "wrong sentence",
		raw:  "$GPXTE,A,A,4.07,L,N*6D",
		err:  "nmea: GPZDA invalid prefix: GPXTE",
	},
}

func TestGPZDA(t *testing.T) {
	for _, tt := range gpzdatests {
		t.Run(tt.name, func(t *testing.T) {
			sent, err := ParseSentence(tt.raw)
			assert.NoError(t, err)
			gpzda, err := NewGPZDA(sent)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				gpzda.Sent = Sent{}
				assert.Equal(t, tt.msg, gpzda)
			}
		})
	}
}
