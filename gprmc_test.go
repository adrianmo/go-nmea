package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGPRMCGoodSentence(t *testing.T) {
	goodMsg := "$GPRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*70"
	s, err := Parse(goodMsg)

	assert.NoError(t, err, "Unexpected error parsing good sentence")
	assert.Equal(t, PrefixGPRMC, s.GetSentence().Type, "Prefix does not match")

	sentence := s.(GPRMC)

	assert.Equal(t, "220516", sentence.Time, "Time does not match")
	assert.Equal(t, "A", sentence.Validity, "Status does not match")
	assert.Equal(t, 173.8, sentence.Speed, "Speed does not match")
	assert.Equal(t, 231.8, sentence.Course, "Course does not match")
	assert.Equal(t, "130694", sentence.Date, "Date does not match")
	assert.Equal(t, -4.2, sentence.Variation, "Variation does not match")
	assert.Equal(t, "5133.8200", sentence.Latitude.PrintGPS(), "Latitude does not match")
	assert.Equal(t, "042.2400", sentence.Longitude.PrintGPS(), "Longitude does not match")
}

func TestGPRMCBadSentence(t *testing.T) {
	badMsg := "$GPRMC,220516,D,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*75"
	_, err := Parse(badMsg)

	assert.Error(t, err, "Parse error not returned")
	assert.Equal(t, "GPRMC decode, invalid validity 'D'", err.Error(), "Incorrect error message")
}

func TestGPRMCWrongSentence(t *testing.T) {
	wrongMsg := "$GPXTE,A,A,4.07,L,N*6D"
	_, err := Parse(wrongMsg)

	assert.Error(t, err, "Parse error not returned")
	assert.Equal(t, "Sentence type 'GPXTE' not implemented", err.Error(), "Incorrect error message")
}
