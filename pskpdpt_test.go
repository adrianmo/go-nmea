package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPSKPDPT(t *testing.T) {
	var testcases = []struct {
		name string
		raw  string
		err  string
		msg  PSKPDPT
	}{
		{
			name: "good sentence, empty location",
			raw:  "$PSKPDPT,0002.5,+00.0,0010,10,03,*77",
			msg: PSKPDPT{
				Depth:              2.5,
				Offset:             0,
				RangeScale:         10,
				BottomEchoStrength: 10,
				ChannelNumber:      03,
				TransducerLocation: "",
			},
		},
		{
			name: "good sentence",
			raw:  "$PSKPDPT,0002.5,-01.1,0010,10,03,AFT*22",
			msg: PSKPDPT{
				Depth:              2.5,
				Offset:             -1.1,
				RangeScale:         10,
				BottomEchoStrength: 10,
				ChannelNumber:      03,
				TransducerLocation: "AFT",
			},
		},
		{
			name: "invalid nmea: Depth",
			raw:  "$PSKPDPT,x0002.5,+00.0,0010,10,03,*0f",
			err:  "nmea: PSKPDPT invalid depth: x0002.5",
		},
		{
			name: "invalid nmea: Offset",
			raw:  "$PSKPDPT,0002.5,+x00.0,0010,10,03,*0f",
			err:  "nmea: PSKPDPT invalid offset: +x00.0",
		},
		{
			name: "invalid nmea: RangeScale",
			raw:  "$PSKPDPT,0002.5,+00.0,x0010,10,03,*0f",
			err:  "nmea: PSKPDPT invalid range scale: x0010",
		},
		{
			name: "invalid nmea: BottomEchoStrength",
			raw:  "$PSKPDPT,0002.5,+00.0,0010,10x,03,*0f",
			err:  "nmea: PSKPDPT invalid bottom echo strength: 10x",
		},
		{
			name: "invalid nmea: ChannelNumber",
			raw:  "$PSKPDPT,0002.5,+00.0,0010,10,0x3,*0f",
			err:  "nmea: PSKPDPT invalid channel number: 0x3",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				sentence := m.(PSKPDPT)
				sentence.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, sentence)
			}
		})
	}
}
