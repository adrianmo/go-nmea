package nmea

const (
	// TypeBEC type of BEC sentence for bearing and distance to waypoint (dead reckoning)
	TypeBEC = "BEC"
)

// BEC - bearing and distance to waypoint (dead reckoning)
// http://www.nmea.de/nmea0183datensaetze.html#bec
// https://www.eye4software.com/hydromagic/documentation/nmea0183/
//
// Format: $--BEC,hhmmss.ss,llll.ll,a,yyyyy.yy,a,x.x,T,x.x,M,x.x,N,c--c*hh<CR><LF>
// Example: $GPBEC,220516,5130.02,N,00046.34,W,213.8,T,218.0,M,0004.6,N,EGLM*33
type BEC struct {
	BaseSentence
	Time                       Time    // UTC Time
	Latitude                   float64 // latitude of waypoint
	Longitude                  float64 // longitude of waypoint
	BearingTrue                float64 // true bearing in degrees
	BearingTrueValid           bool    // is unit of true bearing valid
	BearingMagnetic            float64 // magnetic bearing in degrees
	BearingMagneticValid       bool    // is unit of magnetic bearing valid
	DistanceNauticalMiles      float64 // distance to waypoint in nautical miles
	DistanceNauticalMilesValid bool    // is unit of distance to waypoint nautical miles valid
	DestinationWaypointID      string  // destination waypoint ID
}

// newBEC constructor
func newBEC(s BaseSentence) (BEC, error) {
	p := NewParser(s)
	p.AssertType(TypeBEC)
	return BEC{
		BaseSentence:               s,
		Time:                       p.Time(0, "time"),
		Latitude:                   p.LatLong(1, 2, "latitude"),
		Longitude:                  p.LatLong(3, 4, "longitude"),
		BearingTrue:                p.Float64(5, "true bearing"),
		BearingTrueValid:           p.EnumString(6, "true bearing unit valid", BearingTrue) == BearingTrue,
		BearingMagnetic:            p.Float64(7, "magnetic bearing"),
		BearingMagneticValid:       p.EnumString(8, "magnetic bearing unit valid", BearingMagnetic) == BearingMagnetic,
		DistanceNauticalMiles:      p.Float64(9, "distance to waypoint is nautical miles"),
		DistanceNauticalMilesValid: p.EnumString(10, "is distance to waypoint nautical miles valid", DistanceUnitNauticalMile) == DistanceUnitNauticalMile,
		DestinationWaypointID:      p.String(11, "destination waypoint ID"),
	}, p.Err()
}
