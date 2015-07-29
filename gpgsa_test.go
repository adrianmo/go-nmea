package nmea

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AttrTest struct {
	Attribute string
	Value     string
}

func TestGPGSA(t *testing.T) {
	goodMsg := "$GPGSA,A,3,22,19,18,27,14,03,,,,,,,3.1,2.0,2.4*36"
	s, err := Parse(goodMsg)

	assert.NoError(t, err, "Unexpected error parsing good sentence")
	assert.Equal(t, PrefixGPGSA, s.GetSentence().Type, "Prefix does not match")

	sentence := s.(GPGSA)

	expSV := []string{"22", "19", "18", "27", "14", "03"}
	if !reflect.DeepEqual(expSV, sentence.SV) {
		t.Errorf("SV got %q, expected %q", sentence.SV, expSV)
	}

	// Attributes of the parsed sentence, and their expected values.
	attrs := []AttrTest{
		{Attribute: "Mode", Value: Auto},
		{Attribute: "FixType", Value: Fix3D},
		{Attribute: "PDOP", Value: "3.1"},
		{Attribute: "HDOP", Value: "2.0"},
		{Attribute: "VDOP", Value: "2.4"},
	}

	values := reflect.ValueOf(sentence)
	for _, a := range attrs {
		v := values.FieldByName(a.Attribute).String()
		if v != a.Value {
			t.Errorf("%s got %s, expected %s", a.Attribute, v, a.Value)
		}
	}
	// Make sure bad fix mode is detected.
	badMode := "$GPGSA,F,3,22,19,18,27,14,03,,,,,,,3.1,2.0,2.4*31"
	expErr := "Invalid selection mode [F]"
	s, err = Parse(badMode)
	if err == nil || err.Error() != expErr {
		t.Errorf("Selection mode: Got '%v', expected error '%s'", err, expErr)
	}
	// Make sure bad fix type is detected.
	badFixType := "$GPGSA,A,6,22,19,18,27,14,03,,,,,,,3.1,2.0,2.4*33"
	expErr = "Invalid fix type [6]"
	s, err = Parse(badFixType)
	if err == nil || err.Error() != expErr {
		t.Errorf("Fix type: Got '%v', expected error '%s'", err, expErr)
	}

	wrongMsg := "$GPVTG,A,3,22,19,18,27,14,03,,,,,,,3.1,2.0,2.4*26"
	expErr = "Sentence type 'GPVTG' not implemented"
	s, err = Parse(wrongMsg)
	if err == nil || err.Error() != expErr {
		t.Errorf("Sentence type: Got '%v', expected error '%s'", err, expErr)
	}

	// Test generic parse.
	result, err := Parse(goodMsg)
	if err != nil {
		t.Errorf("Parse error: %v", err)
	}
	if _, ok := result.(GPGSA); !ok {
		t.Errorf("Returned type is not GPGSA")
	}
}
