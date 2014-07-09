package nmea

import "testing"

func TestPSRFTXT(t *testing.T) {
	msg := "$PSRFTXT,Version:  GSWLT3.5.0MMT_3.5.00.00-CONFIG-CL31P2.00 *26"
	expected := "Version:  GSWLT3.5.0MMT_3.5.00.00-CONFIG-CL31P2.00 "
	sentence := PSRFTXT{}
	if err := sentence.Parse(msg); err != nil {
		t.Fatalf("Parse error: %s", err)
	}
	if sentence.Text != expected {
		t.Errorf("Text got '%s', want '%s'", sentence.Text, expected)
	}
	wrongMsg := "$GPVTG,A,3,22,19,18,27,14,03,,,,,,,3.1,2.0,2.4*26"
	expErr := "GPVTG is not a PSRFTXT"
	if err := sentence.Parse(wrongMsg); err == nil || err.Error() != expErr {
		t.Errorf("Sentence type: Got '%v', expected error '%s'", err, expErr)
	}
	// Test generic parse.
	result, err := Parse(msg)
	if err != nil {
		t.Errorf("Parse error: %v", err)
	}
	if _, ok := result.(PSRFTXT); !ok {
		t.Errorf("Returned type is not PSRFTXT")
	}
}
