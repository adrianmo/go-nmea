package nmea

const (
	// TypeGSV type of GSV sentences for satellites in view
	TypeGSV = "GSV"
)

// GSV represents the GPS Satellites in view
// http://aprs.gids.nl/nmea/#glgsv
// https://gpsd.gitlab.io/gpsd/NMEA.html#_gsv_satellites_in_view
// See NMEA ERRATA # 0183 20190515 GSV Sentence:
// https://web.nmea.org/External/WCPages/WCWebContent/webcontentpage.aspx?ContentID=258
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
	SignalID        int64     // GNSS frequency (NMEA 4.1+)
	SystemID        int64     // Deprecated: Compatibility with go-nmea <= v1.10.0
}

// GSVInfo represents information about a visible satellite
type GSVInfo struct {
	SVPRNNumber int64  // SV PRN number, pseudo-random noise or gold code
	Elevation   int64  // Elevation in degrees, 90 maximum
	Azimuth     int64  // Azimuth, degrees from true north, 000 to 359
	SNR         int64  // SNR, 00-99 dB (null when not tracking)
	TalkerID    string // Context for the SignalID
	SignalID    int64  // GNSS frequency (NMEA 4.1+)
}

// newGSV constructor
func newGSV(s BaseSentence, opts ...ParserOption) (Sentence, error) {
	p := NewParser(s, opts...)
	p.AssertType(TypeGSV)
	m := GSV{
		BaseSentence:    s,
		TotalMessages:   p.Int64(0, "total number of messages"),
		MessageNumber:   p.Int64(1, "message number"),
		NumberSVsInView: p.Int64(2, "number of SVs in view"),
	}
	if (len(m.Fields)-3)%4 == 1 {
		m.SignalID = p.HexInt64(len(m.Fields)-1, "signal ID")
		m.SystemID = m.SignalID // Compatibility with go-nmea <= v1.10.0
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
			SignalID:    m.SignalID,
			TalkerID:    p.Talker,
		})
	}
	return m, p.Err()
}
