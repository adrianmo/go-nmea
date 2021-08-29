package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sentencetests = []struct {
	name     string
	raw      string
	datatype string
	talkerid string
	prefix   string
	err      string
	sent     BaseSentence
}{
	{
		name:     "checksum ok",
		raw:      "$GPFOO,1,2,3.3,x,y,zz,*51",
		datatype: "FOO",
		talkerid: "GP",
		prefix:   "GPFOO",
		sent: BaseSentence{
			Talker:   "GP",
			Type:     "FOO",
			Fields:   []string{"1", "2", "3.3", "x", "y", "zz", ""},
			Checksum: "51",
			Raw:      "$GPFOO,1,2,3.3,x,y,zz,*51",
		},
	},
	{
		name:     "trim leading and trailing spaces",
		raw:      "   $GPFOO,1,2,3.3,x,y,zz,*51   ",
		datatype: "FOO",
		talkerid: "GP",
		prefix:   "GPFOO",
		sent: BaseSentence{
			Talker:   "GP",
			Type:     "FOO",
			Fields:   []string{"1", "2", "3.3", "x", "y", "zz", ""},
			Checksum: "51",
			Raw:      "$GPFOO,1,2,3.3,x,y,zz,*51",
		},
	},
	{
		name:     "good parsing",
		raw:      "$GPRMC,235236,A,3925.9479,N,11945.9211,W,44.7,153.6,250905,15.2,E,A*0C",
		datatype: "RMC",
		talkerid: "GP",
		prefix:   "GPRMC",
		sent: BaseSentence{
			Talker:   "GP",
			Type:     "RMC",
			Fields:   []string{"235236", "A", "3925.9479", "N", "11945.9211", "W", "44.7", "153.6", "250905", "15.2", "E", "A"},
			Checksum: "0C",
			Raw:      "$GPRMC,235236,A,3925.9479,N,11945.9211,W,44.7,153.6,250905,15.2,E,A*0C",
		},
	},
	{
		name:     "valid NMEA 4.10 TAG Block",
		raw:      "\\s:Satelite_1,c:1553390539*62\\!AIVDM,1,1,,A,13M@ah0025QdPDTCOl`K6`nV00Sv,0*52",
		datatype: "VDM",
		talkerid: "AI",
		prefix:   "AIVDM",
		sent: BaseSentence{
			Talker:   "AI",
			Type:     "VDM",
			Fields:   []string{"1", "1", "", "A", "13M@ah0025QdPDTCOl`K6`nV00Sv", "0"},
			Checksum: "52",
			Raw:      "!AIVDM,1,1,,A,13M@ah0025QdPDTCOl`K6`nV00Sv,0*52",
			TagBlock: TagBlock{
				Time:   1553390539,
				Source: "Satelite_1",
			},
		},
	},
	{
		name: "checksum bad",
		raw:  "$GPFOO,1,2,3.4,x,y,zz,*51",
		err:  "nmea: sentence checksum mismatch [56 != 51]",
	},
	{
		name: "bad start character",
		raw:  "%GPFOO,1,2,3,x,y,z*1A",
		err:  "nmea: sentence does not start with a '$' or '!'",
	},
	{
		name: "bad checksum delimiter",
		raw:  "$GPFOO,1,2,3,x,y,z",
		err:  "nmea: sentence does not contain checksum separator",
	},
	{
		name: "no start delimiter",
		raw:  "abc$GPRMC,235236,A,3925.9479,N,11945.9211,W,44.7,153.6,250905,15.2,E,A*0C",
		err:  "nmea: sentence does not start with a '$' or '!'",
	},
	{
		name: "no contain delimiter",
		raw:  "GPRMC,235236,A,3925.9479,N,11945.9211,W,44.7,153.6,250905,15.2,E,A*0C",
		err:  "nmea: sentence does not start with a '$' or '!'",
	},
	{
		name: "another bad checksum",
		raw:  "$GPRMC,235236,A,3925.9479,N,11945.9211,W,44.7,153.6,250905,15.2,E,A*0A",
		err:  "nmea: sentence checksum mismatch [0C != 0A]",
	},
	{
		name: "missing TAG Block start delimiter",
		raw:  "s:Satelite_1,c:1553390539*62\\!AIVDM,1,1,,A,13M@ah0025QdPDTCOl`K6`nV00Sv,0*52",
		err:  "nmea: sentence does not start with a '$' or '!'",
	},
	{
		name: "missing TAG Block end delimiter",
		raw:  "\\s:Satelite_1,c:1553390539*62!AIVDM,1,1,,A,13M@ah0025QdPDTCOl`K6`nV00Sv,0*52",
		err:  "nmea: sentence does not start with a '$' or '!'",
	},
	{
		name: "invalid TAG Block contents",
		raw:  "\\\\!AIVDM,1,1,,A,13M@ah0025QdPDTCOl`K6`nV00Sv,0*52",
		err:  "nmea: tagblock does not contain checksum separator",
	},
}

func TestSentences(t *testing.T) {
	for _, tt := range sentencetests {
		t.Run(tt.name, func(t *testing.T) {
			sent, err := parseSentence(tt.raw)
			if tt.err != "" {
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.sent, sent)
				assert.Equal(t, tt.sent.Raw, sent.String())
				assert.Equal(t, tt.datatype, sent.DataType())
				assert.Equal(t, tt.talkerid, sent.TalkerID())
				assert.Equal(t, tt.prefix, sent.Prefix())
			}
		})
	}
}

var prefixtests = []struct {
	name   string
	prefix string
	talker string
	typ    string
}{
	{
		name:   "normal prefix",
		prefix: "GPRMC",
		talker: "GP",
		typ:    "RMC",
	},
	{
		name:   "GNSS prefix",
		prefix: "GNGSA",
		talker: "GN",
		typ:    "GSA",
	},
	{
		name:   "missing type",
		prefix: "GP",
		talker: "GP",
		typ:    "",
	},
	{
		name:   "one character",
		prefix: "X",
		talker: "X",
		typ:    "",
	},
	{
		name:   "proprietary talker",
		prefix: "PGRME",
		talker: "P",
		typ:    "GRME",
	},
	{
		name:   "short proprietary talker",
		prefix: "PX",
		talker: "P",
		typ:    "X",
	},
}

func TestPrefix(t *testing.T) {
	for _, tt := range prefixtests {
		t.Run(tt.name, func(t *testing.T) {
			talker, typ := parsePrefix(tt.prefix)
			assert.Equal(t, tt.talker, talker)
			assert.Equal(t, tt.typ, typ)
		})
	}
}

var parsetests = []struct {
	name string
	raw  string
	err  string
	msg  interface{}
}{
	{
		name: "bad sentence",
		raw:  "SDFSD,2340dfmswd",
		err:  "nmea: sentence does not start with a '$' or '!'",
	},
	{
		name: "bad sentence type",
		raw:  "$INVALID,123,123,*7D",
		err:  "nmea: sentence prefix 'INVALID' not supported",
	},
	{
		name: "bad encapsulated sentence type",
		raw:  "!INVALID,1,2,*7E",
		err:  "nmea: sentence prefix 'INVALID' not supported",
	},
}

func TestParse(t *testing.T) {
	for _, tt := range parsetests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.msg, m)
			}
		})
	}
}
