package nmea

const (
	// TypeVDM type for VDM sentences
	TypeVDM = "VDM"

	// TypeVDO type for VDO sentences
	TypeVDO = "VDO"
)

// VDMVDO is a format used to encapsulate generic binary payloads. It is most commonly used
// with AIS data.
// http://catb.org/gpsd/AIVDM.html
type VDMVDO struct {
	BaseSentence
	NumFragments   int64
	FragmentNumber int64
	MessageID      int64
	Channel        string
	Payload        []byte
}

// SixBitASCIIArmour decodes the 6-bit ascii armor used for VDM and VDO messages
func (p *parser) SixBitASCIIArmour(i int, fillBits int) []byte {
	if fillBits < 0 || fillBits >= 6 {
		return nil
	}

	payload := []byte(p.String(i, "encoded payload"))
	numBits := len(payload)*6 - fillBits

	if numBits < 0 {
		return nil
	}

	result := make([]byte, numBits)
	resultIndex := 0

	for _, v := range payload {
		if v < 48 || v >= 120 {
			return nil
		}

		d := v - 48
		if d > 40 {
			d -= 8
		}

		for i := 5; i >= 0 && resultIndex < len(result); i-- {
			result[resultIndex] = (d >> uint(i)) & 1
			resultIndex++
		}
	}

	return result
}

// newVDMVDO constructor
func newVDMVDO(s BaseSentence) (VDMVDO, error) {
	p := newParser(s)

	m := VDMVDO{
		BaseSentence:   s,
		NumFragments:   p.Int64(0, "number of fragments"),
		FragmentNumber: p.Int64(1, "fragment number"),
		MessageID:      p.Int64(2, "sequence number"),
		Channel:        p.String(3, "channel ID"),
		Payload:        p.SixBitASCIIArmour(4, int(p.Int64(5, "number of padding bits"))),
	}

	if m.Payload == nil {
		p.SetErr("payload", "Decode failed")
	}

	return m, p.Err()
}
