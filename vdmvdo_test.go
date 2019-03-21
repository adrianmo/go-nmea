package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var vdmtests = []struct {
	name string
	raw  string
	err  string
	msg  VDMVDO
}{
	{
		name: "Good single fragment message",
		raw:  "!AIVDM,1,1,,A,13aGt0PP0jPN@9fMPKVDJgwfR>`<,0*55",
		msg: VDMVDO{
			NumFragments:   1,
			FragmentNumber: 1,
			MessageID:      0,
			Channel:        "A",
			Payload:        []byte{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0},
		},
	},
	{
		name: "Good single fragment message with padding",
		raw:  "!AIVDM,1,1,,A,H77nSfPh4U=<E`H4U8G;:222220,2*1F",
		msg: VDMVDO{
			NumFragments:   1,
			FragmentNumber: 1,
			MessageID:      0,
			Channel:        "A",
			Payload:        []byte{0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		},
	},
	{
		name: "Good multipart fragment",
		raw:  "!AIVDM,2,2,4,B,00000000000,2*23",
		msg: VDMVDO{
			NumFragments:   2,
			FragmentNumber: 2,
			MessageID:      4,
			Channel:        "B",
			Payload:        []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	},
	{
		name: "Empty payload",
		raw:  "!AIVDM,1,1,,1,,0*56",
		msg: VDMVDO{
			NumFragments:   1,
			FragmentNumber: 1,
			MessageID:      0,
			Channel:        "1",
			Payload:        []byte{},
		},
	},
	{
		name: "Invalid number of fragments",
		raw:  "!AIVDM,x,1,,1,000 00,0*0F",
		err:  "nmea: AIVDM invalid number of fragments: x",
	},
	{
		name: "Invalid symbol in payload",
		raw:  "!AIVDM,1,1,,1,000 00,0*46",
		err:  "nmea: AIVDM invalid payload: data byte",
	},
	{
		name: "Negative number of fill bits",
		raw:  "!AIVDM,1,1,,1,000,-3*48",
		err:  "nmea: AIVDM invalid payload: fill bits",
	},
	{
		name: "Too high number of fill bits",
		raw:  "!AIVDO,1,1,,1,000,20*56",
		err:  "nmea: AIVDO invalid payload: fill bits",
	},
	{
		name: "Negative number of bits",
		raw:  "!AIVDM,1,1,,1,,2*54",
		err:  "nmea: AIVDM invalid payload: num bits",
	},
}

func TestVDM(t *testing.T) {
	for _, tt := range vdmtests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)

			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				vdm := m.(VDMVDO)
				vdm.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, vdm)
			}
		})
	}
}
