package nmea

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

const (
	// TypePGN is type of PGN sentence for transferring single NMEA2000 frame as NMEA0183 sentence
	TypePGN = "PGN"
)

// PGN - transferring single NMEA2000 frame as NMEA0183 sentence
// https://opencpn.org/wiki/dokuwiki/lib/exe/fetch.php?media=opencpn:software:mxpgn_sentence.pdf
//
// Format: $--PGN,pppppp,aaaa,c--c*hh<CR><LF>
// Example: $MXPGN,01F112,2807,FC7FFF7FFF168012*11
type PGN struct {
	BaseSentence
	PGN      uint32 // PGN of NMEA2000 packet
	IsSend   bool   // is this sentence received or for sending
	Priority uint8  // 0-7
	Address  uint8  // depending on the IsSend field this is Source Address of received packet or Destination for send packet
	Data     []byte // 1-8 bytes. This is single N2K frame. N2K Fast-packets should be assembled from individual frames
}

// newPGN constructor
func newPGN(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypePGN)

	if len(p.Fields) != 3 {
		p.SetErr("fields", "invalid number of fields in sentence")
		return nil, p.Err()
	}
	pgn, err := strconv.ParseUint(p.Fields[0], 16, 24)
	if err != nil {
		p.err = fmt.Errorf("nmea: %s failed to parse PGN field: %w", p.Prefix(), err)
		return nil, p.Err()
	}
	attributes, err := strconv.ParseUint(p.Fields[1], 16, 16)
	if err != nil {
		p.err = fmt.Errorf("nmea: %s failed to parse attributes field: %w", p.Prefix(), err)
		return nil, p.Err()
	}
	dataLength := int((attributes >> 8) & 0b1111) // bits 8-11
	if dataLength*2 != (len(p.Fields[2])) {
		p.SetErr("dlc", "data length does not match actual data length")
		return nil, p.Err()
	}
	data, err := hex.DecodeString(p.Fields[2])
	if err != nil {
		p.err = fmt.Errorf("nmea: %s failed to decode data: %w", p.Prefix(), err)
		return nil, p.Err()
	}

	return PGN{
		BaseSentence: s,
		PGN:          uint32(pgn),
		IsSend:       attributes>>15 == 1,               // bit 15
		Priority:     uint8((attributes >> 12) & 0b111), // bits 12,13,14
		Address:      uint8(attributes),                 // bits 0-7
		Data:         data,
	}, p.Err()
}
