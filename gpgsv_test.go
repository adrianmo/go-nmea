package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var gpgsvtests = []struct {
	name string
	raw  string
	err  string
	msg  GPGSV
}{
	{
		name: "good sentence",
		raw:  "$GPGSV,3,1,11,03,03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*77",
		msg: GPGSV{
			TotalMessages:   3,
			MessageNumber:   1,
			NumberSVsInView: 11,
			Info: []GPGSVInfo{
				GPGSVInfo{SVPRNNumber: 3, Elevation: 3, Azimuth: 111, SNR: 0},
				GPGSVInfo{SVPRNNumber: 4, Elevation: 15, Azimuth: 270, SNR: 0},
				GPGSVInfo{SVPRNNumber: 6, Elevation: 1, Azimuth: 10, SNR: 12},
				GPGSVInfo{SVPRNNumber: 13, Elevation: 6, Azimuth: 292, SNR: 0},
			},
		},
	},
	{
		name: "short",
		raw:  "$GPGSV,3,1,11,03,03,111,00,04,15,270,00,06,01,010,12*4A",
		msg: GPGSV{
			TotalMessages:   3,
			MessageNumber:   1,
			NumberSVsInView: 11,
			Info: []GPGSVInfo{
				GPGSVInfo{SVPRNNumber: 3, Elevation: 3, Azimuth: 111, SNR: 0},
				GPGSVInfo{SVPRNNumber: 4, Elevation: 15, Azimuth: 270, SNR: 0},
				GPGSVInfo{SVPRNNumber: 6, Elevation: 1, Azimuth: 10, SNR: 12},
			},
		},
	},
	{
		name: "invalid number of SVs",
		raw:  "$GPGSV,3,1,11.2,03,03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*6b",
		err:  "nmea: GPGSV invalid number of SVs in view: 11.2",
	},
	{
		name: "invalid total number of messages",
		raw:  "$GPGSV,A3,1,11,03,03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*36",
		err:  "nmea: GPGSV invalid total number of messages: A3",
	},
	{
		name: "invalid message number",
		raw:  "$GPGSV,3,A1,11,03,03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*36",
		err:  "nmea: GPGSV invalid message number: A1",
	},
	{
		name: "invalid SV prn number",
		raw:  "$GPGSV,3,1,11,A03,03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*36",
		err:  "nmea: GPGSV invalid SV prn number: A03",
	},
	{
		name: "invalid elevation",
		raw:  "$GPGSV,3,1,11,03,A03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*36",
		err:  "nmea: GPGSV invalid elevation: A03",
	},
	{
		name: "invalid azimuth",
		raw:  "$GPGSV,3,1,11,03,03,A111,00,04,15,270,00,06,01,010,12,13,06,292,00*36",
		err:  "nmea: GPGSV invalid azimuth: A111",
	},
	{
		name: "invalid SNR",
		raw:  "$GPGSV,3,1,11,03,03,111,A00,04,15,270,00,06,01,010,12,13,06,292,00*36",
		err:  "nmea: GPGSV invalid SNR: A00",
	},
	{
		name: "wrong type",
		raw:  "$GPXTE,A,A,4.07,L,N*6D",
		err:  "nmea: GPGSV invalid prefix: GPXTE",
	},
}

func TestGPGSV(t *testing.T) {
	for _, tt := range gpgsvtests {
		t.Run(tt.name, func(t *testing.T) {
			sent, err := ParseSentence(tt.raw)
			assert.NoError(t, err)
			gpgsv, err := NewGPGSV(sent)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				gpgsv.Sent = Sent{}
				assert.Equal(t, tt.msg, gpgsv)
			}
		})
	}
}
