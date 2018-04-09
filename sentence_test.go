package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sentencetests = []struct {
	name string
	raw  string
	err  string
	sent Sent
}{
	{
		name: "checksum ok",
		raw:  "$GPFOO,1,2,3.3,x,y,zz,*51",
		sent: Sent{
			Type:     "GPFOO",
			Fields:   []string{"1", "2", "3.3", "x", "y", "zz", ""},
			Checksum: "51",
			Raw:      "$GPFOO,1,2,3.3,x,y,zz,*51",
		},
	},
	{
		name: "good parsing",
		raw:  "$GPRMC,235236,A,3925.9479,N,11945.9211,W,44.7,153.6,250905,15.2,E,A*0C",
		sent: Sent{
			Type:     "GPRMC",
			Fields:   []string{"235236", "A", "3925.9479", "N", "11945.9211", "W", "44.7", "153.6", "250905", "15.2", "E", "A"},
			Checksum: "0C",
			Raw:      "$GPRMC,235236,A,3925.9479,N,11945.9211,W,44.7,153.6,250905,15.2,E,A*0C",
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
		err:  "nmea: sentence does not start with a '$'",
	},
	{
		name: "bad checksum delimiter",
		raw:  "$GPFOO,1,2,3,x,y,z",
		err:  "nmea: sentence does not contain checksum separator",
	},
	{
		name: "no start delimiter",
		raw:  "abc$GPRMC,235236,A,3925.9479,N,11945.9211,W,44.7,153.6,250905,15.2,E,A*0C",
		err:  "nmea: sentence does not start with a '$'",
	},
	{
		name: "no contain delimiter",
		raw:  "GPRMC,235236,A,3925.9479,N,11945.9211,W,44.7,153.6,250905,15.2,E,A*0C",
		err:  "nmea: sentence does not start with a '$'",
	},
	{
		name: "another bad checksum",
		raw:  "$GPRMC,235236,A,3925.9479,N,11945.9211,W,44.7,153.6,250905,15.2,E,A*0A",
		err:  "nmea: sentence checksum mismatch [0C != 0A]",
	},
}

func TestSentences(t *testing.T) {
	for _, tt := range sentencetests {
		t.Run(tt.name, func(t *testing.T) {
			sent, err := ParseSentence(tt.raw)
			if tt.err != "" {
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.sent, sent)
			}
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
		err:  "nmea: sentence does not start with a '$'",
	},
	{
		name: "bad sentence type",
		raw:  "$INVALID,123,123,*7D",
		err:  "nmea: sentence type 'INVALID' not implemented",
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
