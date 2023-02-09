package nmea

const (
	// TypeXTE type of XTE sentence for Cross-track error, measured
	TypeXTE = "XTE"
)

// XTE - Cross-track error, measured
// https://gpsd.gitlab.io/gpsd/NMEA.html#_xte_cross_track_error_measured
//
// Format:            $--XTE,A,A,x.x,a,N*hh<CR><LF>
// Format (NMEA 2.3): $--XTE,A,A,x.x,a,N,m*hh<CR><LF>
// Example: $GPXTE,V,V,,,N,S*43
type XTE struct {
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

	// FAA mode indicator (filled in NMEA 2.3 and later)
	FFAMode string
}

// newXTE constructor
func newXTE(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeXTE)
	xte := XTE{
		BaseSentence:             s,
		StatusGeneralWarning:     p.EnumString(0, "general warning", StatusWarningAClearORNotUsedAPB, StatusWarningASetAPB),
		StatusLockWarning:        p.EnumString(1, "lock warning", StatusWarningBSetAPB, StatusWarningBClearAPB),
		CrossTrackErrorMagnitude: p.Float64(2, "cross track error magnitude"),
		DirectionToSteer:         p.EnumString(3, "direction to steer", Left, Right),
		CrossTrackUnits:          p.EnumString(4, "cross track units", DistanceUnitKilometre, DistanceUnitNauticalMile, DistanceUnitStatuteMile, DistanceUnitMetre),
	}
	if len(p.Fields) > 5 {
		xte.FFAMode = p.String(5, "FAA mode") // not enum because some devices have proprietary "non-nmea" values
	}
	return xte, p.Err()
}
