package nmea

const (
	// TypeACK type of ACK sentence for alert acknowledge
	TypeACK = "ACK"
)

// ACK - Acknowledge. This sentence is used to acknowledge an alarm condition reported by a device.
// http://www.nmea.de/nmea0183datensaetze.html#ack
// https://www.furuno.it/docs/INSTALLATION%20MANUALgp170_installation_manual.pdf GPS NAVIGATOR Model GP-170 (page 42)
// https://www.manualslib.com/manual/2226813/Jrc-Jln-900.html?page=239#manual (JRC JLN-900: Installation And Instruction Manual)
//
// Format: $--ACK,xxx*hh<CR><LF>
// Example: $VRACK,001*50
type ACK struct {
	BaseSentence

	// AlertIdentifier is alert identifier (001 to 99999)
	AlertIdentifier int64 // 0
}

// newACKN constructor
func newACK(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeACK)
	return ACK{
		BaseSentence:    s,
		AlertIdentifier: p.Int64(0, "alert identifier"),
	}, p.Err()
}
