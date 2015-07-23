package nmea

import "fmt"

const (
	// PrefixGPGGA prefix
	PrefixGPGGA = "GPGGA"
	// Invalid fix quality.
	Invalid = "0"
	// GPS fix quality
	GPS = "1"
	// DGPS fix quality
	DGPS = "2"
)

// GPGGA represents fix data.
// http://aprs.gids.nl/nmea/#gga
type GPGGA struct {
	Sentence
	// Time of fix.
	Time string
	// Latitude.
	Latitude LatLong
	// Longitude.
	Longitude LatLong
	// Quality of fix.
	FixQuality string
	// Number of satellites in use.
	NumSatellites string
	// Horizontal dilution of precision.
	HDOP string
	// Altitude.
	Altitude string
	// Geoidal separation
	Separation string
	// Age of differential GPD data.
	DGPSAge string
	// DGPS reference station ID.
	DGPSId string
}

// Parse parses the GPGGA sentence into this struct.
// e.g: $GPGGA,034225.077,3356.4650,S,15124.5567,E,1,03,9.7,-25.0,M,21.0,M,,0000*58
func (g *GPGGA) parse() error {
	var err error

	if g.Type != PrefixGPGGA {
		return fmt.Errorf("%s is not a %s", g.Type, PrefixGPGGA)
	}
	g.Time = g.Fields[0]
	g.Latitude, err = NewLatLong(fmt.Sprintf("%s %s", g.Fields[1], g.Fields[2]))
	if err != nil {
		return fmt.Errorf("GPGGA decode error: %s", err)
	}
	g.Longitude, err = NewLatLong(fmt.Sprintf("%s %s", g.Fields[3], g.Fields[4]))
	if err != nil {
		return fmt.Errorf("GPGGA decode error: %s", err)
	}
	g.FixQuality = g.Fields[5]
	if g.FixQuality != Invalid && g.FixQuality != GPS && g.FixQuality != DGPS {
		return fmt.Errorf("Invalid fix quality [%s]", g.FixQuality)
	}
	g.NumSatellites = g.Fields[6]
	g.HDOP = g.Fields[7]
	g.Altitude = g.Fields[8]
	g.Separation = g.Fields[10]
	g.DGPSAge = g.Fields[12]
	g.DGPSId = g.Fields[13]
	return nil
}
