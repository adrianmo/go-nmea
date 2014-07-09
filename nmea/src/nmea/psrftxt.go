package nmea

import "fmt"

const (
	PrefixPSRFTXT = "PSRFTXT"
)

// PSRFTXT is proprietary text.
type PSRFTXT struct {
	Sentence
	// Text in PSRFTXT message.
	Text string
}

// Parse parses the PSRFTXT message into this struct.
// e.g: $PSRFTXT,Version:  GSWLT3.5.0MMT_3.5.00.00-CONFIG-CL31P2.00 *26
func (p *PSRFTXT) Parse(s string) error {
	err := p.Sentence.Parse(s)
	if err != nil {
		// Sentence decode error.
		return err
	}
	if p.SType != PrefixPSRFTXT {
		return fmt.Errorf("%s is not a %s", p.SType, PrefixPSRFTXT)
	}
	p.Text = p.Fields[0]
	return nil
}
