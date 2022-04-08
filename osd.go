package nmea

const (
	// TypeOSD type for OSD sentence for Own Ship Data
	TypeOSD = "OSD"

	// OSDReferenceBottomTrackingLog is reference for bottom tracking log
	OSDReferenceBottomTrackingLog = "B"
	// OSDReferenceManual is reference for manually entered
	OSDReferenceManual = "M"
	// OSDReferenceWaterReferenced is reference for water referenced
	OSDReferenceWaterReferenced = "W"
	// OSDReferenceRadarTracking is reference for radar tracking of fixed target
	OSDReferenceRadarTracking = "R"
	// OSDReferencePositioningSystemGroundReference is reference for positioning system ground reference
	OSDReferencePositioningSystemGroundReference = "P"
)

// OSD - Own Ship Data
// https://gpsd.gitlab.io/gpsd/NMEA.html#_osd_own_ship_data
// https://github.com/nohal/OpenCPN/wiki/ARPA-targets-tracking-implementation#osd---own-ship-data
//
// Format: $--OSD,x.x,A,x.x,a,x.x,a,x.x,x.x,a*hh<CR><LF>
// Example: $RAOSD,179.0,A,179.0,M,00.0,M,,,N*76
type OSD struct {
	BaseSentence
	// Heading is Heading in degrees
	Heading float64

	// HeadingStatus is Heading status
	// * A - data valid
	// * V - data invalid
	HeadingStatus string

	// VesselTrueCourse is Vessel Course, degrees True
	VesselTrueCourse float64

	// CourseReference is Course Reference, B/M/W/R/P
	// * B - bottom tracking log
	// * M - manually entered
	// * W - water referenced
	// * R - radar tracking of fixed target
	// * P - positioning system ground reference
	CourseReference string

	// VesselSpeed is Vessel Speed
	VesselSpeed float64

	// SpeedReference is Speed Reference, B/M/W/R/P
	// * B - bottom tracking log
	// * M - manually entered
	// * W - water referenced
	// * R - radar tracking of fixed target
	// * P - positioning system ground reference.
	SpeedReference string

	// VesselSetTrue is Vessel Set, degrees True - Manually entered
	VesselSetTrue float64

	// VesselDrift is Vessel drift (speed) - Manually entered
	VesselDrift float64

	// SpeedUnits is Speed Units
	// * K - km/h
	// * N - Knots
	// * S - statute miles/h
	SpeedUnits string
}

// newOSD constructor
func newOSD(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeOSD)
	m := OSD{
		BaseSentence:     s,
		Heading:          p.Float64(0, "heading"),
		HeadingStatus:    p.EnumString(1, "heading status", StatusValid, StatusInvalid),
		VesselTrueCourse: p.Float64(2, "vessel course true"),
		CourseReference: p.EnumString(
			3,
			"course reference",
			OSDReferenceBottomTrackingLog,
			OSDReferenceManual,
			OSDReferenceWaterReferenced,
			OSDReferenceRadarTracking,
			OSDReferencePositioningSystemGroundReference,
		),
		VesselSpeed: p.Float64(4, "vessel speed"),
		SpeedReference: p.EnumString(
			5,
			"speed reference",
			OSDReferenceBottomTrackingLog,
			OSDReferenceManual,
			OSDReferenceWaterReferenced,
			OSDReferenceRadarTracking,
			OSDReferencePositioningSystemGroundReference,
		),
		VesselSetTrue: p.Float64(6, "vessel set"),
		VesselDrift:   p.Float64(7, "vessel drift"),
		SpeedUnits:    p.EnumString(8, "speed units", DistanceUnitKilometre, DistanceUnitNauticalMile, DistanceUnitStatuteMile),
	}
	return m, p.Err()
}
