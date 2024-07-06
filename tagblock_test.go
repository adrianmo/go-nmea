package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tagblocktests = []struct {
	name  string
	raw   string
	err   string
	block TagBlock
	len   int
}{
	{

		name: "Test NMEA tag block",
		raw:  "\\s:Satelite_1,c:1553390539*62\\!AIVDM,1,2,3",
		block: TagBlock{
			Time:   1553390539,
			Source: "Satelite_1",
		},
		len: 30,
	},
	{

		name: "Test NMEA tag block with head",
		raw:  "\\s:satelite,c:1564827317*25\\!AIVDM,1,2,3",
		block: TagBlock{
			Time:   1564827317,
			Source: "satelite",
		},
		len: 28,
	},
	{

		name: "Test unknown tag",
		raw:  "\\x:NorSat_1,c:1564827317*42\\!AIVDM,1,2,3",
		block: TagBlock{
			Time:   1564827317,
			Source: "",
		},
		len: 28,
	},
	{
		name: "Test unix timestamp",
		raw:  "\\x:NorSat_1,c:1564827317*42\\!AIVDM,1,2,3",
		block: TagBlock{
			Time:   1564827317,
			Source: "",
		},
		len: 28,
	},
	{

		name: "Test milliseconds timestamp",
		raw:  "\\x:NorSat_1,c:1564827317000*72\\!AIVDM,1,2,3",
		block: TagBlock{
			Time:   1564827317000,
			Source: "",
		},
		len: 31,
	},
	{

		name: "Test all input types",
		raw:  "\\s:satelite,c:1564827317,r:1553390539,d:ara,g:bulk,n:13,t:helloworld*3F\\!AIVDM,1,2,3",
		block: TagBlock{
			Time:         1564827317,
			RelativeTime: 1553390539,
			Destination:  "ara",
			Grouping:     "bulk",
			Source:       "satelite",
			Text:         "helloworld",
			LineCount:    13,
		},
		len: 72,
	},
	{

		name: "Test empty tag in tagblock",
		raw:  "\\s:satelite,,r:1553390539,d:ara,g:bulk,n:13,t:helloworld*68\\!AIVDM,1,2,3",
		err:  "nmea: tagblock field is malformed (should be <key>:<value>) []",
	},
	{

		name: "Test Invalid checksum",
		raw:  "\\s:satelite,c:1564827317*49\\!AIVDM,1,2,3",
		err:  "nmea: tagblock checksum mismatch [25 != 49]",
	},
	{

		name: "Test no checksum",
		raw:  "\\s:satelite,c:156482731749\\!AIVDM,1,2,3",
		err:  "nmea: tagblock does not contain checksum separator",
	},
	{

		name: "Test invalid timestamp",
		raw:  "\\s:satelite,c:gjadslkg*30\\!AIVDM,1,2,3",
		err:  "nmea: tagblock unable to parse uint64 [gjadslkg]",
	},
	{

		name: "Test invalid linecount",
		raw:  "\\s:satelite,n:gjadslkg*3D\\!AIVDM,1,2,3",
		err:  "nmea: tagblock unable to parse uint64 [gjadslkg]",
	},
	{

		name: "Test invalid relative time",
		raw:  "\\s:satelite,r:gjadslkg*21\\!AIVDM,1,2,3",
		err:  "nmea: tagblock unable to parse uint64 [gjadslkg]",
	},
	{
		name: "Test no tagblock",
		raw:  "!AIVDM,1,2,3",
	},
}

func TestParseTagBlock(t *testing.T) {
	for _, tt := range tagblocktests {
		t.Run(tt.name, func(t *testing.T) {
			b, n, err := ParseTagBlock(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.block, b)
			assert.Equal(t, tt.len, n)
		})
	}
}
