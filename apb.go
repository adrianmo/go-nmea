package nmea

const (
	// TypeAPB type of APB sentence for Autopilot Sentence "B"
	TypeAPB = "APB"

	// StatusWarningASetAPB indicates LORAN-C Blink or SNR warning
	StatusWarningASetAPB = "V"
	// StatusWarningAClearORNotUsedAPB general warning flag or other navigation systems when a reliable fix is not available
	StatusWarningAClearORNotUsedAPB = "A"

	// StatusWarningBSetAPB means Loran-C Cycle Lock warning OK or not used
	StatusWarningBSetAPB = "A"
	// StatusWarningBClearAPB means Loran-C Cycle Lock warning flag
	StatusWarningBClearAPB = "V"
)

// Autopilot related constants (used in APB, APA, AAM)
const (
	// WPStatusPerpendicularPassedA is warning for passing the perpendicular of the course line of waypoint
	WPStatusPerpendicularPassedA = "A"
	// WPStatusPerpendicularPassedV indicates for not passing of the perpendicular of the course line of waypoint
	WPStatusPerpendicularPassedV = "V"

	// WPStatusArrivalCircleEnteredA is warning of entering to waypoint circle
	WPStatusArrivalCircleEnteredA = "A"
	// WPStatusArrivalCircleEnteredV indicates of not yet entered into waypoint circle
	WPStatusArrivalCircleEnteredV = "V"
)

// APB - Autopilot Sentence "B" for heading/tracking
// https://gpsd.gitlab.io/gpsd/NMEA.html#_apb_autopilot_sentence_b
// https://www.tronico.fi/OH6NT/docs/NMEA0183.pdf (page 5)
//
// Format:           $--APB,A,A,x.x,a,N,A,A,x.x,a,c--c,x.x,a,x.x,a*hh<CR><LF>
// Format NMEA 2.3+: $--APB,A,A,x.x,a,N,A,A,x.x,a,c--c,x.x,a,x.x,a,a*hh<CR><LF>
// Example: $GPAPB,A,A,0.10,R,N,V,V,011,M,DEST,011,M,011,M*82
//			$ECAPB,A,A,0.0,L,M,V,V,175.2,T,Antechamber_Bay,175.2,T,175.2,T*48
type APB struct {
	BaseSentence

	// StatusGeneralWarning is used for warnings
	//  * V = LORAN-C Blink or SNR warning
	//  * A = general warning flag or other navigation systems when a reliable fix is not available
	StatusGeneralWarning string

	// StatusLockWarning is used for lock warning
	//  * V = Loran-C Cycle Lock warning flag
	//  * A = OK or not used
	StatusLockWarning string

	// CrossTrackErrorMagnitude is Cross Track Error Magnitude
	CrossTrackErrorMagnitude float64

	// DirectionToSteer is Direction to steer,
	//  * L = left
	//  * R = right
	DirectionToSteer string

	// CrossTrackUnits is cross track units
	// * N = nautical miles
	// * K = for kilometers
	CrossTrackUnits string

	// StatusArrivalCircleEntered is warning of arrival to waypoint circle
	// * A = Arrival Circle Entered
	// * V = not entered
	StatusArrivalCircleEntered string

	// StatusPerpendicularPassed is warning for perpendicular passing of waypoint
	// * A = Perpendicular passed at waypoint
	// * V = not passed
	StatusPerpendicularPassed string

	// BearingOriginToDest is Bearing origin to destination
	BearingOriginToDest float64

	// BearingOriginToDestType is Bearing origin to dest type
	// * M = Magnetic
	// * T = True
	BearingOriginToDestType string

	// DestinationWaypointID is Destination waypoint ID
	DestinationWaypointID string

	// BearingPresentToDest is Bearing, present position to Destination
	BearingPresentToDest float64

	// BearingPresentToDestType is Bearing present to dest type
	// * M = Magnetic
	// * T = True
	BearingPresentToDestType string

	// Heading is heading to steer to destination waypoint
	Heading float64

	// HeadingType is Heading type
	// * M = Magnetic
	// * T = True
	HeadingType string

	// FAA mode indicator (filled in NMEA 2.3 and later)
	FFAMode string
}

// newAPB constructor
func newAPB(s BaseSentence) (APB, error) {
	p := NewParser(s)
	p.AssertType(TypeAPB)
	apb := APB{
		BaseSentence:               s,
		StatusGeneralWarning:       p.EnumString(0, "general warning", StatusWarningAClearORNotUsedAPB, StatusWarningASetAPB),
		StatusLockWarning:          p.EnumString(1, "lock warning", StatusWarningBSetAPB, StatusWarningBClearAPB),
		CrossTrackErrorMagnitude:   p.Float64(2, "cross track error magnitude"),
		DirectionToSteer:           p.EnumString(3, "direction to steer", Left, Right),
		CrossTrackUnits:            p.EnumString(4, "cross track units", DistanceUnitKilometre, DistanceUnitNauticalMile, DistanceUnitStatuteMile, DistanceUnitMetre),
		StatusArrivalCircleEntered: p.EnumString(5, "arrival circle entered status", WPStatusArrivalCircleEnteredA, WPStatusArrivalCircleEnteredV),
		StatusPerpendicularPassed:  p.EnumString(6, "perpendicularly passed status", WPStatusPerpendicularPassedA, WPStatusPerpendicularPassedV),
		BearingOriginToDest:        p.Float64(7, "origin bearing to destination"),
		BearingOriginToDestType:    p.EnumString(8, "origin bearing to destination type", HeadingMagnetic, HeadingTrue),
		DestinationWaypointID:      p.String(9, "destination waypoint ID"),
		BearingPresentToDest:       p.Float64(10, "present bearing to destination"),
		BearingPresentToDestType:   p.EnumString(11, "present bearing to destination type", HeadingMagnetic, HeadingTrue),
		Heading:                    p.Float64(12, "heading"),
		HeadingType:                p.EnumString(13, "heading type", HeadingMagnetic, HeadingTrue),
	}
	if len(p.Fields) > 14 {
		apb.FFAMode = p.String(14, "FAA mode") // not enum because some devices have proprietary "non-nmea" values
	}
	return apb, p.Err()
}
