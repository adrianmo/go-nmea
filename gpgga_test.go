package nmea

import (
	"reflect"
	"testing"
)

type AttrTest struct {
	Attribute string
	Value     string
}

func TestGPGGA(t *testing.T) {
	goodMsg := "$GPGGA,034225.077,3356.4650,S,15124.5567,E,1,03,9.7,-25.0,M,21.0,M,,0000*51"
	sentence, err := Parse(goodMsg)

	if err != nil {
		t.Fatalf("Parse error: %s", err)
	}

	// Attributes of the parsed sentence, and their expected values.
	attrs := []AttrTest{
		{Attribute: "Time", Value: "034225.077"},
		{Attribute: "FixQuality", Value: GPS},
		{Attribute: "NumSatellites", Value: "03"},
		{Attribute: "HDOP", Value: "9.7"},
		{Attribute: "Altitude", Value: "-25.0"},
		{Attribute: "Separation", Value: "21.0"},
		{Attribute: "DGPSAge", Value: ""},
		{Attribute: "DGPSId", Value: "0000"},
	}

	s := reflect.ValueOf(sentence)
	for _, a := range attrs {
		v := s.FieldByName(a.Attribute).String()
		if v != a.Value {
			t.Errorf("%s got %s, expected %s", a.Attribute, v, a.Value)
		}
	}

	sGPGGA := sentence.(GPGGA)
	if sGPGGA.Latitude.PrintGPS() != "3356.4650" {
		t.Errorf("Latitude got %s, expected 3356.4650", sGPGGA.Latitude.PrintGPS())
	}
	if sGPGGA.Longitude.PrintGPS() != "15124.5567" {
		t.Errorf("Longitude got %s, expected 15124.5567", sGPGGA.Longitude.PrintGPS())
	}

	// Make sure bad fix mode is detected.
	badMode := "$GPGGA,034225.077,3356.4650,S,15124.5567,E,5,03,9.7,-25.0,M,21.0,M,,0000*55"
	expErr := "Invalid fix quality [5]"
	sentence, err = Parse(badMode)
	if err == nil || err.Error() != expErr {
		t.Errorf("Selection mode: Got '%v', expected error '%s'", err, expErr)
	}

	// Test generic parse.
	result, err := Parse(goodMsg)
	if err != nil {
		t.Errorf("Parse error: %v", err)
	}
	if _, ok := result.(GPGGA); !ok {
		t.Errorf("Returned type is not GPGGA")
	}

}
