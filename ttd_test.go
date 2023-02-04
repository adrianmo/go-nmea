package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTTD(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  TTD
	}{
		{
			name: "Good single fragment message",
			raw:  "!RATTD,1A,01,1,177KQJ5000G?tO`K>RA1wUbN0TKH,0*72",
			msg: TTD{
				NumFragments:   26,
				FragmentNumber: 1,
				MessageID:      1,
				Payload: []byte{
					0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x1, // 10
					0x1, 0x1, 0x0, 0x0, 0x0, 0x1, 0x1, 0x1, 0x0, 0x1, // 20
					0x1, 0x0, 0x1, 0x1, 0x1, 0x0, 0x0, 0x0, 0x0, 0x1, // 30
					0x0, 0x1, 0x1, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x1, // 40
					0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 50
					0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 60
					0x0, 0x1, 0x0, 0x1, 0x1, 0x1, 0x0, 0x0, 0x1, 0x1, // 70
					0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x0, 0x0, 0x0, 0x1, // 80
					0x1, 0x1, 0x1, 0x1, 0x1, 0x0, 0x1, 0x0, 0x0, 0x0, // 90
					0x0, 0x1, 0x1, 0x0, 0x1, 0x1, 0x0, 0x0, 0x1, 0x1, // 100
					0x1, 0x0, 0x1, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x1, // 110
					0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, // 120
					0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x0, 0x0, 0x1, // 130
					0x0, 0x1, 0x1, 0x0, 0x1, 0x0, 0x1, 0x0, 0x0, 0x1, // 140
					0x1, 0x1, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, // 150
					0x1, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x1, 0x1, 0x0, // 160
					0x1, 0x1, 0x0, 0x1, 0x1, 0x0, 0x0, 0x0, // 168
				},
			},
		},
		{
			name: "Good single fragment message with padding",
			raw:  "!RATTD,1A,01,1,H77nSfPh4U=<E`H4U8G;:222220,2*32",
			msg: TTD{
				NumFragments:   26,
				FragmentNumber: 1,
				MessageID:      1,
				Payload:        []byte{0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			},
		},
		{
			name: "Empty payload",
			raw:  "!RATTD,1A,01,1,,,0*27",
			msg: TTD{
				NumFragments:   26,
				FragmentNumber: 1,
				MessageID:      1,
				Payload:        []byte{},
			},
		},
		{
			name: "Invalid number of fragments",
			raw:  "!RATTD,x,01,1,177KQJ5000G?tO`K>RA1wUbN0TKH,0*7A",
			err:  "nmea: RATTD invalid number of fragments: x",
		},
		{
			name: "Invalid symbol in payload",
			raw:  "!RATTD,1A,01,1,1 1,0*2b",
			err:  "nmea: RATTD invalid payload: data byte",
		},
		{
			name: "Negative number of fill bits",
			raw:  "!RATTD,1A,01,1,177KQJ5000G?tO`K>RA1wUbN0TKH,-1*5e",
			err:  "nmea: RATTD invalid payload: fill bits",
		},
		{
			name: "Too high number of fill bits",
			raw:  "!RATTD,1A,01,1,177KQJ5000G?tO`K>RA1wUbN0TKH,20*40",
			err:  "nmea: RATTD invalid payload: fill bits",
		},
		{
			name: "Negative number of bits",
			raw:  "!RATTD,1A,01,1,,2*09",
			err:  "nmea: RATTD invalid payload: num bits",
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
				ttd := m.(TTD)
				ttd.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, ttd)
			}
		})
	}
}
