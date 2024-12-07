package nmea

const (
	// TypePKNID type for PKNID sentances
	TypePKNID = "KNID"
)


// PKNID is a Kenwood Propritary sentance used for GPS data communications in NEXTEDGE Digital.
// $PKNID,<0>,<1>,<2>,<3>,*hh<CR><LF>
// Format:  $PKNID,xx,Uxxxx,xxx,xx,*xx<CR><LF>
// Example: $PKNID,00,U00065519,207,00,*??
type PKNID struct {
	BaseSentence
	SentanceVersion	string	// 00 to 15
	UnitID		string	// U00001 to U65519 or U00000001 to U16776415 (U is FIXED)
	Status		string	// 001 to 255
	Extension	string	// 00 to 99
}

// newPKNID constructor
func newPKNID(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypePKNID)

	return PKNID{
		BaseSentence:		s,
		SentanceVersion:	p.String(0, "sentance version, range of 00 to 15"),
		UnitID:			p.String(1, "unit ID, NXDN range U00001 to U65519, DMR range of  U00000001 to U16776415"),
		Status:			p.String(2, "status NXDN, range of 001 to 255"),
		Extension:		p.String(3, "reserved for future use, range of 00 to 99"),
	}, p.Err()
}
