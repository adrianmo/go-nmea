package nmea

const (
	// TypeTLL type of TLL sentence for Target latitude and longitude
	TypeTLL = "TLL"

	// RadarTargetLost is used when target is lost
	RadarTargetLost = "L"
	// RadarTargetAcquisition is used when target is acquired
	RadarTargetAcquisition = "Q"
	// RadarTargetTracking is used when tracking target
	RadarTargetTracking = "T"
)

// TLL - Target latitude and longitude
// https://gpsd.gitlab.io/gpsd/NMEA.html#_tll_target_latitude_and_longitude
// https://github.com/nohal/OpenCPN/wiki/ARPA-targets-tracking-implementation#tll---target-latitude-and-longitude
//
// Format: $--TLL,xx,llll.ll,a,yyyyy.yy,a,c--c,hhmmss.ss,a,a*hh<CR><LF>
// Example: $RATLL,,3647.422,N,01432.592,E,,,,*58
type TLL struct {
	BaseSentence
	TargetNumber    int64   // Target number 00 – 99
	TargetLatitude  float64 // Target latitude + N/S
	TargetLongitude float64 // Target longitude + E/W
	TargetName      string  // Target name
	TimeUTC         Time    // UTC of data, hh is hours, mm is minutes, ss.ss is seconds.
	TargetStatus    string  // Target status (L=lost, Q=acquisition, T=tracking)
	ReferenceTarget string  // Reference target, R= reference target; null (,,)= otherwise
}

// newTLL constructor
func newTLL(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeTLL)
	return TLL{
		BaseSentence:    s,
		TargetNumber:    p.Int64(0, "target number"),
		TargetLatitude:  p.LatLong(1, 2, "latitude"),
		TargetLongitude: p.LatLong(3, 4, "longitude"),
		TargetName:      p.String(5, "target name"),
		TimeUTC:         p.Time(6, "UTC time"),
		TargetStatus:    p.EnumString(7, "target status", RadarTargetLost, RadarTargetAcquisition, RadarTargetTracking),
		ReferenceTarget: p.EnumString(8, "reference target", "R"),
	}, p.Err()
}
