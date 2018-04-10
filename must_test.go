package nmea

// MustParseLatLong parses the supplied string into the LatLong.
// It panics if an error is encountered
func MustParseLatLong(s string) LatLong {
	l, err := ParseLatLong(s)
	if err != nil {
		panic(err)
	}
	return l
}

// MustParseGPS parses a GPS/NMEA coordinate or panics if it fails.
func MustParseGPS(s string) LatLong {
	l, err := ParseGPS(s)
	if err != nil {
		panic(err)
	}
	return l
}

// MustParseDMS parses a coordinate in degrees, minutes, seconds and
// panics on failure
func MustParseDMS(s string) LatLong {
	l, err := ParseDMS(s)
	if err != nil {
		panic(err)
	}
	return l
}

// ParseDecimal parses a decimal format coordinate and panics on error.
func MustParseDecimal(s string) LatLong {
	l, err := ParseDecimal(s)
	if err != nil {
		panic(err)
	}
	return l
}

// MustParseTime parses wall clock and panics on failure
func MustParseTime(s string) Time {
	t, err := ParseTime(s)
	if err != nil {
		panic(err)
	}
	return t
}

// MustParseDate parses a date and panics on failure
func MustParseDate(s string) Date {
	d, err := ParseDate(s)
	if err != nil {
		panic(err)
	}
	return d
}
