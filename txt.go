package nmea

import "strings"

const (
	// TypeTXT type for TXT sentences for the transmission of text messages
	TypeTXT = "TXT"
)

// TXT is sentence for the transmission of short text messages, longer text messages may be transmitted by using
// multiple sentences. This sentence is intended to convey human readable textual information for display purposes.
// The TXT sentence shall not be used for sending commands and making device configuration changes.
// https://www.nmea.org/Assets/20160520%20txt%20amendment.pdf
//
// Format: $--TXT,xx,xx,xx,c-c*hh<CR><LF>
// Example: $GNTXT,01,01,02,u-blox AG - www.u-blox.com*4E
type TXT struct {
	BaseSentence
	TotalNumber int64 // total number of sentences, 01 to 99
	Number      int64 // number of current sentences, 01 to 99
	ID          int64 // identifier of the text message, 01 to 99
	// Message contains ASCII characters, and code delimiters if needed, up to the maximum permitted sentence length
	// (i.e., up to 61 characters including any code delimiters)
	Message string
}

// newTXT constructor
func newTXT(s BaseSentence) (TXT, error) {
	p := NewParser(s)
	p.AssertType(TypeTXT)
	m := TXT{
		BaseSentence: s,
		TotalNumber:  p.Int64(0, "total number of sentences"),
		Number:       p.Int64(1, "sentence number"),
		ID:           p.Int64(2, "sentence identifier"),
		Message:      strings.Join(p.Fields[3:], FieldSep),
	}
	return m, p.Err()
}
