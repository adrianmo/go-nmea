package nmea

const (
	// TypeBWW type of BWW sentence for bearing (from destination) destination waypoint to origin waypoint
	TypeBWW = "BWW"
)

// BWW - bearing (from destination) destination waypoint to origin waypoint
// Replaces by BOD in NMEA4+ (according to GPSD docs)
// If your system supports RMB it is better to use RMB as it is more common (according to OpenCPN docs)
// https://gpsd.gitlab.io/gpsd/NMEA.html#_bww_bearing_waypoint_to_waypoint
// http://www.nmea.de/nmea0183datensaetze.html#bww
//
// Format: $--BWW,x.x,T,x.x,M,c--c,c--c*hh<CR><LF>
// Example: $GPBWW,097.0,T,103.2,M,POINTB,POINTA*41
type BWW struct {
	BaseSentence
	BearingTrue           float64 // true bearing in degrees
	BearingTrueType       string  // is type of true bearing
	BearingMagnetic       float64 // magnetic bearing in degrees
	BearingMagneticType   string  // is type of magnetic bearing
	DestinationWaypointID string  // destination waypoint ID
	OriginWaypointID      string  // origin waypoint ID
}

// newBWW constructor
func newBWW(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeBWW)
	bod := BWW{
		BaseSentence:          s,
		BearingTrue:           p.Float64(0, "true bearing"),
		BearingTrueType:       p.EnumString(1, "true bearing type", BearingTrue),
		BearingMagnetic:       p.Float64(2, "magnetic bearing"),
		BearingMagneticType:   p.EnumString(3, "magnetic bearing type", BearingMagnetic),
		DestinationWaypointID: p.String(4, "destination waypoint ID"),
		OriginWaypointID:      p.String(5, "origin waypoint ID"),
	}
	return bod, p.Err()
}
