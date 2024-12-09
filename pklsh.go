package nmea

const (
	// TypePKLSH type for PKLSH sentances
	TypePKLSH = "KLSH"
)

// PKLSH is a Kenwood Propritary sentance used for GPS data communications in FleetSync.
//
//	adds UnitID and Fleet to $GPGLL sentance
//
// $PKLSH,<0>,<1>,<2>,<3>,<4>,<5>,<6>,<7>*hh<CR><LF>
// Format:  $PKLSH,xxxx.xxxx,x,xxxxx.xxxx,x,xxxxxx,x,xxx,xxxx,*xx<CR><LF>
// Example: $PKLSH,4000.0000,N,13500.0000,E,021720,A,100,2000,*??
type PKLSH struct {
	BaseSentence
	Latitude  float64 // Latitude
	Longitude float64 // Longitude
	Time      Time    // Time Stamp
	Validity  string  // validity - A=valid, V=invalid
	Fleet     string  // 100 to 349
	UnitID    string  // 1000 to 4999
}

// newPKLSH constructor
func newPKLSH(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypePKLSH)

	return PKLSH{
		BaseSentence: s,
		Latitude:     p.LatLong(0, 1, "latitude"),
		Longitude:    p.LatLong(2, 3, "longitude"),
		Time:         p.Time(4, "time"),
		Validity:     p.EnumString(5, "validity", ValidGLL, InvalidGLL),
		Fleet:        p.String(6, "fleet, range of 100 to 349"),
		UnitID:       p.String(7, "subscriber unit id, range of 1000 to 4999"),
	}, p.Err()
}
