package nmea

const (
	// TypeEVE type of EVE sentence for General Event Message
	TypeEVE = "EVE"
)

// EVE - General Event Message
// Source: "Interfacing Voyage Data Recorder Systems, AutroSafe Interactive Fire-Alarm System, 116-P-BSL336/EE, RevA 2007-01-25,
// Autronica Fire and Security AS " (page 34 | p.8.1.5)
// https://product.autronicafire.com/fileshare/fileupload/14251/bsl336_ee.pdf
//
// Format: $FREVE,hhmmss,c--c,c--c*hh<CR><LF>
// Example: $FREVE,000001,DZ00513,Fire Alarm On: TEST DZ201 Name*0A
type EVE struct {
	BaseSentence
	Time    Time   // Event Time
	TagCode string // Tag code
	Message string // Event text
}

// newEVE constructor
func newEVE(s BaseSentence) (EVE, error) {
	p := NewParser(s)
	p.AssertType(TypeEVE)
	return EVE{
		BaseSentence: s,
		Time:         p.Time(0, "time"),
		TagCode:      p.String(1, "tag code"),
		Message:      p.String(2, "event message text"),
	}, p.Err()
}
