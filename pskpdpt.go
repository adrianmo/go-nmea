package nmea

const (
	// TypePSKPDPT type for proprietary Skipper PSKPDPT sentences
	TypePSKPDPT = "SKPDPT"
)

// PSKPDPT - Depth of Water for multiple transducer installation
// https://www.alphatronmarine.com/files/products/120-echonav-skipper-gds101-instoper-manual-12-6-2017_1556099135_47f5f8d1.pdf (page 56, Edition: 2017.06.12)
// https://www.kongsberg.com/globalassets/maritime/km-products/product-documents/164821aa_rd301_instruction_manual_lr.pdf (page 2, 857-164821aa)
//
// Format: $PSKPDPT,x.x,x.x,x.x,xx,xx,c--c*hh<CR><LF>
// Example: $PSKPDPT,0002.5,+00.0,0010,10,03,*77
type PSKPDPT struct {
	BaseSentence
	// Depth is water depth relative to transducer, meters
	Depth float64
	// Offset from transducer, meters
	Offset float64
	// RangeScale is Maximum range scale in use, meters
	RangeScale float64
	// BottomEchoStrength is Bottom echo strength (0,9)
	BottomEchoStrength int64
	// ChannelNumber is Echo sounder channel number (0-99) (1 = 38 kHz. 2 = 50 kHz. 3 = 200 kHz)
	ChannelNumber int64
	// TransducerLocation is Transducer location. Text string, indicating transducer position: FWD/AFT/PORT/STB.
	// If position is not preset by operator, empty field is provided.
	TransducerLocation string
}

// newPSKPDPT constructor
func newPSKPDPT(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypePSKPDPT)
	sentence := PSKPDPT{
		BaseSentence:       s,
		Depth:              p.Float64(0, "depth"),
		Offset:             p.Float64(1, "offset"),
		RangeScale:         p.Float64(2, "range scale"),
		BottomEchoStrength: p.Int64(3, "bottom echo strength"),
		ChannelNumber:      p.Int64(4, "channel number"),
		TransducerLocation: p.String(5, "transducer location"),
	}
	return sentence, p.Err()
}
