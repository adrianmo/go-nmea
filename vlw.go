package nmea

const (
	// TypeVLW type of VLW sentence for Distance Traveled through Water
	TypeVLW = "VLW"
)

// VLW - Distance Traveled through Water
// https://gpsd.gitlab.io/gpsd/NMEA.html#_vlw_distance_traveled_through_water
//
// Format: 			 $--VLW,x.x,N,x.x,N*hh<CR><LF>
// Format (NMEA 3+): $--VLW,x.x,N,x.x,N,x.x,N,x.x,N*hh<CR><LF>
// Example: $IIVLW,10.1,N,3.2,N*7C
// Example: $IIVLW,10.1,N,3.2,N,0,N,0,N*7C
type VLW struct {
	BaseSentence
	TotalInWater           float64 // Total cumulative water distance, nm
	TotalInWaterUnit       string  // N = Nautical Miles
	SinceResetInWater      float64 // Water distance since Reset, nm
	SinceResetInWaterUnit  string  // N = Nautical Miles
	TotalOnGround          float64 // Total cumulative ground distance, nm (NMEA 3 and above)
	TotalOnGroundUnit      string  // N = Nautical Miles (NMEA 3 and above)
	SinceResetOnGround     float64 // Ground distance since reset, nm (NMEA 3 and above)
	SinceResetOnGroundUnit string  // N = Nautical Miles (NMEA 3 and above)
}

// newVLW constructor
func newVLW(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeVLW)

	vlw := VLW{
		BaseSentence:          s,
		TotalInWater:          p.Float64(0, "total cumulative water distance"),
		TotalInWaterUnit:      p.EnumString(1, "total cumulative water distance unit", DistanceUnitNauticalMile),
		SinceResetInWater:     p.Float64(2, "water distance since reset"),
		SinceResetInWaterUnit: p.EnumString(3, "water distance since reset unit", DistanceUnitNauticalMile),
	}
	if len(p.Fields) > 4 {
		vlw.TotalOnGround = p.Float64(4, "total cumulative ground distance")
		vlw.TotalOnGroundUnit = p.EnumString(5, "total cumulative ground distance unit", DistanceUnitNauticalMile)
		vlw.SinceResetOnGround = p.Float64(6, "ground distance since reset")
		vlw.SinceResetOnGroundUnit = p.EnumString(7, "ground distance since reset unit", DistanceUnitNauticalMile)
	}
	return vlw, p.Err()
}
