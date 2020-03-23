package nmea

import (
	"fmt"
	"strings"
)

const (
	// SentenceStart is the token to indicate the start of a sentence.
	SentenceStart = "$"

	// SentenceStartEncapsulated is the token to indicate the start of encapsulated data.
	SentenceStartEncapsulated = "!"

	// FieldSep is the token to delimit fields of a sentence.
	FieldSep = ","

	// ChecksumSep is the token to delimit the checksum of a sentence.
	ChecksumSep = "*"
)

// Sentence interface for all NMEA sentence
type Sentence interface {
	fmt.Stringer
	Prefix() string
	DataType() string
	TalkerID() string
}

// BaseSentence contains the information about the NMEA sentence
type BaseSentence struct {
	Talker   string   // The talker id (e.g GP)
	Type     string   // The data type (e.g GSA)
	Fields   []string // Array of fields
	Checksum string   // The Checksum
	Raw      string   // The raw NMEA sentence received
}

// Prefix returns the talker and type of message
func (s BaseSentence) Prefix() string {
	return s.Talker + s.Type
}

// DataType returns the type of the message
func (s BaseSentence) DataType() string {
	return s.Type
}

// TalkerID returns the talker of the message
func (s BaseSentence) TalkerID() string {
	return s.Talker
}

// String formats the sentence into a string
func (s BaseSentence) String() string { return s.Raw }

// parseSentence parses a raw message into it's fields
func parseSentence(raw string) (BaseSentence, error) {
	raw = strings.TrimSpace(raw)
	startIndex := strings.IndexAny(raw, SentenceStart+SentenceStartEncapsulated)
	if startIndex != 0 {
		return BaseSentence{}, fmt.Errorf("nmea: sentence does not start with a '$' or '!'")
	}
	sumSepIndex := strings.Index(raw, ChecksumSep)
	if sumSepIndex == -1 {
		return BaseSentence{}, fmt.Errorf("nmea: sentence does not contain checksum separator")
	}
	var (
		fieldsRaw   = raw[startIndex+1 : sumSepIndex]
		fields      = strings.Split(fieldsRaw, FieldSep)
		checksumRaw = strings.ToUpper(raw[sumSepIndex+1:])
		checksum    = Checksum(fieldsRaw)
	)
	// Validate the checksum
	if checksum != checksumRaw {
		return BaseSentence{}, fmt.Errorf(
			"nmea: sentence checksum mismatch [%s != %s]", checksum, checksumRaw)
	}
	talker, typ := parsePrefix(fields[0])
	return BaseSentence{
		Talker:   talker,
		Type:     typ,
		Fields:   fields[1:],
		Checksum: checksumRaw,
		Raw:      raw,
	}, nil
}

// parsePrefix takes the first field and splits it into a talker id and data type.
func parsePrefix(s string) (string, string) {
	if strings.HasPrefix(s, "PMTK") {
		return "PMTK", s[4:]
	}
	if strings.HasPrefix(s, "P") {
		return "P", s[1:]
	}
	if len(s) < 2 {
		return s, ""
	}
	return s[:2], s[2:]
}

// Checksum xor all the bytes in a string an return it
// as an uppercase hex string
func Checksum(s string) string {
	var checksum uint8
	for i := 0; i < len(s); i++ {
		checksum ^= s[i]
	}
	return fmt.Sprintf("%02X", checksum)
}

// Parse parses the given string into the correct sentence type.
func Parse(raw string) (Sentence, error) {
	s, err := parseSentence(raw)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(s.Raw, SentenceStart) {
		// MTK message types share the same format
		// so we return the same struct for all types.
		switch s.Talker {
		case TypeMTK:
			return newMTK(s)
		}
		switch s.Type {
		case TypeRMC:
			return newRMC(s)
		case TypeGGA:
			return newGGA(s)
		case TypeGSA:
			return newGSA(s)
		case TypeGLL:
			return newGLL(s)
		case TypeVTG:
			return newVTG(s)
		case TypeZDA:
			return newZDA(s)
		case TypePGRME:
			return newPGRME(s)
		case TypeGSV:
			return newGSV(s)
		case TypeHDT:
			return newHDT(s)
		case TypeGNS:
			return newGNS(s)
		case TypeTHS:
			return newTHS(s)
		case TypeWPL:
			return newWPL(s)
		case TypeRTE:
			return newRTE(s)
		case TypeVHW:
			return newVHW(s)
		case TypeDPT:
			return newDPT(s)
		case TypeDBT:
			return newDBT(s)
		case TypeDBS:
			return newDBS(s)
		}
	}
	if strings.HasPrefix(s.Raw, SentenceStartEncapsulated) {
		switch s.Type {
		case TypeVDM, TypeVDO:
			return newVDMVDO(s)
		}
	}
	return nil, fmt.Errorf("nmea: sentence prefix '%s' not supported", s.Prefix())
}
