package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDSE(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  DSE
	}{
		{
			name: "good sentence, single dataset",
			raw:  "$CDDSE,1,1,A,3380400790,00,46504437*15",
			msg: DSE{
				TotalNumber:     1,
				Number:          1,
				Acknowledgement: AcknowledgementAutomaticDSE,
				MSSI:            "3380400790",
				DataSets: []DSEDataSet{
					{Code: "00", Data: "46504437"},
				},
			},
		},
		{
			name: "good sentence, single dataset",
			raw:  "$CDDSE,1,1,A,3380400790,00,46504437,01,16501437*17",
			msg: DSE{
				TotalNumber:     1,
				Number:          1,
				Acknowledgement: AcknowledgementAutomaticDSE,
				MSSI:            "3380400790",
				DataSets: []DSEDataSet{
					{Code: "00", Data: "46504437"},
					{Code: "01", Data: "16501437"},
				},
			},
		},
		{
			name: "invalid nmea: field count",
			raw:  "$CDDSE,1,1,x,3380400790,46504437*00",
			err:  "DSE is missing fields for parsing data sets",
		},
		{
			name: "invalid nmea: data set field count",
			raw:  "$CDDSE,1,1,A,3380400790,00,46504437,01*38",
			err:  "DSE data set field count is not exactly dividable by 2",
		},
		{
			name: "invalid nmea: Acknowledgement",
			raw:  "$CDDSE,1,1,x,3380400790,00,46504437*2c",
			err:  "nmea: CDDSE invalid acknowledgement: x",
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
				hdt := m.(DSE)
				hdt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, hdt)
			}
		})
	}
}
