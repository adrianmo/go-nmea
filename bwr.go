package nmea

const (
	// TypeBWR type of BWR sentence for bearing and distance to waypoint (Rhumb Line)
	TypeBWR = "BWR"
)

// BWR - bearing and distance to waypoint (Rhumb Line). This is calculated along rumb line instead of along the great circle.
// https://gpsd.gitlab.io/gpsd/NMEA.html#_bwr_bearing_and_distance_to_waypoint_rhumb_line
//
// Format:             $--BWR,hhmmss.ss,llll.ll,a,yyyyy.yy,a,x.x,T,x.x,M,x.x,N,c--c*hh<CR><LF>
// Format (NMEA 2.3+): $--BWR,hhmmss.ss,llll.ll,a,yyyyy.yy,a,x.x,T,x.x,M,x.x,N,c--c,m*hh<CR><LF>
// Example: $GPBWR,081837,,,,,,T,,M,,N,*02
//          $GPBWR,220516,5130.02,N,00046.34,W,213.8,T,218.0,M,0004.6,N,EGLM*30
type BWR struct {
	BaseSentence
	Time                      Time    // UTC Time
	Latitude                  float64 // latitude of waypoint
	Longitude                 float64 // longitude of waypoint
	BearingTrue               float64 // true bearing in degrees
	BearingTrueType           string  // is type of true bearing
	BearingMagnetic           float64 // magnetic bearing in degrees
	BearingMagneticType       string  // is type of magnetic bearing
	DistanceNauticalMiles     float64 // distance to waypoint in nautical miles
	DistanceNauticalMilesUnit string  // is unit of distance to waypoint nautical miles
	DestinationWaypointID     string  // destination waypoint ID
	FFAMode                   string  // FAA mode indicator (filled in NMEA 2.3 and later)
}

// newBWR constructor
func newBWR(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeBWR)
	bwc := BWR{
		BaseSentence:              s,
		Time:                      p.Time(0, "time"),
		Latitude:                  p.LatLong(1, 2, "latitude"),
		Longitude:                 p.LatLong(3, 4, "longitude"),
		BearingTrue:               p.Float64(5, "true bearing"),
		BearingTrueType:           p.EnumString(6, "true bearing type", BearingTrue),
		BearingMagnetic:           p.Float64(7, "magnetic bearing"),
		BearingMagneticType:       p.EnumString(8, "magnetic bearing type", BearingMagnetic),
		DistanceNauticalMiles:     p.Float64(9, "distance to waypoint is nautical miles"),
		DistanceNauticalMilesUnit: p.EnumString(10, "is distance to waypoint nautical miles unit", DistanceUnitNauticalMile),
		DestinationWaypointID:     p.String(11, "destination waypoint ID"),
	}
	if len(p.Fields) > 12 {
		bwc.FFAMode = p.String(12, "FAA mode") // not enum because some devices have proprietary "non-nmea" values
	}
	return bwc, p.Err()
}
