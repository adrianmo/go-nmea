package nmea

const (
	// TypeRMB type of RMB sentence for recommended minimum navigation information
	TypeRMB = "RMB"

	// DataStatusWarningClearRMB means data is OK
	DataStatusWarningClearRMB = "A"
	// DataStatusWarningSetRMB means warning flag set
	DataStatusWarningSetRMB = "V"
)

// RMB - Recommended Minimum Navigation Information. To be sent by a navigation receiver when a destination waypoint
// is active. Alternative to BOD and BWW sentences.
// https://gpsd.gitlab.io/gpsd/NMEA.html#_rmb_recommended_minimum_navigation_information
// http://aprs.gids.nl/nmea/#rmb
//
// Format:            $--RMB,A,x.x,a,c--c,c--c,llll.ll,a,yyyyy.yy,a,x.x,x.x,x.x,A*hh<CR><LF>
// Format (NMEA2.3+): $--RMB,A,x.x,a,c--c,c--c,llll.ll,a,yyyyy.yy,a,x.x,x.x,x.x,A,m*hh<CR><LF>
// Example: $GPRMB,A,0.66,L,003,004,4917.24,N,12309.57,W,001.3,052.5,000.5,V*0B
type RMB struct {
	BaseSentence

	// DataStatus is status of data,
	// * A = OK
	// * V = Navigation receiver warning
	DataStatus string

	// Cross Track error (nautical miles, 9.9 max)
	CrossTrackErrorNauticalMiles float64

	// DirectionToSteer is Direction to steer,
	//  * L = left
	//  * R = right
	DirectionToSteer string

	// OriginWaypointID is origin (FROM) waypoint ID
	OriginWaypointID string

	// DestinationWaypointID is destination (TO) waypoint ID
	DestinationWaypointID string

	// DestinationLatitude is destination waypoint latitude
	DestinationLatitude float64

	// DestinationLongitude is destination waypoint longitude
	DestinationLongitude float64

	// RangeToDestinationNauticalMiles is range to destination, nautical miles (999,9 max)
	RangeToDestinationNauticalMiles float64

	// TrueBearingToDestination is true bearing to destination, degrees
	TrueBearingToDestination float64

	// VelocityToDestinationKnots is velocity towards destination, knots
	VelocityToDestinationKnots float64

	// ArrivalStatus is Arrival Status
	// * A = arrival circle entered
	// * V = not arrived
	ArrivalStatus string

	// FAA mode indicator (filled in NMEA 2.3 and later)
	FFAMode string
}

// newRMB constructor
func newRMB(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeRMB)
	rmb := RMB{
		BaseSentence:                    s,
		DataStatus:                      p.EnumString(0, "data status", DataStatusWarningClearRMB, DataStatusWarningSetRMB),
		CrossTrackErrorNauticalMiles:    p.Float64(1, "cross track error"),
		DirectionToSteer:                p.EnumString(2, "direction to steer", Left, Right),
		OriginWaypointID:                p.String(3, "origin waypoint ID"),
		DestinationWaypointID:           p.String(4, "destination waypoint ID"),
		DestinationLatitude:             p.LatLong(5, 6, "latitude"),
		DestinationLongitude:            p.LatLong(7, 8, "latitude"),
		RangeToDestinationNauticalMiles: p.Float64(9, "range to destination"),
		TrueBearingToDestination:        p.Float64(10, "true bearing to destination"),
		VelocityToDestinationKnots:      p.Float64(11, "velocity to destination"),
		ArrivalStatus:                   p.EnumString(12, "arrival status", WPStatusArrivalCircleEnteredA, WPStatusArrivalCircleEnteredV),
		FFAMode:                         "",
	}
	if len(p.Fields) > 13 {
		rmb.FFAMode = p.String(13, "FAA mode") // not enum because some devices have proprietary "non-nmea" values
	}
	return rmb, p.Err()
}
