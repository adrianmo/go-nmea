package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type AttrTest struct {
	Attribute string
	Value     string
}

func TestGPGGAGoodSentence(t *testing.T) {
	goodMsg := "$GPGGA,034225.077,3356.4650,S,15124.5567,E,1,03,9.7,-25.0,M,21.0,M,,0000*51"
	sentence, err := Parse(goodMsg)

	assert.NoError(t, err, "Unexpected error parsing good sentence")

	lat, _ := NewLatLong("3356.4650 S")
	lon, _ := NewLatLong("15124.5567 E")
	// Attributes of the parsed sentence, and their expected values.
	expected := GPGGA{
		Sentence: Sentence{
			Type:     "GPGGA",
			Fields:   []string{"034225.077", "3356.4650", "S", "15124.5567", "E", "1", "03", "9.7", "-25.0", "M", "21.0", "M", "", "0000"},
			Checksum: "51",
			Raw:      "$GPGGA,034225.077,3356.4650,S,15124.5567,E,1,03,9.7,-25.0,M,21.0,M,,0000*51",
		},
		Time:          "034225.077",
		Latitude:      lat,
		Longitude:     lon,
		FixQuality:    GPS,
		NumSatellites: "03",
		HDOP:          "9.7",
		Altitude:      "-25.0",
		Separation:    "21.0",
		DGPSAge:       "",
		DGPSId:        "0000",
	}

	assert.EqualValues(t, expected, sentence, "Sentence values do not match")
}

func TestGPGGABadFixQuality(t *testing.T) {
	// Make sure bad fix mode is detected.
	badMode := "$GPGGA,034225.077,3356.4650,S,15124.5567,E,5,03,9.7,-25.0,M,21.0,M,,0000*55"
	_, err := Parse(badMode)

	assert.Error(t, err, "Parse error not returned")
	assert.Equal(t, err.Error(), "Invalid fix quality [5]", "Error message not as expected")
}
