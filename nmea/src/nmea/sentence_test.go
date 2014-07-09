package nmea

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestSentence(t *testing.T) {
	raw := "$GPFOO,1,2,3.3,x,y,zz,*51"
	rawGoodSum := strings.Split(raw, string(checksumSep))[1]
	rawBadSum := "2C"
	s := Sentence{}

	// Test checksum calculation works.
	s.Raw = raw
	s.Checksum = rawGoodSum
	if err := s.sumOk(); err != nil {
		t.Errorf("checksum failure: %s", err)
	}
	// Test that a bad checksum is flagged.
	s.Checksum = rawBadSum
	expected := "[51 != 2C]"
	if err := s.sumOk(); err.Error() != expected {
		t.Errorf("sumOk(): Got error %s, expected %s", err.Error(), expected)
	}
	// And when parsing.
	rawSplit := strings.Split(raw, checksumSep)
	badRaw := fmt.Sprintf("%s*33", rawSplit[0])
	expected = "Sentence checksum mismatch [51 != 33]"
	if err := s.Parse(badRaw); err.Error() != expected {
		t.Errorf("Parse(): Got error %s, expected %s", err.Error(), expected)
	}

	// Check that a bad start character is flagged.
	rawBadStart := "%GPFOO,1,2,3,x,y,z*1A"
	expected = "Sentence does not start with a '$'"
	if err := s.Parse(rawBadStart); err.Error() != expected {
		t.Errorf("Parse(): Got error %s, expected %s", err.Error(), expected)
	}

	// Check that a bad checksum delimiter is flagged.
	rawBadSumSep := "$GPFOO,1,2,3,x,y,z"
	expected = "Sentence does not contain single checksum separator"
	if err := s.Parse(rawBadSumSep); err.Error() != expected {
		t.Errorf("Parse(): Got error %s, expected %s", err.Error(), expected)
	}

	// Check for good parsing.
	if err := s.Parse(raw); err != nil {
		t.Errorf("Parse error: %s", err)
	}

	expectedFields := []string{"1", "2", "3.3", "x", "y", "zz", ""}
	if !reflect.DeepEqual(expectedFields, s.Fields) {
		t.Errorf("s.Fields: Got %q, expected %q", s.Fields, expectedFields)
	}
	if s.SType != "GPFOO" {
		t.Errorf("t.SType: Got %s, expected GPFOO", s.SType)
	}
	if s.Raw != raw {
		t.Errorf("t.Raw: Got %s, expected %s", s.Raw, raw)
	}
	if s.Checksum != rawGoodSum {
		t.Errorf("t.Checksum: Got %s, expected %s", s.Checksum, rawGoodSum)
	}

	// Ensure Parse works with a generic sentence.
	result, err := Parse("$GPFOO,1,2,3.3,x,y,zz,*51")
	if err != nil {
		t.Errorf("Parse return error: %v", err)
	}
	r, ok := result.(Sentence)
	if !ok {
		t.Errorf("Returned type not a Sentence.")
	}
	if r.SType != "GPFOO" {
		t.Errorf("Parsed was not a GPFOO")
	}
	// Ensure Parse returns errors when appropriate.
	result, err = Parse("$GPFOO,1,2,3.3,x,y,zz,*52")
	if result != nil || err == nil {
		t.Errorf("Result want nil, got %v. Err want error, got %v.", result, err)
	}
	expected = "Sentence checksum mismatch [51 != 52]"
	if err.Error() != expected {
		t.Errorf("Error expected %v, got %v", expected, err)
	}
}
