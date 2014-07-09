package nmea

import (
	"testing"
)

func TestGPRMC(t *testing.T) {
	goodMsg := "$GPRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*70"
	sentence := GPRMC{}
	if err := sentence.Parse(goodMsg); err != nil {
		t.Fatalf("Parse error: %s", err)
	}

	if sentence.Time != "220516" {
		t.Errorf("Time got %s, expected %s", sentence.Time, "220516")
	}
	if sentence.Status != "A" {
		t.Errorf("Status got %s, expected %s", sentence.Status, "A")
	}
	if sentence.Speed != 173.8 {
		t.Errorf("Speed got %s, expected %s", sentence.Speed, 173.8)
	}
	if sentence.Course != 231.8 {
		t.Errorf("Course got %s, expected %s", sentence.Course, 231.8)
	}
	if sentence.Date != "130694" {
		t.Errorf("Date got %s, expected %s", sentence.Date, "130694")
	}
	if sentence.Variation != -4.2 {
		t.Errorf("Variation got %s, expected %s", sentence.Variation, -4.2)
	}

	if sentence.Latitude.PrintGPS() != "5133.8200" {
		t.Errorf("Latitude got %s, expected 5133.8200", sentence.Latitude.PrintGPS())
	}
	if sentence.Longitude.PrintGPS() != "042.2400" {
		t.Errorf("Longitude got %s, expected 042.2400", sentence.Longitude.PrintGPS())
	}

	badMsg := "$GPRMC,220516,D,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*75"
	sentence = GPRMC{}
	err := sentence.Parse(badMsg)
	if err == nil {
		t.Fatalf("parse error not returned.")
	}
	expected := "GPRMC decode, invalid status 'D'"
	if err.Error() != expected {
		t.Fatalf("parse error got %s, expected %s", err.Error(), expected)
	}

	wrongMsg := "$GPVTG,A,3,22,19,18,27,14,03,,,,,,,3.1,2.0,2.4*26"
	expErr := "GPVTG is not a GPRMC"
	if err := sentence.Parse(wrongMsg); err == nil || err.Error() != expErr {
		t.Errorf("Sentence type: Got '%v', expected error '%s'", err, expErr)
	}

	// Test generic parse.
	result, err := Parse(goodMsg)
	if err != nil {
		t.Errorf("Parse error: %v", err)
	}
	if _, ok := result.(GPRMC); !ok {
		t.Errorf("Returned type is not GPRMC")
	}

}
