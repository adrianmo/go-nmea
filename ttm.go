package nmea

const (
	// TypeTTM type of TTM sentence for Tracked Target Message
	TypeTTM = "TTM"
)

// TTM - Tracked Target Message
// https://gpsd.gitlab.io/gpsd/NMEA.html#_ttm_tracked_target_message
// https://github.com/nohal/OpenCPN/wiki/ARPA-targets-tracking-implementation#ttm---tracked-target-message
//
// Format: $--TTM,xx,x.x,x.x,a,x.x,x.x,a,x.x,x.x,a,c--c,a,a*hh<CR><LF>
// Format: $--TTM,xx,x.x,x.x,a,x.x,x.x,a,x.x,x.x,a,c--c,a,a,hhmmss.ss,a*hh<CR><LF>
// Example: $RATTM,02,1.43,170.5,T,0.16,264.4,T,1.42,36.9,N,,T,,,M*2A
type TTM struct {
	BaseSentence
	TargetNumber      int64   // Target number 00 – 99
	TargetDistance    float64 // Target Distance
	Bearing           float64 // Bearing from own ship, degrees
	BearingType       string  // Type of target Bearing, T = True, R = Relative
	TargetSpeed       float64 // Target Speed
	TargetCourse      float64 // Target Course
	CourseType        string  // target course type,  T = True, R = Relative
	DistanceCPA       float64 // Distance of closest-point-of-approach
	TimeCPA           float64 // Time until closest-point-of-approach "-" means increasing
	SpeedUnits        string  // Speed/distance units, K/N/S
	TargetName        string  // Target name
	TargetStatus      string  // Target status (L=lost, Q=acquisition, T=tracking)
	ReferenceTarget   string  // Reference target, R= reference target; null (,,)= otherwise
	TimeUTC           Time    // UTC of data, hh is hours, mm is minutes, ss.ss is seconds.
	TypeOfAcquisition string  // Type, A = Auto, M = Manual, R = Reported
}

// newTTM constructor
func newTTM(s BaseSentence) (TTM, error) {
	p := NewParser(s)
	p.AssertType(TypeTTM)
	return TTM{
		BaseSentence:      s,
		TargetNumber:      p.Int64(0, "target number"),
		TargetDistance:    p.Float64(1, "target Distance"),
		Bearing:           p.Float64(2, "bearing"),
		BearingType:       p.EnumString(3, "bearing type", "T", "R"),
		TargetSpeed:       p.Float64(4, "target speed"),
		TargetCourse:      p.Float64(5, "target course"),
		CourseType:        p.EnumString(6, "course type", "T", "R"),
		DistanceCPA:       p.Float64(7, "distance CPA"),
		TimeCPA:           p.Float64(8, "time of CPA"),
		SpeedUnits:        p.EnumString(9, "speed units", DistanceUnitKilometre, DistanceUnitNauticalMile, DistanceUnitStatuteMile),
		TargetName:        p.String(10, "target name"),
		TargetStatus:      p.EnumString(11, "target status", RadarTargetLost, RadarTargetAcquisition, RadarTargetTracking),
		ReferenceTarget:   p.EnumString(12, "reference target", "R"),
		TimeUTC:           p.Time(13, "UTC time"),
		TypeOfAcquisition: p.EnumString(14, "type of acquisition", "A", "M", "R"),
	}, p.Err()
}
