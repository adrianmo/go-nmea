package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tagblocktests = []struct {
	//name: "Tagblock ok",
	//raw: "",
	name string
	raw  string
	err  string
	msg  TagBlock
}{
	{

		name: "Test NMEA tag block",
		raw:  "\\s:Satelite_1,c:1553390539*62\\!AIVDM,1,1,,A,13M@ah0025QdPDTCOl`K6`nV00Sv,0*52",
		msg: TagBlock{
			Time:   1553390539,
			Source: "Satelite_1",
		},
	},
	{

		name: "Test NMEA tag block with head",
		raw:  "UdPbC?\\s:satelite,c:1564827317*25\\!AIVDM,1,1,,A,19NSRaP02A0fo91kwnaMKbjR08:J,0*15",
		msg: TagBlock{
			Time:   1564827317,
			Source: "satelite",
			Head:   "UdPbC?",
		},
	},
	{

		name: "Test unknown tag",
		raw:  "UdPbC?\\x:NorSat_1,c:1564827317*42\\!AIVDM,1,1,,A,19NSRaP02A0fo91kwnaMKbjR08:J,0*15",
		msg: TagBlock{
			Time:   1564827317,
			Source: "",
			Head:   "UdPbC?",
		},
	},
	{

		name: "Test all input types",
		raw:  "UdPbC?\\s:satelite,c:1564827317,r:1553390539,d:ara,g:bulk,n:13,t:helloworld*3F\\!AIVDM,1,1,,A,19NSRaP02A0fo91kwnaMKbjR08:J,0*15",
		msg: TagBlock{
			Time:         1564827317,
			RelativeTime: 1553390539,
			Destination:  "ara",
			Grouping:     "bulk",
			Source:       "satelite",
			Head:         "UdPbC?",
			Text:         "helloworld",
			LineCount:    13,
		},
	},
	{

		name: "Test Invalid checksum",
		raw:  "UdPbC?\\s:satelite,c:1564827317*49\\!AIVDM,1,1,,A,19NSRaP02A0fo91kwnaMKbjR08:J,0*15",
		err:  "nmea: tagblock checksum mismatch [25 != 49]",
	},
	{

		name: "Test no checksum",
		raw:  "UdPbC?\\s:satelite,c:156482731749\\!AIVDM,1,1,,A,19NSRaP02A0fo91kwnaMKbjR08:J,0*15",
		err:  "nmea: tagblock does not contain checksum separator",
	},
	{

		name: "Test invalid timestamp",
		raw:  "UdPbC?\\s:satelite,c:gjadslkg*30\\!AIVDM,1,1,,A,19NSRaP02A0fo91kwnaMKbjR08:J,0*15",
		err:  "nmea: tagblock unable to parse uint32 [c:gjadslkg]",
	},
	{

		name: "Test invalid linecount",
		raw:  "UdPbC?\\s:satelite,n:gjadslkg*3D\\!AIVDM,1,1,,A,19NSRaP02A0fo91kwnaMKbjR08:J,0*15",
		err:  "nmea: tagblock unable to parse uint32 [n:gjadslkg]",
	},
	{

		name: "Test invalid relative time",
		raw:  "UdPbC?\\s:satelite,r:gjadslkg*21\\!AIVDM,1,1,,A,19NSRaP02A0fo91kwnaMKbjR08:J,0*15",
		err:  "nmea: tagblock unable to parse uint32 [r:gjadslkg]",
	},
}

func TestTagBlock(t *testing.T) {
	for _, tt := range tagblocktests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				vdm := m.(VDMVDO)
				assert.Equal(t, tt.msg, vdm.BaseSentence.TagBlock)
			}
		})
	}
}
