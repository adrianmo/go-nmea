package nmea

const (
	// TypePKLID type for PKLID sentances
	TypePKLID = "KLID"
)


// PKLID is a Kenwood Propritary sentance used for GPS data communications in FleetSync.
// $PKLID,<0>,<1>,<2>,<3>,<4>*hh<CR><LF>
// Format:  $PKLID,xx,xxx,xxxx,xx,xx,*xx<CR><LF>
// Example: $PKLID,00,100,2000,15,00,*??
type PKLID struct {
	BaseSentence
	SentanceVersion	string	// 00 to 15
	Fleet		string	// 100 to 349
	UnitID		string	// 1000 to 4999
	Status		string	// 10 to 99
	Extension	string	// 00 to 99
}

// newPKLID constructor
func newPKLID(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypePKLID)

	return PKLID{
		BaseSentence:		s,
		SentanceVersion:	p.String(0, "sentance version, range of 00 to 15"),
		Fleet:			p.String(1, "fleet, range of 100 to 349"),
		UnitID:			p.String(2, "subscriber unit id, range of 1000 to 4999"),
		Status:			p.String(3, "subscriber unit status id, range of 10 to 99"),
		Extension:		p.String(4, "reserved for future use, range of 00 to 99"),
	}, p.Err()
}
