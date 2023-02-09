package nmea

// Latitude / longitude representation.

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const (
	// StatusValid indicated status having valid value
	StatusValid = "A"
	// StatusInvalid indicated status having invalid value
	StatusInvalid = "V"
)

const (
	// UnitAmpere is unit for current in Amperes
	UnitAmpere = "A"
	// UnitBars is unit for pressure in Bars
	UnitBars = "B"
	// UnitBinary is unit for binary data
	UnitBinary = "B"
	// UnitCelsius is unit for temperature in Celsius
	UnitCelsius = TemperatureCelsius
	// UnitFahrenheit is unit for temperature in Fahrenheit
	UnitFahrenheit = TemperatureFahrenheit
	// UnitDegrees is unit for angular displacement in Degrees
	UnitDegrees = "D"
	// UnitHertz is unit for frequency in Hertz
	UnitHertz = "H"
	// UnitLitresPerSecond is unit for volumetric flow in Litres per second
	UnitLitresPerSecond = "I"
	// UnitKelvin is unit of temperature in Kelvin
	UnitKelvin = TemperatureKelvin
	// UnitKilogramPerCubicMetre is unit of density in kilogram per cubic metre
	UnitKilogramPerCubicMetre = "K"
	// UnitMeters is unit of distance in Meters
	UnitMeters = DistanceUnitMetre
	// UnitNewtons is unit of force in Newtons (1 kg*m/s2)
	UnitNewtons = "N"
	// UnitCubicMeters is unit of volume in cubic meters
	UnitCubicMeters = "M"
	// UnitRevolutionsPerMinute is unit of rotational speed or the frequency of rotation around a fixed axis in revolutions per minute (RPM)
	UnitRevolutionsPerMinute = "R"
	// UnitPercent is percent of full range
	UnitPercent = "P"
	// UnitPascal is unit of pressure in Pascals
	UnitPascal = "P"
	// UnitPartsPerThousand is in parts-per notation set of pseudo-unit to describe small values of miscellaneous dimensionless quantities, e.g. mole fraction or mass fraction.
	UnitPartsPerThousand = "S"
	// UnitVolts is unit of voltage in Volts
	UnitVolts = "V"
)

const (
	// SpeedKnots is a unit of speed equal to one nautical mile per hour, exactly 1.852 km/h (approximately 1.151 mph or 0.514 m/s)
	SpeedKnots = "N"
	// SpeedMeterPerSecond is unit of speed of 1 meter per second
	SpeedMeterPerSecond = "M"
	// SpeedKilometerPerHour is unit of speed of 1 kilometer per hour
	SpeedKilometerPerHour = "K"
)

const (
	// TemperatureCelsius is unit of temperature measured in celsius. °C = (°F − 32) / 1,8
	TemperatureCelsius = "C"
	// TemperatureFahrenheit is unit of temperature measured in fahrenheits. °F = °C * 1,8 + 32
	TemperatureFahrenheit = "F"
	// TemperatureKelvin is unit of temperature measured in kelvins. K = °C + 273,15
	TemperatureKelvin = "K"
)

// In navigation, the heading of a vessel or object is the compass direction in which the craft's bow or nose is pointed.
// Note that the heading may not necessarily be the direction that the vehicle actually travels, which is known as
// its course or track.
// https://en.wikipedia.org/wiki/Heading_(navigation)
const (
	// HeadingMagnetic - Magnetic heading is your direction relative to magnetic north, read from your magnetic compass.
	// Magnetic north is the point on the Earth's surface where its magnetic field points directly downwards.
	HeadingMagnetic = "M"
	// HeadingTrue - True heading is your direction relative to true north, or the geographic north pole.
	// True north is the northern axis of rotation of the Earth. It is the point where the lines of longitude converge
	// on maps.
	HeadingTrue = "T"
)

// In nautical navigation the absolute bearing is the clockwise angle between north and an object observed from the vessel.
// https://en.wikipedia.org/wiki/Bearing_(angle)
const (
	// BearingMagnetic is the clockwise angle between Earth's magnetic north and an object observed from the vessel.
	BearingMagnetic = "M"
	// BearingTrue  is the clockwise angle between Earth's true (geographical) north and an object observed from the vessel.
	BearingTrue = "T"
)

// FAAMode is type for FAA mode indicator (NMEA 2.3 and later).
// In NMEA 2.3, several sentences (APB, BWC, BWR, GLL, RMA, RMB, RMC, VTG, WCV, and XTE) got a new last field carrying
// the signal integrity information needed by the FAA.
// Source: https://www.xj3.nl/dokuwiki/doku.php?id=nmea
// Note: there can be other values (proprietary).
const (
	// FAAModeAutonomous is Autonomous mode
	FAAModeAutonomous = "A"
	// FAAModeDifferential is Differential Mode
	FAAModeDifferential = "D"
	// FAAModeEstimated is Estimated (dead-reckoning) mode
	FAAModeEstimated = "E"
	// FAAModeRTKFloat is RTK Float mode
	FAAModeRTKFloat = "F"
	// FAAModeManualInput is Manual Input Mode
	FAAModeManualInput = "M"
	// FAAModeDataNotValid is Data Not Valid
	FAAModeDataNotValid = "N"
	// FAAModePrecise is Precise (NMEA4.00+)
	FAAModePrecise = "P"
	// FAAModeRTKInteger is RTK Integer mode
	FAAModeRTKInteger = "R"
	// FAAModeSimulated is Simulated Mode
	FAAModeSimulated = "S"
)

// Navigation Status (NMEA 4.1 and later)
const (
	// NavStatusAutonomous is Autonomous mode
	NavStatusAutonomous = "A"
	// NavStatusDifferential is Differential Mode
	NavStatusDifferential = "D"
	// NavStatusEstimated is Estimated (dead-reckoning) mode
	NavStatusEstimated = "E"
	// NavStatusManualInput is Manual Input Mode
	NavStatusManualInput = "M"
	// NavStatusSimulated is Simulated Mode
	NavStatusSimulated = "S"
	// NavStatusDataNotValid is Data Not Valid
	NavStatusDataNotValid = "N"
	// NavStatusDataValid is valid
	NavStatusDataValid = "V"
)

const (
	// DistanceUnitKilometre is unit for distance in kilometres (1km = 1000m)
	DistanceUnitKilometre = "K"
	// DistanceUnitNauticalMile is unit for distance in nautical miles (1nmi = 1852m)
	DistanceUnitNauticalMile = "N"
	// DistanceUnitStatuteMile is unit for distance in statute miles (1smi = 5,280 feet = 1609.344m)
	DistanceUnitStatuteMile = "S"
	// DistanceUnitMetre is unit for distance in metres
	DistanceUnitMetre = "M"
	// DistanceUnitFeet is unit for distance in feets (1f = 0.3048m)
	DistanceUnitFeet = "f"
	// DistanceUnitFathom is unit for distance in fathoms (1fm = 6ft = 1,8288m)
	DistanceUnitFathom = "F"
)

const (
	// Degrees value
	Degrees = '\u00B0'
	// Minutes value
	Minutes = '\''
	// Seconds value
	Seconds = '"'
	// Point value
	Point = '.'
	// North value
	North = "N"
	// South value
	South = "S"
	// East value
	East = "E"
	// West value
	West = "W"
	// Left value
	Left = "L"
	// Right value
	Right = "R"
)

// ParseLatLong parses the supplied string into the LatLong.
//
// Supported formats are:
// - DMS (e.g. 33° 23' 22")
// - Decimal (e.g. 33.23454)
// - GPS (e.g 15113.4322 S)
func ParseLatLong(s string) (float64, error) {
	var l float64
	if v, err := ParseDMS(s); err == nil {
		l = v
	} else if v, err := ParseGPS(s); err == nil {
		l = v
	} else if v, err := ParseDecimal(s); err == nil {
		l = v
	} else {
		return 0, fmt.Errorf("cannot parse [%s], unknown format", s)
	}
	return l, nil
}

// ParseGPS parses a GPS/NMEA coordinate.
// e.g `15113.4322 S`
func ParseGPS(s string) (float64, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid format: %s", s)
	}
	dir := parts[1]
	value, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("parse error: %s", err.Error())
	}

	degrees := math.Floor(value / 100)
	minutes := value - (degrees * 100)
	value = degrees + minutes/60

	if dir == North || dir == East {
		return value, nil
	} else if dir == South || dir == West {
		return 0 - value, nil
	} else {
		return 0, fmt.Errorf("invalid direction [%s]", dir)
	}
}

// FormatGPS formats a GPS/NMEA coordinate
func FormatGPS(l float64) string {
	padding := ""
	degrees := math.Floor(math.Abs(l))
	fraction := (math.Abs(l) - degrees) * 60
	if fraction < 10 {
		padding = "0"
	}
	return fmt.Sprintf("%d%s%.4f", int(degrees), padding, fraction)
}

// ParseDecimal parses a decimal format coordinate.
// e.g: 151.196019
func ParseDecimal(s string) (float64, error) {
	// Make sure it parses as a float.
	l, err := strconv.ParseFloat(s, 64)
	if err != nil || s[0] != '-' && len(strings.Split(s, ".")[0]) > 3 {
		return 0.0, errors.New("parse error (not decimal coordinate)")
	}
	return l, nil
}

// ParseDMS parses a coordinate in degrees, minutes, seconds.
// - e.g. 33° 23' 22"
func ParseDMS(s string) (float64, error) {
	degrees := 0
	minutes := 0
	seconds := 0.0
	// Whether a number has finished parsing (i.e whitespace after it)
	endNumber := false
	// Temporary parse buffer.
	tmpBytes := []byte{}
	var err error

	for i, r := range s {
		switch {
		case unicode.IsNumber(r) || r == '.':
			if !endNumber {
				tmpBytes = append(tmpBytes, s[i])
			} else {
				return 0, errors.New("parse error (no delimiter)")
			}
		case unicode.IsSpace(r) && len(tmpBytes) > 0:
			endNumber = true
		case r == Degrees:
			if degrees, err = strconv.Atoi(string(tmpBytes)); err != nil {
				return 0, errors.New("parse error (degrees)")
			}
			tmpBytes = tmpBytes[:0]
			endNumber = false
		case s[i] == Minutes:
			if minutes, err = strconv.Atoi(string(tmpBytes)); err != nil {
				return 0, errors.New("parse error (minutes)")
			}
			tmpBytes = tmpBytes[:0]
			endNumber = false
		case s[i] == Seconds:
			if seconds, err = strconv.ParseFloat(string(tmpBytes), 64); err != nil {
				return 0, errors.New("parse error (seconds)")
			}
			tmpBytes = tmpBytes[:0]
			endNumber = false
		case unicode.IsSpace(r) && len(tmpBytes) == 0:
			continue
		default:
			return 0, fmt.Errorf("parse error (unknown symbol [%d])", s[i])
		}
	}
	if len(tmpBytes) > 0 {
		return 0, fmt.Errorf("parse error (trailing data [%s])", string(tmpBytes))
	}
	val := float64(degrees) + (float64(minutes) / 60.0) + (float64(seconds) / 60.0 / 60.0)
	return val, nil
}

// FormatDMS returns the degrees, minutes, seconds format for the given LatLong.
func FormatDMS(l float64) string {
	val := math.Abs(l)
	degrees := int(math.Floor(val))
	minutes := int(math.Floor(60 * (val - float64(degrees))))
	seconds := 3600 * (val - float64(degrees) - (float64(minutes) / 60))
	return fmt.Sprintf("%d\u00B0 %d' %f\"", degrees, minutes, seconds)
}

// Time type
type Time struct {
	Valid       bool
	Hour        int
	Minute      int
	Second      int
	Millisecond int
}

// String representation of Time
func (t Time) String() string {
	seconds := float64(t.Second) + float64(t.Millisecond)/1000
	return fmt.Sprintf("%02d:%02d:%07.4f", t.Hour, t.Minute, seconds)
}

// timeRe is used to validate time strings
var timeRe = regexp.MustCompile(`^\d{6}(\.\d*)?$`)

// ParseTime parses wall clock time.
// e.g. hhmmss.ssss
// An empty time string will result in an invalid time.
func ParseTime(s string) (Time, error) {
	if s == "" {
		return Time{}, nil
	}
	if !timeRe.MatchString(s) {
		return Time{}, fmt.Errorf("parse time: expected hhmmss.ss format, got '%s'", s)
	}
	hour, _ := strconv.Atoi(s[:2])
	minute, _ := strconv.Atoi(s[2:4])
	second, _ := strconv.ParseFloat(s[4:], 64)
	whole, frac := math.Modf(second)
	return Time{true, hour, minute, int(whole), int(round(frac * 1000))}, nil
}

// round is implemented here because it wasn't added until go1.10
// this code is taken directly from the math.Round documentation
// TODO: use math.Round after a reasonable amount of time
func round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}

// Date type
type Date struct {
	Valid bool
	DD    int
	MM    int
	YY    int
}

// String representation of date
func (d Date) String() string {
	return fmt.Sprintf("%02d/%02d/%02d", d.DD, d.MM, d.YY)
}

// ParseDate field ddmmyy format
func ParseDate(ddmmyy string) (Date, error) {
	if ddmmyy == "" {
		return Date{}, nil
	}
	if len(ddmmyy) != 6 {
		return Date{}, fmt.Errorf("parse date: exptected ddmmyy format, got '%s'", ddmmyy)
	}
	dd, err := strconv.Atoi(ddmmyy[0:2])
	if err != nil {
		return Date{}, errors.New(ddmmyy)
	}
	mm, err := strconv.Atoi(ddmmyy[2:4])
	if err != nil {
		return Date{}, errors.New(ddmmyy)
	}
	yy, err := strconv.Atoi(ddmmyy[4:6])
	if err != nil {
		return Date{}, errors.New(ddmmyy)
	}
	return Date{true, dd, mm, yy}, nil
}

// LatDir returns the latitude direction symbol
func LatDir(l float64) string {
	if l < 0.0 {
		return South
	}
	return North
}

// LonDir returns the longitude direction symbol
func LonDir(l float64) string {
	if l < 0.0 {
		return East
	}
	return West
}

// Float64 is a nullable float64 value
type Float64 struct {
	Value float64
	Valid bool
}

// Int64 is a nullable int64 value
type Int64 struct {
	Value int64
	Valid bool
}
