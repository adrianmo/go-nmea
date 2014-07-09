package nmea

import (
	"reflect"
	"testing"
)

func TestGPVTG(t *testing.T) {
	goodMsg := "$GPVTG,054.7,T,034.4,M,005.5,N,010.2,K*48"
	sentence := GPVTG{}
	if err := sentence.Parse(goodMsg); err != nil {
		t.Fatalf("Parse error: %s", err)
	}

	attrs := []AttrTest{
		{Attribute: "TrueTrack", Value: "054.7"},
		{Attribute: "MagneticTrack", Value: "034.4"},
		{Attribute: "SpeedKnots", Value: "005.5"},
		{Attribute: "SpeedKPH", Value: "010.2"},
	}

	s := reflect.ValueOf(sentence)
	for _, a := range attrs {
		v := s.FieldByName(a.Attribute).String()
		if v != a.Value {
			t.Errorf("%s got %s, expected %s", a.Attribute, v, a.Value)
		}
	}

	wrongMsg := "$GPGSA,A,3,22,19,18,27,14,03,,,,,,,3.1,2.0,2.4*36"
	expected := "GPGSA is not a GPVTG"

	if err := sentence.Parse(wrongMsg); err.Error() != expected {
		t.Errorf("Got error '%s', expected '%s'", err, expected)
	}

	badMsg := "$GPVTG,054.7,G,034.4,M,005.5,N,010.2,K*5B"
	expected = "field expected 'T' got 'G'"

	if err := sentence.Parse(badMsg); err.Error() != expected {
		t.Errorf("Got error '%s', expected '%s'", err, expected)
	}

	// Test generic parse.
	result, err := Parse(goodMsg)
	if err != nil {
		t.Errorf("Parse error: %v", err)
	}
	if _, ok := result.(GPVTG); !ok {
		t.Errorf("Returned type is not GPVTG")
	}

}
