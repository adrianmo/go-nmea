package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksumOK(t *testing.T) {
	s := Sentence{Raw: "$GPFOO,1,2,3.3,x,y,zz,*51", Checksum: "51"}
	assert.NoError(t, s.sumOk(), "Checksum check failed")
}

func TestChecksumBad(t *testing.T) {
	s := Sentence{Raw: "$GPFOO,1,2,3.3,x,y,zz,*51", Checksum: "2C"}
	expected := "[51 != 2C]"
	assert.Error(t, s.sumOk(), "Expected '%s'", expected)
}

func TestChecksumBadRaw(t *testing.T) {
	badRaw := "$GPFOO,1,2,3.3,x,y,zz,*33"
	expected := "Sentence checksum mismatch [51 != 33]"
	_, err := Parse(badRaw)
	assert.Error(t, err, "Expected '%s'", expected)
}

func TestBadStartCharacter(t *testing.T) {
	// Check that a bad start character is flagged.
	rawBadStart := "%GPFOO,1,2,3,x,y,z*1A"
	expected := "Sentence does not contain a '$'"
	_, err := Parse(rawBadStart)
	assert.Error(t, err, "Expected '%s'", expected)
}

func TestBadChecksumDelimiter(t *testing.T) {
	// Check that a bad checksum delimiter is flagged.
	rawBadSumSep := "$GPFOO,1,2,3,x,y,z"
	expected := "Sentence does not contain single checksum separator"
	_, err := Parse(rawBadSumSep)
	assert.Error(t, err, "Expected '%s'", expected)
}

func TestGoodParsing(t *testing.T) {
	// Check for good parsing.
	raw := "$GPRMC,235236,A,3925.9479,N,11945.9211,W,44.7,153.6,250905,15.2,E,A*0C"
	_, err := Parse(raw)
	assert.NoError(t, err, "Parse error")
}

func TestGoodFields(t *testing.T) {
	raw := "$GPRMC,235236,A,3925.9479,N,11945.9211,W,44.7,153.6,250905,15.2,E,A*0C"
	expectedFields := []string{"235236", "A", "3925.9479", "N", "11945.9211", "W", "44.7", "153.6", "250905", "15.2", "E", "A"}
	m, _ := Parse(raw)
	assert.EqualValues(t, expectedFields, m.GetSentence().Fields, "Got '%q', expected '%q'", m.GetSentence().Fields, expectedFields)
}

func TestGoodSentenceType(t *testing.T) {
	raw := "$GPRMC,235236,A,3925.9479,N,11945.9211,W,44.7,153.6,250905,15.2,E,A*0C"
	expected := "GPRMC"
	m, _ := Parse(raw)
	assert.Equal(t, expected, m.GetSentence().Type, "Got '%s', expected '%s'", m.GetSentence().Type, expected)
}

func TestGoodRawSentence(t *testing.T) {
	raw := "$GPRMC,235236,A,3925.9479,N,11945.9211,W,44.7,153.6,250905,15.2,E,A*0C"
	m, _ := Parse(raw)
	assert.Equal(t, raw, m.GetSentence().Raw, "Bad raw sentence")
}

func TestReturnValues(t *testing.T) {
	// Ensure Parse returns errors when appropriate.
	result, err := Parse("$GPRMC,235236,A,3925.9479,N,11945.9211,W,44.7,153.6,250905,15.2,E,A*0A")
	assert.Nil(t, result, "Result should be nil")
	assert.NotNil(t, err, "Err should be an error")
	assert.Equal(t, "Sentence checksum mismatch [0C != 0A]", err.Error(), "Error sentence mismatch")
}
