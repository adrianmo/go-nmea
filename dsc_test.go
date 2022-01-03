package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDSC(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  DSC
	}{
		{
			name: "good sentence",
			raw:  "$CDDSC,12,3380400790,12,06,00,1423108312,2019, ,  , S, E  *4a",
			msg: DSC{
				FormatSpecifier:             "12",
				Address:                     "3380400790",
				Category:                    "12",
				DistressCauseOrTeleCommand1: "06",
				CommandTypeOrTeleCommand2:   "00",
				PositionOrCanal:             "1423108312",
				TimeOrTelephoneNumber:       "2019",
				MSSI:                        " ",
				DistressCause:               "  ",
				Acknowledgement:             "S",
				ExpansionIndicator:          " E  ",
			},
		},
		{
			name: "good sentence Distress Alert Cancel",
			raw:  "$CDDSC,12,3381581370,12,06,00,1423108312,0236,3381581370, , S,    *20",
			msg: DSC{
				FormatSpecifier:             "12",
				Address:                     "3381581370",
				Category:                    "12",
				DistressCauseOrTeleCommand1: "06",
				CommandTypeOrTeleCommand2:   "00",
				PositionOrCanal:             "1423108312",
				TimeOrTelephoneNumber:       "0236",
				MSSI:                        "3381581370",
				DistressCause:               " ",
				Acknowledgement:             "S",
				ExpansionIndicator:          "    ",
			},
		},
		{
			name: "good sentence Non-Distress Call - Reply to Position Request\n",
			raw:  "$CDDSC,20,3381581370,00,21,26,1423108312,1902, , , B, E  *7B",
			msg: DSC{
				FormatSpecifier:             "20",
				Address:                     "3381581370",
				Category:                    "00",
				DistressCauseOrTeleCommand1: "21",
				CommandTypeOrTeleCommand2:   "26",
				PositionOrCanal:             "1423108312",
				TimeOrTelephoneNumber:       "1902",
				MSSI:                        " ",
				DistressCause:               " ",
				Acknowledgement:             "B",
				ExpansionIndicator:          " E  ",
			},
		},
		{
			name: "invalid nmea: Acknowledgement",
			raw:  "$CDDSC,20,3380400790,00,21,26,1423108312,2021,,,x, E*69",
			err:  "nmea: CDDSC invalid acknowledgement: x",
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
				dsc := m.(DSC)
				dsc.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, dsc)
			}
		})
	}
}
