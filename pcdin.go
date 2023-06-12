package nmea

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

const (
	// TypePCDIN is type of PCDIN sentence for SeaSmart.Net Protocol
	TypePCDIN = "CDIN"
)

// PCDIN - SeaSmart.Net Protocol transfers NMEA2000 message as NMEA0183 sentence
// http://www.seasmart.net/pdf/SeaSmart_HTTP_Protocol_RevG_043012.pdf (SeaSmart.Net Protocol Specification Version 1.7)
//
// Note: older SeaSmart.Net Protocol versions have different amount of fields
//
// Format:  $PCDIN,hhhhhh,hhhhhhhh,hh,h--h*hh<CR><LF>
// Example: $PCDIN,01F112,000C72EA,09,28C36A0000B40AFD*56
type PCDIN struct {
	BaseSentence
	PGN       uint32 // PGN of NMEA2000 packet
	Timestamp uint32 // ticks since something
	Source    uint8  // 0-255
	Data      []byte // can be more than 8 bytes i.e can contain assembled fast packets
}

// newPCDIN constructor
func newPCDIN(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypePCDIN)

	if len(p.Fields) != 4 {
		p.SetErr("fields", "invalid number of fields in sentence")
		return nil, p.Err()
	}
	pgn, err := strconv.ParseUint(p.Fields[0], 16, 24)
	if err != nil {
		p.err = fmt.Errorf("failed to parse PGN field, err: %w", err)
		return nil, p.Err()
	}
	timestamp, err := strconv.ParseUint(p.Fields[1], 16, 32)
	if err != nil {
		p.err = fmt.Errorf("failed to parse timestamp field, err: %w", err)
		return nil, p.Err()
	}
	source, err := strconv.ParseUint(p.Fields[2], 16, 8)
	if err != nil {
		p.err = fmt.Errorf("failed to parse source field, err: %w", err)
		return nil, p.Err()
	}
	data, err := hex.DecodeString(p.Fields[3])
	if err != nil {
		p.err = fmt.Errorf("failed to decode data, err: %w", err)
		return nil, p.Err()
	}

	return PCDIN{
		BaseSentence: s,
		PGN:          uint32(pgn),
		Timestamp:    uint32(timestamp),
		Source:       uint8(source),
		Data:         data,
	}, p.Err()
}
