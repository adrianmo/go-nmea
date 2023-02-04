package nmea

import "errors"

const (
	// TypeALC type of ALC sentence for cyclic alert list
	TypeALC = "ALC"
)

// ALC - Cyclic alert list
// https://fcc.report/FCC-ID/ADB9ZWRTR100/2768717.pdf (page 6) FURUNO MARINE RADAR, model FAR-15XX manual
//
// Format: $--ALC,xx,xx,xx,xx, aaa,x.x,x.x,x.x,’’’’’’’’’,*hh<CR><LF>
// Example: $FBALC,02,01,03,01,FEB,01,02,03*0A
type ALC struct {
	BaseSentence

	// NumFragments is total number of ALC sentences this message (01, 16)
	NumFragments int64 // 0

	//  FragmentNumber is current fragment/sentence number (01 - 16)
	FragmentNumber int64 // 1

	// MessageID is sequential message identifier (00 - 99)
	MessageID int64 // 2

	// Number of alert entries (0 - 3)
	EntriesNumber int64 // 3

	// Additional alert entries. Each entry identifies a certain alert with a certain state.
	// It is not allowed that an alert entry is split between two ALC sentences
	AlertEntries []ALCAlertEntry // 4
}

// ALCAlertEntry is instance of alert entry for ALC sentence
type ALCAlertEntry struct {
	// ManufacturerMnemonicCode is manufacturer mnemonic code
	ManufacturerMnemonicCode string // i+4

	// AlertIdentifier is alert identifier (001 to 99999)
	AlertIdentifier int64 // i+5

	// AlertInstance is alert instance
	AlertInstance int64 // i+6

	// RevisionCounter is revision counter (1 - 99)
	RevisionCounter int64 // i+7
}

// newALC constructor
func newALC(s BaseSentence) (ALC, error) {
	p := NewParser(s)
	p.AssertType(TypeALC)
	alc := ALC{
		BaseSentence:   s,
		NumFragments:   p.Int64(0, "number of fragments"),
		FragmentNumber: p.Int64(1, "fragment number"),
		MessageID:      p.Int64(2, "message ID"),
		EntriesNumber:  p.Int64(3, "entries number"),
		AlertEntries:   nil,
	}

	fieldCount := len(p.Fields)
	if fieldCount == 4 {
		return alc, p.Err()
	}
	if fieldCount%4 != 0 {
		return alc, errors.New("ALC data set field count is not exactly dividable by 4")
	}
	alc.AlertEntries = make([]ALCAlertEntry, 0, (fieldCount-4)/4)
	for i := 4; i < fieldCount; i = i + 4 {
		tmp := ALCAlertEntry{
			ManufacturerMnemonicCode: p.String(i, "manufacturer mnemonic code"),
			AlertIdentifier:          p.Int64(i+1, "alert identifier"),
			AlertInstance:            p.Int64(i+2, "alert instance"),
			RevisionCounter:          p.Int64(i+3, "revision counter"),
		}
		alc.AlertEntries = append(alc.AlertEntries, tmp)
	}

	return alc, p.Err()
}
