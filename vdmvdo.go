package nmea

const (
	// TypeVDM type for VDM sentences
	TypeVDM = "VDM"

	// TypeVDO type for VDO sentences
	TypeVDO = "VDO"
)

// VDMVDO is sentence ($--VDM or $--VDO) used to encapsulate generic binary payloads. It is most commonly used with AIS data.
// https://gpsd.gitlab.io/gpsd/AIVDM.html
//
// Example: !AIVDM,1,1,,B,177KQJ5000G?tO`K>RA1wUbN0TKH,0*5C
type VDMVDO struct {
	BaseSentence
	NumFragments   int64
	FragmentNumber int64
	MessageID      int64
	Channel        string
	Payload        []byte
}

// newVDMVDO constructor
func newVDMVDO(s BaseSentence) (VDMVDO, error) {
	p := NewParser(s)
	m := VDMVDO{
		BaseSentence:   s,
		NumFragments:   p.Int64(0, "number of fragments"),
		FragmentNumber: p.Int64(1, "fragment number"),
		MessageID:      p.Int64(2, "sequence number"),
		Channel:        p.String(3, "channel ID"),
		Payload:        p.SixBitASCIIArmour(4, int(p.Int64(5, "number of padding bits")), "payload"),
	}
	return m, p.Err()
}
