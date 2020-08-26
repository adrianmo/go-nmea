package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tagblocktests = []struct {
	name string
	raw  string
	err  string
	msg  TagBlock
}{
	{

		name: "Test NMEA tag block",
		raw:  "s:Satelite_1,c:1553390539*62",
		msg: TagBlock{
			Time:   1553390539,
			Source: "Satelite_1",
		},
	},
	{

		name: "Test NMEA tag block with head",
		raw:  "s:satelite,c:1564827317*25",
		msg: TagBlock{
			Time:   1564827317,
			Source: "satelite",
		},
	},
	{

		name: "Test unknown tag",
		raw:  "x:NorSat_1,c:1564827317*42",
		msg: TagBlock{
			Time:   1564827317,
			Source: "",
		},
	},
	{
		name: "Test unix timestamp",
		raw:  "x:NorSat_1,c:1564827317*42",
		msg: TagBlock{
			Time:   1564827317,
			Source: "",
		},
	},
	{

		name: "Test milliseconds timestamp",
		raw:  "x:NorSat_1,c:1564827317000*72",
		msg: TagBlock{
			Time:   1564827317000,
			Source: "",
		},
	},
	{

		name: "Test all input types",
		raw:  "s:satelite,c:1564827317,r:1553390539,d:ara,g:bulk,n:13,t:helloworld*3F",
		msg: TagBlock{
			Time:         1564827317,
			RelativeTime: 1553390539,
			Destination:  "ara",
			Grouping:     "bulk",
			Source:       "satelite",
			Text:         "helloworld",
			LineCount:    13,
		},
	},
	{

		name: "Test empty tag in tagblock",
		raw:  "s:satelite,,r:1553390539,d:ara,g:bulk,n:13,t:helloworld*68",
		err:  "nmea: tagblock field is malformed (should be <key>:<value>) []",
	},
	{

		name: "Test Invalid checksum",
		raw:  "s:satelite,c:1564827317*49",
		err:  "nmea: tagblock checksum mismatch [25 != 49]",
	},
	{

		name: "Test no checksum",
		raw:  "s:satelite,c:156482731749",
		err:  "nmea: tagblock does not contain checksum separator",
	},
	{

		name: "Test invalid timestamp",
		raw:  "s:satelite,c:gjadslkg*30",
		err:  "nmea: tagblock unable to parse uint64 [gjadslkg]",
	},
	{

		name: "Test invalid linecount",
		raw:  "s:satelite,n:gjadslkg*3D",
		err:  "nmea: tagblock unable to parse uint64 [gjadslkg]",
	},
	{

		name: "Test invalid relative time",
		raw:  "s:satelite,r:gjadslkg*21",
		err:  "nmea: tagblock unable to parse uint64 [gjadslkg]",
	},
}

func TestTagBlock(t *testing.T) {
	for _, tt := range tagblocktests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := parseTagBlock(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.msg, m)
			}
		})
	}
}
