package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pgrmttests = []struct {
	name string
	raw  string
	err  string
	msg  PGRMT
}{
	{
		name: "typical sentence",
		raw:  "$PGRMT,GPS24xd-HVS VER 2.30,,,,,,,,*10",
		msg: PGRMT{
			ModelAndFirmwareVersion: "GPS24xd-HVS VER 2.30",
		},
	},
	{
		name: "all good",
		raw:  "$PGRMT,GOOD GPS VER 1.0,P,P,R,R,P,C,32,R*39",
		msg: PGRMT{
			ModelAndFirmwareVersion: "GOOD GPS VER 1.0",
			ROMChecksumTest:         PassPGRMT,
			ReceiverFailureDiscrete: PassPGRMT,
			StoredDataLost:          DataRetainedPGRMT,
			RealtimeClockLost:       DataRetainedPGRMT,
			OscillatorDriftDiscrete: PassPGRMT,
			DataCollectionDiscrete:  DataCollectingPGRMT,
			SensorTemperature:       32,
			SensorConfigurationData: DataRetainedPGRMT,
		},
	},
	{
		name: "all bad",
		raw:  "$PGRMT,BAD GPS VER 1.0,F,F,L,L,F,,-64,L*18",
		msg: PGRMT{
			ModelAndFirmwareVersion: "BAD GPS VER 1.0",
			ROMChecksumTest:         FailPGRMT,
			ReceiverFailureDiscrete: FailPGRMT,
			StoredDataLost:          DataLostPGRMT,
			RealtimeClockLost:       DataLostPGRMT,
			OscillatorDriftDiscrete: FailPGRMT,
			SensorTemperature:       -64,
			SensorConfigurationData: DataLostPGRMT,
		},
	},
}

func TestPGRMT(t *testing.T) {
	for _, tt := range pgrmttests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				pgrmt := m.(PGRMT)
				pgrmt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, pgrmt)
			}
		})
	}
}
