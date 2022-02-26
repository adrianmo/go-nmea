package nmea

const (
	// TypeRSD type of RSD sentence for RADAR System Data
	TypeRSD = "RSD"

	// RSDDisplayRotationCourseUp is when display rotation is course up
	RSDDisplayRotationCourseUp = "C"
	// RSDDisplayRotationHeadingUp is when display rotation is ship heading up
	RSDDisplayRotationHeadingUp = "H"
	// RSDDisplayRotationNorthUp is when display rotation is (true) north up
	RSDDisplayRotationNorthUp = "N"
)

// RSD - RADAR System Data
// https://gpsd.gitlab.io/gpsd/NMEA.html#_rsd_radar_system_data
// https://github.com/nohal/OpenCPN/wiki/ARPA-targets-tracking-implementation#rsd---radar-system-data
//
// Format: $--RSD,x.x,x.x,x.x,x.x,x.x,x.x,x.x,x.x,x.x,x.x,x.x,a,a*hh<CR><LF>
// Example: $RARSD,0.00,,2.50,005.0,0.00,,4.50,355.0,,,3.0,N,H*51
// Example: $RARSD,,,,,,,,,0.808,326.9,0.750,N,N*58
// Example: $RARSD,0.00,,0.40,,,,,,,,3.0,N,N*53
type RSD struct {
	BaseSentence
	Origin1Range         float64 // Origin 1 range
	Origin1Bearing       float64 // Origin 1 bearing (degrees from 0°)
	VariableRangeMarker1 float64 // Variable Range Marker 1
	BearingLine1         float64 // Bearing Line 1

	Origin2Range         float64 // Origin 2 range
	Origin2Bearing       float64 // Origin 2 bearing (degrees from 0°)
	VariableRangeMarker2 float64 // Variable Range Marker 2
	BearingLine2         float64 // Bearing Line 2

	CursorRangeFromOwnShip float64 // Cursor Range From Own Ship
	CursorBearingDegrees   float64 // Cursor Bearing (degrees clockwise from 0°)

	RangeScale      float64 // Range scale
	RangeUnit       string  // Range units (K = kilometers, N = nautical miles, S = statute miles)
	DisplayRotation string  // Display rotation (C = course up, H = heading up, N - North up)
}

// newRSD constructor
func newRSD(s BaseSentence) (RSD, error) {
	p := NewParser(s)
	p.AssertType(TypeRSD)
	return RSD{
		BaseSentence:         s,
		Origin1Range:         p.Float64(0, "origin 1 range"),
		Origin1Bearing:       p.Float64(1, "origin 1 bearing"),
		VariableRangeMarker1: p.Float64(2, "variable range marker 1"),
		BearingLine1:         p.Float64(3, "bearing line 1"),

		Origin2Range:         p.Float64(4, "origin 2 range"),
		Origin2Bearing:       p.Float64(5, "origin 2 bearing"),
		VariableRangeMarker2: p.Float64(6, "variable range marker 2"),
		BearingLine2:         p.Float64(7, "bearing line 2"),

		CursorRangeFromOwnShip: p.Float64(8, "cursor range from own ship"),
		CursorBearingDegrees:   p.Float64(9, "cursor bearing"),

		RangeScale:      p.Float64(10, "range scale"),
		RangeUnit:       p.EnumString(11, "range units", DistanceUnitKilometre, DistanceUnitNauticalMile, DistanceUnitStatuteMile),
		DisplayRotation: p.EnumString(12, "display rotation", RSDDisplayRotationCourseUp, RSDDisplayRotationHeadingUp, RSDDisplayRotationNorthUp),
	}, p.Err()
}
