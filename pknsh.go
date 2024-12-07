package nmea

const (
	// TypePKNSH type for PKLSH sentances
	TypePKNSH = "KNSH"
)


// PKNSH is a Kenwood Propritary sentance used for GPS data communications in NEXTEDGE Digital.
//   adds UnitID and Fleet to $GPGLL sentance
// $PKNSH,<0>,<1>,<2>,<3>,<4>,<5>,<6>*hh<CR><LF>
// Format:  $PKNSH,xxxx.xxxx,x,xxxxx.xxxx,x,xxxxxx,x,Uxxxxx,*xx<CR><LF>
// Example: $PKNSH,4000.0000,N,13500.0000,E,021720,A,U00001,*??
type PKNSH struct {
	BaseSentence
	Latitude  float64       // Latitude
	Longitude float64       // Longitude
	Time      Time          // Time Stamp
	Validity  string        // validity - A=valid, V=invalid
	UnitID    string        // U00001 to U65519 or U00000001 to U16776415 (U is FIXED)
}

// newPKNSH constructor
func newPKNSH(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypePKNSH)

	return PKNSH{
		BaseSentence: s,
		Latitude:     p.LatLong(0, 1, "latitude"),
		Longitude:    p.LatLong(2, 3, "longitude"),
		Time:         p.Time(4, "time"),
		Validity:     p.EnumString(5, "validity", ValidGLL, InvalidGLL),
		UnitID:       p.String(6, "unit ID, NXDN range U00001 to U65519, DMR range of  U00000001 to U16776415"),
	}, p.Err()
}
