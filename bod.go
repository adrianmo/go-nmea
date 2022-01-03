package nmea

const (
	// TypeBOD type of BOD sentence for bearing waypoint to waypoint
	TypeBOD = "BOD"
)

// BOD - bearing waypoint to waypoint (origin to destination).
// Replaced by BWW in NMEA4+ (according to GPSD docs)
// If your system supports RMB it is better to use RMB as it is more common (according to OpenCPN docs)
// https://gpsd.gitlab.io/gpsd/NMEA.html#_bod_bearing_waypoint_to_waypoint
//
// Format: $--BOD,x.x,T,x.x,M,c--c,c--c*hh<CR><LF>
// Example: $GPBOD,099.3,T,105.6,M,POINTB*64
//			$GPBOD,097.0,T,103.2,M,POINTB,POINTA*4A
type BOD struct {
	BaseSentence
	BearingTrue           float64 // true bearing in degrees
	BearingTrueType       string  // is type of true bearing
	BearingMagnetic       float64 // magnetic bearing in degrees
	BearingMagneticType   string  // is type of magnetic bearing
	DestinationWaypointID string  // destination waypoint ID
	OriginWaypointID      string  // origin waypoint ID
}

// newBOD constructor
func newBOD(s BaseSentence) (BOD, error) {
	p := NewParser(s)
	p.AssertType(TypeBOD)
	bod := BOD{
		BaseSentence:          s,
		BearingTrue:           p.Float64(0, "true bearing"),
		BearingTrueType:       p.EnumString(1, "true bearing type", BearingTrue),
		BearingMagnetic:       p.Float64(2, "magnetic bearing"),
		BearingMagneticType:   p.EnumString(3, "magnetic bearing type", BearingMagnetic),
		DestinationWaypointID: p.String(4, "destination waypoint ID"),
		OriginWaypointID:      "",
	}
	// According to GSPD docs: OriginWaypointID is not transmitted in the GOTO mode, without an active route on your GPS.
	// in that case you have only DestinationWaypointID
	if len(p.Fields) > 5 {
		bod.OriginWaypointID = p.String(5, "origin waypoint ID")
	}
	return bod, p.Err()
}
