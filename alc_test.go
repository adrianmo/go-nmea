package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestALC(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  ALC
	}{
		{
			name: "good sentence, single entry",
			raw:  "$FBALC,02,01,03,01,FEB,01,02,03*0A",
			msg: ALC{
				NumFragments:   2,
				FragmentNumber: 1,
				MessageID:      3,
				EntriesNumber:  1,
				AlertEntries: []ALCAlertEntry{
					{
						ManufacturerMnemonicCode: "FEB",
						AlertIdentifier:          01,
						AlertInstance:            02,
						RevisionCounter:          03,
					},
				},
			},
		},
		{
			name: "good sentence, multiple entries",
			raw:  "$FBALC,02,01,03,02,FEB,01,02,03,TEB,02,03,04*5f",
			msg: ALC{
				NumFragments:   2,
				FragmentNumber: 1,
				MessageID:      3,
				EntriesNumber:  2,
				AlertEntries: []ALCAlertEntry{
					{
						ManufacturerMnemonicCode: "FEB",
						AlertIdentifier:          01,
						AlertInstance:            02,
						RevisionCounter:          03,
					},
					{
						ManufacturerMnemonicCode: "TEB",
						AlertIdentifier:          02,
						AlertInstance:            03,
						RevisionCounter:          04,
					},
				},
			},
		},
		{
			name: "good sentence, no entries",
			raw:  "$FBALC,02,01,03,00*4a",
			msg: ALC{
				NumFragments:   2,
				FragmentNumber: 1,
				MessageID:      3,
				EntriesNumber:  0,
				AlertEntries:   nil,
			},
		},
		{
			name: "invalid nmea: invalid number of fields",
			raw:  "$FBALC,02,01,03,01,FEB,01,02*25",
			err:  "ALC data set field count is not exactly dividable by 4",
		},
		{
			name: "invalid nmea: NumFragments",
			raw:  "$FBALC,0x,01,03,01,FEB,01,02,03*40",
			err:  "nmea: FBALC invalid number of fragments: 0x",
		},
		{
			name: "invalid nmea: FragmentNumber",
			raw:  "$FBALC,02,0a,03,01,FEB,01,02,03*5a",
			err:  "nmea: FBALC invalid fragment number: 0a",
		},
		{
			name: "invalid nmea: MessageID",
			raw:  "$FBALC,02,01,0x,01,FEB,01,02,03*41",
			err:  "nmea: FBALC invalid message ID: 0x",
		},
		{
			name: "invalid nmea: EntriesNumber",
			raw:  "$FBALC,02,01,03,0x,FEB,01,02,03*43",
			err:  "nmea: FBALC invalid entries number: 0x",
		},
		{
			name: "invalid nmea: AlertIdentifier",
			raw:  "$FBALC,02,01,03,01,FEB,0x,02,03*43",
			err:  "nmea: FBALC invalid alert identifier: 0x",
		},
		{
			name: "invalid nmea: AlertInstance",
			raw:  "$FBALC,02,01,03,01,FEB,01,0x,03*40",
			err:  "nmea: FBALC invalid alert instance: 0x",
		},
		{
			name: "invalid nmea: RevisionCounter",
			raw:  "$FBALC,02,01,03,01,FEB,01,02,0x*41",
			err:  "nmea: FBALC invalid revision counter: 0x",
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
				arc := m.(ALC)
				arc.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, arc)
			}
		})
	}
}
