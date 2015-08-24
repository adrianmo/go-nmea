package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGPGLLGoodSentence(t *testing.T) {
	goodMsg := "$GPGLL,3926.7952,N,12000.5947,W,022732,A,A*58"
	s, err := Parse(goodMsg)

	assert.NoError(t, err, "Unexpected error parsing good sentence")
	assert.Equal(t, PrefixGPGLL, s.GetSentence().Type, "Prefix does not match")

	sentence := s.(GPGLL)

	assert.Equal(t, "3926.7952", sentence.Latitude.PrintGPS(), "Latitude does not match")
	assert.Equal(t, "12000.5947", sentence.Longitude.PrintGPS(), "Longitude does not match")
	assert.Equal(t, "022732", sentence.Time, "Time does not match")
	assert.Equal(t, "A", sentence.Validity, "Status does not match")
}

func TestGPGLLBadSentence(t *testing.T) {
	badMsg := "$GPGLL,3926.7952,N,12000.5947,W,022732,D,A*5D"
	_, err := Parse(badMsg)

	assert.Error(t, err, "Parse error not returned")
	assert.Equal(t, "GPGLL decode, invalid validity 'D'", err.Error(), "Incorrect error message")
}

func TestGPGLLWrongSentence(t *testing.T) {
	wrongMsg := "$GPVTG,A,3,22,19,18,27,14,03,,,,,,,3.1,2.0,2.4*26"
	_, err := Parse(wrongMsg)

	assert.Error(t, err, "Parse error not returned")
	assert.Equal(t, "Sentence type 'GPVTG' not implemented", err.Error(), "Incorrect error message")
}
