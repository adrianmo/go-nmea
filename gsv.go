package nmea

const (
	// TypeGSV type of GSV sentences for satellites in view
	TypeGSV = "GSV"
)

// GSV represents the GPS Satellites in view
// http://aprs.gids.nl/nmea/#glgsv
// https://gpsd.gitlab.io/gpsd/NMEA.html#_gsv_satellites_in_view
//
// Format:              $--GSV,x,x,x,x,x,x,x,...*hh<CR><LF>
// Format (NMEA 4.1+):  $--GSV,x,x,x,x,x,x,x,...,x*hh<CR><LF>
// Example: $GPGSV,3,1,11,09,76,148,32,05,55,242,29,17,33,054,30,14,27,314,24*71
// Example (NMEA 4.1+): $GAGSV,3,1,09,02,00,179,,04,09,321,,07,11,134,11,11,10,227,,7*7F
type GSV struct {
	BaseSentence
	TotalMessages   int64     // Total number of messages of this type in this cycle
	MessageNumber   int64     // Message number
	NumberSVsInView int64     // Total number of SVs in view
	Info            []GSVInfo // visible satellite info (0-4 of these)
	// SystemID is (GNSS) System ID (NMEA 4.1+)
	// 1 - GPS
	// 2 - GLONASS
	// 3 - Galileo
	// 4 - BeiDou
	// 5 - QZSS
	// 6 - NavID (IRNSS)
	SystemID int64
}

// GSVInfo represents information about a visible satellite
type GSVInfo struct {
	SVPRNNumber int64 // SV PRN number, pseudo-random noise or gold code
	Elevation   int64 // Elevation in degrees, 90 maximum
	Azimuth     int64 // Azimuth, degrees from true north, 000 to 359
	SNR         int64 // SNR, 00-99 dB (null when not tracking)
}

// newGSV constructor
func newGSV(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeGSV)
	m := GSV{
		BaseSentence:    s,
		TotalMessages:   p.Int64(0, "total number of messages"),
		MessageNumber:   p.Int64(1, "message number"),
		NumberSVsInView: p.Int64(2, "number of SVs in view"),
	}
	i := 0
	for ; i < 4; i++ {
		if 6+i*4 >= len(m.Fields) {
			break
		}
		m.Info = append(m.Info, GSVInfo{
			SVPRNNumber: p.Int64(3+i*4, "SV prn number"),
			Elevation:   p.Int64(4+i*4, "elevation"),
			Azimuth:     p.Int64(5+i*4, "azimuth"),
			SNR:         p.Int64(6+i*4, "SNR"),
		})
	}
	idxSID := (6 + (i-1)*4) + 1
	if len(p.Fields) == idxSID+1 {
		m.SystemID = p.Int64(idxSID, "system ID")
	}
	return m, p.Err()
}
