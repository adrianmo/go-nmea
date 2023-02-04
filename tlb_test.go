package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTLB(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  TLB
	}{
		{
			name: "good sentence, single target",
			raw:  "$RATLB,1,XXX*20",
			msg: TLB{
				Targets: []TLBTarget{
					{TargetNumber: 1, TargetLabel: "XXX"},
				},
			},
		},
		{
			name: "good sentence, multiple targets",
			raw:  "$RATLB,1,XXX,2.0,YYY*55",
			msg: TLB{
				Targets: []TLBTarget{
					{TargetNumber: 1, TargetLabel: "XXX"},
					{TargetNumber: 2, TargetLabel: "YYY"},
				},
			},
		},
		{
			name: "invalid nmea: field count",
			raw:  "$RATLB,1*54",
			err:  "TLB is missing fields for parsing target pairs",
		},
		{
			name: "invalid nmea: data set field count",
			raw:  "$RATLB,1,XXX,2.0*20",
			err:  "TLB data set field count is not exactly dividable by 2",
		},
		{
			name: "invalid nmea: target number",
			raw:  "$RATLB,x,XXX*69",
			err:  "nmea: RATLB invalid target number: x",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				tlb := m.(TLB)
				tlb.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, tlb)
			}
		})
	}
}
