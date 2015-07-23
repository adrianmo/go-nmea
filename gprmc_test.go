package nmea

import (
	"testing"
)

func TestGPRMC(t *testing.T) {
	goodMsg := "$GPRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*70"
	s, err := Parse(goodMsg)
	if err != nil {
		t.Fatalf("Parse error: %s", err)
	}

	if s.GetSentence().Type != PrefixGPRMC {
		t.Errorf("Returned type is not GPRMC")
	}

	sentence := s.(GPRMC)

	if sentence.Time != "220516" {
		t.Errorf("Time got %s, expected %s", sentence.Time, "220516")
	}
	if sentence.Validity != "A" {
		t.Errorf("Status got %s, expected %s", sentence.Validity, "A")
	}
	if sentence.Speed != 173.8 {
		t.Errorf("Speed got %.2f, expected %.2f", sentence.Speed, 173.8)
	}
	if sentence.Course != 231.8 {
		t.Errorf("Course got %.2f, expected %.2f", sentence.Course, 231.8)
	}
	if sentence.Date != "130694" {
		t.Errorf("Date got %s, expected %s", sentence.Date, "130694")
	}
	if sentence.Variation != -4.2 {
		t.Errorf("Variation got %.2f, expected %.2f", sentence.Variation, -4.2)
	}

	if sentence.Latitude.PrintGPS() != "5133.8200" {
		t.Errorf("Latitude got %s, expected 5133.8200", sentence.Latitude.PrintGPS())
	}
	if sentence.Longitude.PrintGPS() != "042.2400" {
		t.Errorf("Longitude got %s, expected 042.2400", sentence.Longitude.PrintGPS())
	}

	badMsg := "$GPRMC,220516,D,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*75"
	s, err = Parse(badMsg)

	if err == nil {
		t.Fatalf("parse error not returned.")
	}
	expected := "GPRMC decode, invalid validity 'D'"
	if err.Error() != expected {
		t.Fatalf("parse error got %s, expected %s", err.Error(), expected)
	}

	wrongMsg := "$GPVTG,A,3,22,19,18,27,14,03,,,,,,,3.1,2.0,2.4*26"
	expErr := "Sentence type 'GPVTG' not implemented"
	s, err = Parse(wrongMsg)

	if err == nil || err.Error() != expErr {
		t.Errorf("Sentence type: Got '%v', expected error '%s'", err, expErr)
	}

	// Test generic parse.
	result, err := Parse(goodMsg)
	if err != nil {
		t.Errorf("Parse error: %v", err)
	}
	if result.GetSentence().Type != PrefixGPRMC {
		t.Errorf("Returned type is not GPRMC")
	}

}
