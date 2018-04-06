package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var glgsvtests = []struct {
	name string
	raw  string
	err  string
	msg  GLGSV
}{
	{
		name: "good sentence",
		raw:  "$GLGSV,3,1,11,03,03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*6B",
		msg: GLGSV{
			TotalMessages:   3,
			MessageNumber:   1,
			NumberSVsInView: 11,
			Info: []GLGSVInfo{
				GLGSVInfo{SVPRNNumber: 3, Elevation: 3, Azimuth: 111, SNR: 0},
				GLGSVInfo{SVPRNNumber: 4, Elevation: 15, Azimuth: 270, SNR: 0},
				GLGSVInfo{SVPRNNumber: 6, Elevation: 1, Azimuth: 10, SNR: 12},
				GLGSVInfo{SVPRNNumber: 13, Elevation: 6, Azimuth: 292, SNR: 0},
			},
		},
	},
	{
		name: "short sentence",
		raw:  "$GLGSV,3,1,11,03,03,111,00,04,15,270,00,06,01,010,12*56",
		msg: GLGSV{
			TotalMessages:   3,
			MessageNumber:   1,
			NumberSVsInView: 11,
			Info: []GLGSVInfo{
				GLGSVInfo{SVPRNNumber: 3, Elevation: 3, Azimuth: 111, SNR: 0},
				GLGSVInfo{SVPRNNumber: 4, Elevation: 15, Azimuth: 270, SNR: 0},
				GLGSVInfo{SVPRNNumber: 6, Elevation: 1, Azimuth: 10, SNR: 12},
			},
		},
	},
	{
		name: "invalid number of svs",
		raw:  "$GLGSV,3,1,11.2,03,03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*77",
		err:  "nmea: GLGSV invalid number of SVs in view: 11.2",
	},
	{
		name: "invalid number of messages",
		raw:  "$GLGSV,A3,1,11,03,03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*2A",
		err:  "nmea: GLGSV invalid total number of messages: A3",
	},
	{
		name: "invalid message number",
		raw:  "$GLGSV,3,A1,11,03,03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*2A",
		err:  "nmea: GLGSV invalid message number: A1",
	},
	{
		name: "invalid SV prn number",
		raw:  "$GLGSV,3,1,11,A03,03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*2A",
		err:  "nmea: GLGSV invalid SV prn number: A03",
	},
	{
		name: "invalid elevation",
		raw:  "$GLGSV,3,1,11,03,A03,111,00,04,15,270,00,06,01,010,12,13,06,292,00*2A",
		err:  "nmea: GLGSV invalid elevation: A03",
	},
	{
		name: "invalid azimuth",
		raw:  "$GLGSV,3,1,11,03,03,A111,00,04,15,270,00,06,01,010,12,13,06,292,00*2A",
		err:  "nmea: GLGSV invalid azimuth: A111",
	},
	{
		name: "invalid SNR",
		raw:  "$GLGSV,3,1,11,03,03,111,A00,04,15,270,00,06,01,010,12,13,06,292,00*2A",
		err:  "nmea: GLGSV invalid SNR: A00",
	},
}

func TestGLGSV(t *testing.T) {
	for _, tt := range glgsvtests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				glgsv := m.(GLGSV)
				glgsv.Sent = Sent{}
				assert.Equal(t, tt.msg, glgsv)
			}
		})
	}
}
