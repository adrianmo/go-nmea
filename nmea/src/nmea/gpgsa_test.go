package nmea

import (
	"reflect"
	"testing"
)

func TestGPGSA(t *testing.T) {
	goodMsg := "$GPGSA,A,3,22,19,18,27,14,03,,,,,,,3.1,2.0,2.4*36"
	sentence := GPGSA{}
	if err := sentence.Parse(goodMsg); err != nil {
		t.Fatalf("Parse error: %s", err)
	}
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

	s := reflect.ValueOf(sentence)
	for _, a := range attrs {
		v := s.FieldByName(a.Attribute).String()
		if v != a.Value {
			t.Errorf("%s got %s, expected %s", a.Attribute, v, a.Value)
		}
	}
	// Make sure bad fix mode is detected.
	badMode := "$GPGSA,F,3,22,19,18,27,14,03,,,,,,,3.1,2.0,2.4*31"
	expErr := "Invalid selection mode [F]"
	if err := sentence.Parse(badMode); err == nil || err.Error() != expErr {
		t.Errorf("Selection mode: Got '%v', expected error '%s'", err, expErr)
	}
	// Make sure bad fix type is detected.
	badFixType := "$GPGSA,A,6,22,19,18,27,14,03,,,,,,,3.1,2.0,2.4*33"
	expErr = "Invalid fix type [6]"
	if err := sentence.Parse(badFixType); err == nil || err.Error() != expErr {
		t.Errorf("Fix type: Got '%v', expected error '%s'", err, expErr)
	}

	wrongMsg := "$GPVTG,A,3,22,19,18,27,14,03,,,,,,,3.1,2.0,2.4*26"
	expErr = "GPVTG is not a GPGSA"
	if err := sentence.Parse(wrongMsg); err == nil || err.Error() != expErr {
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
