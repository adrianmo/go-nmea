package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestABM(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  ABM
	}{
		{
			name: "Good single fragment message",
			raw:  "!AIABM,26,2,1,3381581370,3,8,177KQJ5000G?tO`K>RA1wUbN0TKH,0*02",
			msg: ABM{
				NumFragments:     26,
				FragmentNumber:   2,
				MessageID:        1,
				MMSI:             "3381581370",
				Channel:          "3",
				VDLMessageNumber: 8,
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
			raw:  "!AIABM,26,2,1,3381581370,3,8,H77nSfPh4U=<E`H4U8G;:222220,2*42",
			msg: ABM{
				NumFragments:     26,
				FragmentNumber:   2,
				MessageID:        1,
				MMSI:             "3381581370",
				Channel:          "3",
				VDLMessageNumber: 8,
				Payload: []byte{
					0, 1, 1, 0, 0, 0, 0, 0, 0, 1,
					1, 1, 0, 0, 0, 1, 1, 1, 1, 1,
					0, 1, 1, 0, 1, 0, 0, 0, 1, 1,
					1, 0, 1, 1, 1, 0, 1, 0, 0, 0,
					0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
					0, 1, 0, 0, 1, 0, 0, 1, 0, 1,
					0, 0, 1, 1, 0, 1, 0, 0, 1, 1,
					0, 0, 0, 1, 0, 1, 0, 1, 1, 0,
					1, 0, 0, 0, 0, 1, 1, 0, 0, 0,
					0, 0, 0, 1, 0, 0, 1, 0, 0, 1,
					0, 1, 0, 0, 1, 0, 0, 0, 0, 1,
					0, 1, 1, 1, 0, 0, 1, 0, 1, 1,
					0, 0, 1, 0, 1, 0, 0, 0, 0, 0,
					1, 0, 0, 0, 0, 0, 1, 0, 0, 0,
					0, 0, 1, 0, 0, 0, 0, 0, 1, 0,
					0, 0, 0, 0, 1, 0, 0, 0, 0, 0,
				},
			},
		},
		{
			name: "Empty payload",
			raw:  "!AIABM,26,2,1,3381581370,3,8,,0*7b",
			msg: ABM{
				NumFragments:     26,
				FragmentNumber:   2,
				MessageID:        1,
				MMSI:             "3381581370",
				Channel:          "3",
				VDLMessageNumber: 8,
				Payload:          []byte{},
			},
		},
		{
			name: "Invalid number of fragments",
			raw:  "!AIABM,x,2,1,3381581370,3,8,177KQJ5000G?tO`K>RA1wUbN0TKH,0*7e",
			err:  "nmea: AIABM invalid number of fragments: x",
		},
		{
			name: "Invalid VDLMessageNumber",
			raw:  "!AIABM,26,2,1,3381581370,3,x,177KQJ5000G?tO`K>RA1wUbN0TKH,0*42",
			err:  "nmea: AIABM invalid VDL message number: x",
		},
		{
			name: "Invalid symbol in payload",
			raw:  "!AIABM,26,2,1,3381581370,3,8,1 1,0*5b",
			err:  "nmea: AIABM invalid payload: data byte",
		},
		{
			name: "Negative number of fill bits",
			raw:  "!AIABM,26,2,1,3381581370,3,8,177KQJ5000G?tO`K>RA1wUbN0TKH,-1*2e",
			err:  "nmea: AIABM invalid payload: fill bits",
		},
		{
			name: "Too high number of fill bits",
			raw:  "!AIABM,26,2,1,3381581370,3,8,177KQJ5000G?tO`K>RA1wUbN0TKH,20*30",
			err:  "nmea: AIABM invalid payload: fill bits",
		},
		{
			name: "Negative number of bits",
			raw:  "!AIABM,26,2,1,3381581370,3,8,,2*79",
			err:  "nmea: AIABM invalid payload: num bits",
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
				abm := m.(ABM)
				abm.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, abm)
			}
		})
	}
}
