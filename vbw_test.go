package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVBW(t *testing.T) {
	var tests = []struct {
		name string
		raw  string
		err  string
		msg  VBW
	}{
		{
			name: "good sentence",
			raw:  "$VMVBW,-7.1,0.1,A,,,V,,V,,V*65",
			msg: VBW{
				LongitudinalWaterSpeedKnots:         -7.1,
				TransverseWaterSpeedKnots:           0.1,
				WaterSpeedStatusValid:               true,
				WaterSpeedStatus:                    "A",
				LongitudinalGroundSpeedKnots:        0,
				TransverseGroundSpeedKnots:          0,
				GroundSpeedStatusValid:              false,
				GroundSpeedStatus:                   "V",
				SternTraverseWaterSpeedKnots:        0,
				SternTraverseWaterSpeedStatusValid:  false,
				SternTraverseWaterSpeedStatus:       "V",
				SternTraverseGroundSpeedKnots:       0,
				SternTraverseGroundSpeedStatusValid: false,
				SternTraverseGroundSpeedStatus:      "V",
			},
		},
		{
			name: "invalid nmea: LongitudinalWaterSpeedKnots",
			raw:  "$VMVBW,x,0.1,A,,,V,,V,,V*18",
			err:  "nmea: VMVBW invalid longitudinal water speed: x",
		},
		{
			name: "invalid nmea: TransverseWaterSpeedKnots",
			raw:  "$VMVBW,0.1,x,A,0.3,0.4,A,0.5,A,0.6,A*0b",
			err:  "nmea: VMVBW invalid transverse water speed: x",
		},
		{
			name: "invalid nmea: WaterSpeedStatusValid",
			raw:  "$VMVBW,0.1,0.2,X,0.3,0.4,A,0.5,A,0.6,A*46",
			err:  "nmea: VMVBW invalid water speed status: X",
		},
		{
			name: "invalid nmea: LongitudinalGroundSpeedKnots",
			raw:  "$VMVBW,0.1,0.2,A,X,0.4,A,0.5,A,0.6,A*2a",
			err:  "nmea: VMVBW invalid longitudinal ground speed: X",
		},
		{
			name: "invalid nmea: TransverseGroundSpeedKnots",
			raw:  "$VMVBW,0.1,0.2,A,0.3,X,A,0.5,A,0.6,A*2d",
			err:  "nmea: VMVBW invalid transverse ground speed: X",
		},
		{
			name: "invalid nmea: GroundSpeedStatusValid",
			raw:  "$VMVBW,0.1,0.2,A,0.3,0.4,X,0.5,A,0.6,A*46",
			err:  "nmea: VMVBW invalid ground speed status: X",
		},
		{
			name: "invalid nmea: SternTraverseWaterSpeedKnots",
			raw:  "$VMVBW,0.1,0.2,A,0.3,0.4,A,X,A,0.6,A*2c",
			err:  "nmea: VMVBW invalid stern traverse water speed: X",
		},
		{
			name: "invalid nmea: SternTraverseWaterSpeedStatusValid",
			raw:  "$VMVBW,0.1,0.2,A,0.3,0.4,A,0.5,X,0.6,A*46",
			err:  "nmea: VMVBW invalid stern water speed status: X",
		},
		{
			name: "invalid nmea: SternTraverseGroundSpeedKnots",
			raw:  "$VMVBW,0.1,0.2,A,0.3,0.4,A,0.5,A,X,A*2f",
			err:  "nmea: VMVBW invalid stern traverse ground speed: X",
		},
		{
			name: "invalid nmea: SternTraverseGroundSpeedStatusValid",
			raw:  "$VMVBW,0.1,0.2,A,0.3,0.4,A,0.5,A,0.6,X*46",
			err:  "nmea: VMVBW invalid stern ground speed status: X",
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
				mm := m.(VBW)
				mm.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, mm)
			}
		})
	}
}
