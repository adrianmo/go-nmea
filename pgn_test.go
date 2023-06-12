package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPGN(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  PGN
	}{
		{
			name: "good sentence",
			raw:  "$MXPGN,01F112,2807,FC7FFF7FFF168012*11",
			msg: PGN{
				PGN:      127250, //  0x1F112 Vessel Heading
				IsSend:   false,
				Priority: 2,
				Address:  7,
				Data:     []byte{0xFC, 0x7f, 0xFF, 0x7f, 0xFF, 0x16, 0x80, 0x12},
			},
		},
		{
			name: "invalid number of fields",
			raw:  "$MXPGN,01F112,FC7FFF7FFF168012*30",
			err:  "nmea: MXPGN invalid fields: invalid number of fields in sentence",
		},
		{
			name: "invalid PGN field",
			raw:  "$MXPGN,0xF112,2807,FC7FFF7FFF168012*58",
			err:  "failed to parse PGN field, err: strconv.ParseUint: parsing \"0xF112\": invalid syntax",
		},
		{
			name: "invalid attributes field",
			raw:  "$MXPGN,01F112,x807,FC7FFF7FFF168012*5b",
			err:  "failed to parse attributes field, err: strconv.ParseUint: parsing \"x807\": invalid syntax",
		},
		{
			name: "invalid data length field",
			raw:  "$MXPGN,01F112,2207,FC7FFF7FFF168012*1b",
			err:  "nmea: MXPGN invalid dlc: data length does not match actual data length",
		},
		{
			name: "invalid hex data",
			raw:  "$MXPGN,01F112,2807,xC7FFF7FFF168012*2f",
			err:  "failed to decode data, err: encoding/hex: invalid byte: U+0078 'x'",
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
				pgrme := m.(PGN)
				pgrme.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, pgrme)
			}
		})
	}
}
