package nmea

import "errors"

const (
	// TypeTLB type of TLB target label.
	TypeTLB = "TLB"
)

// TLB is sentence for target label.
// https://fcc.report/FCC-ID/ADB9ZWRTR100/2768717.pdf (page 8) FURUNO MARINE RADAR, model FAR-15XX manual
//
// Format: $--TLB,x.x,c--c,x.x,c--c,...x.x,c--c*hh<CR><LF>
// Example: $CDTLB,1,XXX,2.0,YYY*41
type TLB struct {
	BaseSentence
	Targets []TLBTarget
}

// TLBTarget is instance of target for TLB sentence
type TLBTarget struct {
	// TargetNumber is target number “n” reported by the device (1 - 1023)
	TargetNumber float64
	// TargetLabel is label assigned to target “n” (TT=000 - 999, AIS=000000000 - 999999999). Could be empty.
	TargetLabel string
}

// newTLB constructor
func newTLB(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeTLB)
	tlb := TLB{
		BaseSentence: s,
		Targets:      make([]TLBTarget, 0),
	}
	fieldCount := len(p.Fields)
	if fieldCount < 2 {
		return tlb, errors.New("TLB is missing fields for parsing target pairs")
	}
	if fieldCount%2 != 0 {
		return tlb, errors.New("TLB data set field count is not exactly dividable by 2")
	}
	tlb.Targets = make([]TLBTarget, 0, fieldCount/2)
	for i := 0; i < fieldCount; i = i + 2 {
		tmp := TLBTarget{
			TargetNumber: p.Float64(0+i, "target number"),
			TargetLabel:  p.String(1+i, "target label"),
		}
		tlb.Targets = append(tlb.Targets, tmp)
	}
	return tlb, p.Err()
}
