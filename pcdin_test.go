package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPCDIN(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  PCDIN
	}{
		{
			name: "good sentence",
			raw:  "$PCDIN,01F112,000C72EA,09,28C36A0000B40AFD*56",
			msg: PCDIN{
				PGN:       127250, //  0x1F112 Vessel Heading
				Timestamp: 815850,
				Source:    9,
				Data:      []byte{0x28, 0xC3, 0x6A, 0x00, 0x00, 0xB4, 0x0A, 0xFD},
			},
		},
		{
			name: "invalid number of fields",
			raw:  "$PCDIN,01F112,000C72EA,28C36A0000B40AFD*73",
			err:  "nmea: PCDIN invalid fields: invalid number of fields in sentence",
		},
		{
			name: "invalid PGN field",
			raw:  "$PCDIN,x1F112,000C72EA,09,28C36A0000B40AFD*1e",
			err:  "failed to parse PGN field, err: strconv.ParseUint: parsing \"x1F112\": invalid syntax",
		},
		{
			name: "invalid timestamp field",
			raw:  "$PCDIN,01F112,x00C72EA,09,28C36A0000B40AFD*1e",
			err:  "failed to parse timestamp field, err: strconv.ParseUint: parsing \"x00C72EA\": invalid syntax",
		},
		{
			name: "invalid source field",
			raw:  "$PCDIN,01F112,000C72EA,x9,28C36A0000B40AFD*1e",
			err:  "failed to parse source field, err: strconv.ParseUint: parsing \"x9\": invalid syntax",
		},
		{
			name: "invalid hex data",
			raw:  "$PCDIN,01F112,000C72EA,09,x8C36A0000B40AFD*1c",
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
				pgrme := m.(PCDIN)
				pgrme.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, pgrme)
			}
		})
	}
}
