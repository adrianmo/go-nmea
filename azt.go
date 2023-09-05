package nmea

const (
	// TypeAZT type for AZT sentences
	TypeAZT = "AZT"
)

// AZT - Azimuth Thruster Message
// Format: $--AZT,x.x,x.x,x.x,x.x,x.x,x.x,x.x,x.x,x.x,x.x,x.x,x.x,x.x,x.x,x.x,x.x,x*hh<CR><LF>
// Example: $IIAZT,1,45.5,90.0,50.0,60.0,75.0,30.0,40.0,80.0,70.0,3,85.0,0,0,0,12345,9876,5432*2B
type AZT struct {
	BaseSentence
	ThrusterNo                            int64   // 1...9 Azimuth thruster number 1..9.
	SteeringCommandSetValue               float64 // 0...359.9 ° Steering command set value.
	SteeringMeasurementActualValue        float64 // 0...359.9 ° Steering measurement actual value.
	PrimeMoverCommandSetValue             float64 // -100...+100 % Prime mover (P.M.) RPM command set value. +100% equals nominal RPM (negative optional).
	PrimeMoverMeasurementActualValue      float64 // -100...+100 % P.M. RPM measurement actual value. +100% equals nominal RPM (negative optional).
	VariableSlippingClutchCommandSetValue float64 // 0...100 % Variable slipping clutch command set value. 100% equals fully engaged clutch (no slipping), 0% equals clutch disengaged. <zero value if not used>.
	PitchCommandSetValue                  float64 // -100...+100 % Pitch command set value. -100% equals max. negative pitch, 0% equals zero pitch, +100% equals max. positive pitch. <zero value if not used>.
	PitchMeasurementActualValue           float64 // -100...+100 % Pitch measurement actual value. <zero value if not used>.
	PMLoadLimitSetValue                   float64 // 0...120 % P.M. load limit set value (LLS). <zero value if not used>.
	PMLoadLimitCurrentMaxValue            float64 // 0...120 % P.M. load limit current max. allowed value. <zero value if not used>.
	PMLoadMeasurementActualValue          float64 // 0...120 % P.M. load measurement actual value (FPS). <zero value if not used>.
	ActiveControlStationNumber            int64   // 0...9 Active control station number in range of 1…9. Value is zero if no control station is selected.
	PropellerRPMMeasurementActualValue    float64 // 0...100 % Propeller RPM measurement actual value. +100% equals nominal RPM (used with slipping clutch).
	Reserved1                             float64 // Reserved for future use. <zero value>.
	Reserved2                             float64 // Reserved for future use. <zero value>.
	Reserved3                             float64 // Reserved for future use. <zero value>.
	ValueErrorStatusWord                  int64   // Value error status word for values 1...15, each bit indicates a single value error condition state either OK or ERROR.
	ControlStateWord1                     int64   // Control state word 1, each bit indicates a separate condition state either ON or OFF.
	ControlStateWord2                     int64   // Control state word 2 (pitch), each bit indicates a separate condition state either ON or OFF.
}

// newAZT constructor
func newAZT(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeAZT)
	return AZT{
		BaseSentence:                          s,
		ThrusterNo:                            p.Int64(0, "x"),
		SteeringCommandSetValue:               p.Float64(1, "value1"),
		SteeringMeasurementActualValue:        p.Float64(2, "value2"),
		PrimeMoverCommandSetValue:             p.Float64(3, "value3"),
		PrimeMoverMeasurementActualValue:      p.Float64(4, "value4"),
		VariableSlippingClutchCommandSetValue: p.Float64(5, "value5"),
		PitchCommandSetValue:                  p.Float64(6, "value6"),
		PitchMeasurementActualValue:           p.Float64(7, "value7"),
		PMLoadLimitSetValue:                   p.Float64(8, "value8"),
		PMLoadLimitCurrentMaxValue:            p.Float64(9, "value9"),
		PMLoadMeasurementActualValue:          p.Float64(10, "value10"),
		ActiveControlStationNumber:            p.Int64(11, "value11"),
		PropellerRPMMeasurementActualValue:    p.Float64(12, "value12"),
		Reserved1:                             p.Float64(13, "value13"),
		Reserved2:                             p.Float64(14, "value14"),
		Reserved3:                             p.Float64(15, "value15"),
		ValueErrorStatusWord:                  p.Int64(16, "vesw"),
		ControlStateWord1:                     p.Int64(17, "csw1"),
		ControlStateWord2:                     p.Int64(18, "csw2"),
	}, p.Err()
}
