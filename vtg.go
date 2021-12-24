package nmea

const (
	// TypeVTG type for VTG sentences
	TypeVTG = "VTG"
)

// VTG represents track & speed data.
// http://aprs.gids.nl/nmea/#vtg
// https://gpsd.gitlab.io/gpsd/NMEA.html#_vtg_track_made_good_and_ground_speed
//
// Format:             $--VTG,x.x,T,x.x,M,x.x,N,x.x,K*hh<CR><LF>
// Format (NMEA 2.3+): $--VTG,x.x,T,x.x,M,x.x,N,x.x,K,m*hh<CR><LF>
// Example: $GPVTG,45.5,T,67.5,M,30.45,N,56.40,K*4B
//          $GPVTG,220.86,T,,M,2.550,N,4.724,K,A*34
type VTG struct {
	BaseSentence
	TrueTrack        float64
	MagneticTrack    float64
	GroundSpeedKnots float64
	GroundSpeedKPH   float64
	FFAMode          string // FAA mode indicator (filled in NMEA 2.3 and later)
}

// newVTG parses the VTG sentence into this struct.
// e.g: $GPVTG,360.0,T,348.7,M,000.0,N,000.0,K*43
func newVTG(s BaseSentence) (VTG, error) {
	p := NewParser(s)
	p.AssertType(TypeVTG)
	vtg := VTG{
		BaseSentence:     s,
		TrueTrack:        p.Float64(0, "true track"),
		MagneticTrack:    p.Float64(2, "magnetic track"),
		GroundSpeedKnots: p.Float64(4, "ground speed (knots)"),
		GroundSpeedKPH:   p.Float64(6, "ground speed (km/h)"),
	}
	if len(p.Fields) > 8 {
		vtg.FFAMode = p.String(8, "FAA mode") // not enum because some devices have proprietary "non-nmea" values
	}
	return vtg, p.Err()
}
