package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXDR(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  XDR
	}{
		{
			name: "good sentence with 1 measurement",
			raw:  "$SDXDR,C,23.15,C,WTHI*70",
			msg: XDR{
				Measurements: []XDRMeasurement{
					{
						TransducerType: "C",
						Value:          23.15,
						Unit:           "C",
						TransducerName: "WTHI",
					},
				},
			},
		},
		{
			name: "good sentence with 5 measurements",
			raw:  "$HCXDR,A,171,D,PITCH,A,-37,D,ROLL,G,367,,MAGX,G,2420,,MAGY,G,-8984,,MAGZ*41",
			msg: XDR{
				Measurements: []XDRMeasurement{
					{TransducerType: "A", Value: 171, Unit: "D", TransducerName: "PITCH"},
					{TransducerType: "A", Value: -37, Unit: "D", TransducerName: "ROLL"},
					{TransducerType: "G", Value: 367, Unit: "", TransducerName: "MAGX"},
					{TransducerType: "G", Value: 2420, Unit: "", TransducerName: "MAGY"},
					{TransducerType: "G", Value: -8984, Unit: "", TransducerName: "MAGZ"},
				},
			},
		},
		{
			name: "good sentence with 4 measurements",
			raw:  "$WIXDR,C,9.7,C,2,U,24.1,N,0,U,24.4,V,1,U,3.510,V,2*46",
			msg: XDR{
				Measurements: []XDRMeasurement{
					{TransducerType: "C", Value: 9.7, Unit: "C", TransducerName: "2"},
					// U+N - Voltage+Newtons? This is real sentence from actual vessel nmea0183 bus. Maybe misconfigured device?
					{TransducerType: "U", Value: 24.1, Unit: "N", TransducerName: "0"},
					{TransducerType: "U", Value: 24.4, Unit: "V", TransducerName: "1"},
					{TransducerType: "U", Value: 3.510, Unit: "V", TransducerName: "2"},
				},
			},
		},
		{
			name: "invalid nmea: odd number of fields",
			raw:  "$HCXDR,A,171,D,PITCH,A,-37,D,ROLL,G,367,,MAGX,G,2420,MAGY,G,-8984,,MAGZ*6d",
			err:  "XDR field count is not exactly dividable by 4",
		},
		{
			name: "invalid nmea: TransducerType",
			raw:  "$SDXDR,x,23.15,C,WTHI*4b",
			err:  "nmea: SDXDR invalid transducer type: x",
		},
		{
			name: "invalid nmea: Value",
			raw:  "$SDXDR,C,23.x,C,WTHI*0C",
			err:  "nmea: SDXDR invalid measurement value: 23.x",
		},
		{
			name: "invalid nmea: Unit",
			raw:  "$SDXDR,C,23.15,x,WTHI*4b",
			err:  "nmea: SDXDR invalid measurement unit: x",
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
				xdr := m.(XDR)
				xdr.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, xdr)
			}
		})
	}
}
