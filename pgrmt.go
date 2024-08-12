package nmea

const (
	// TypePGRMT type for PGRMT sentences
	TypePGRMT = "GRMT"
	// PassPGRMT Self-Test Passed
	PassPGRMT = "P"
	// FailPGRMT Self-Test Failed
	FailPGRMT = "F"
	// DataRetainedPGRMT Data Retained
	DataRetainedPGRMT = "R"
	// DataLostPGRMT Data Lost
	DataLostPGRMT = "L"
	// DataCollectingPGRMT Data Collecting
	DataCollectingPGRMT = "C"
)

// PGRMT is Sensor Status Information (Garmin proprietary sentence)
// https://developer.garmin.com/downloads/legacy/uploads/2015/08/190-00684-00.pdf
// $PGRMT,<0>,<1>,<2>,<3>,<4>,<5>,<6>,<7>,<8>*hh<CR><LF>
// Format: $PGRMT,xxxxxxxxxx,A,A,A,A,A,A,N,A*hh<CR><LF>
// Example: $PGRMT,GPS24xd-HVS VER 2.30,,,,,,,,*10
type PGRMT struct {
	BaseSentence
	ModelAndFirmwareVersion string
	ROMChecksumTest         string  // "P" = pass, "F" = fail
	ReceiverFailureDiscrete string  // "P" = pass, "F" = fail
	StoredDataLost          string  // "R" = retained, "L" = lost
	RealtimeClockLost       string  // "R" = retained, "L" = lost
	OscillatorDriftDiscrete string  // "P" = pass, "F" = fail
	DataCollectionDiscrete  string  // "C" = collecting, "" = not collecting
	SensorTemperature       float64 // Degrees C
	SensorConfigurationData string  // "R" = retained, "L" = lost
}

// newPGRMT constructor
func newPGRMT(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypePGRMT)

	return PGRMT{
		BaseSentence:            s,
		ModelAndFirmwareVersion: p.String(0, "product, model and software version"),
		ROMChecksumTest:         p.EnumString(1, "rom checksum test", PassPGRMT, FailPGRMT),
		ReceiverFailureDiscrete: p.EnumString(2, "receiver failure discrete", PassPGRMT, FailPGRMT),
		StoredDataLost:          p.EnumString(3, "stored data lost", DataRetainedPGRMT, DataLostPGRMT),
		RealtimeClockLost:       p.EnumString(4, "realtime clock lost", DataRetainedPGRMT, DataLostPGRMT),
		OscillatorDriftDiscrete: p.EnumString(5, "oscillator drift discrete", PassPGRMT, FailPGRMT),
		DataCollectionDiscrete:  p.EnumString(6, "data collection discrete", DataCollectingPGRMT),
		SensorTemperature:       p.Float64(7, "sensor temperature in degrees celsius"),
		SensorConfigurationData: p.EnumString(8, "sensor configuration data", DataRetainedPGRMT, DataLostPGRMT),
	}, p.Err()
}
