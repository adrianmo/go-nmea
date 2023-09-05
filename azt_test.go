package nmea

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var azttests = []struct {
	name string
	raw  string
	err  string
	msg  AZT
}{
	{
		name: "good sentence",
		raw:  "$IIAZT,1,45.5,90.0,50.0,60.0,75.0,30.0,40.0,80.0,70.0,70.0,3,85.0,0,0,0,12345,9876,5432*70",
		msg: AZT{
			ThrusterNo:                            1,
			SteeringCommandSetValue:               45.5,
			SteeringMeasurementActualValue:        90.0,
			PrimeMoverCommandSetValue:             50.0,
			PrimeMoverMeasurementActualValue:      60.0,
			VariableSlippingClutchCommandSetValue: 75.0,
			PitchCommandSetValue:                  30.0,
			PitchMeasurementActualValue:           40.0,
			PMLoadLimitSetValue:                   80.0,
			PMLoadLimitCurrentMaxValue:            70.0,
			PMLoadMeasurementActualValue:          70.0,
			ActiveControlStationNumber:            3,
			PropellerRPMMeasurementActualValue:    85.0,
			Reserved1:                             0,
			Reserved2:                             0,
			Reserved3:                             0,
			ValueErrorStatusWord:                  12345,
			ControlStateWord1:                     9876,
			ControlStateWord2:                     5432,
		},
	},
	{
		name: "bad validity",
		raw:  "$IIAZT,1,45.5,90.0,50.0,60.0,75.0,30.0,40.0,80.0,70.0,70.0,3,85.0,0,0,0,12345,9876,5432*74",
		err:  "nmea: sentence checksum mismatch [70 != 74]",
	},
}

func TestAZT(t *testing.T) {
	for _, tt := range azttests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				azt := m.(AZT)
				azt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, azt)
			}
		})
	}
}
