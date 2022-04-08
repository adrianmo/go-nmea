package nmea

const (
	// TypePSONCMS is type of PSONCMS sentence for proprietary Xsens IMU/VRU/AHRS device
	TypePSONCMS = "SONCMS"
)

// PSONCMS is proprietary Xsens IMU/VRU/AHRS device sentence for quaternion, acceleration, rate of turn,
// magnetic Field, sensor temperature.
// https://www.xsens.com/hubfs/Downloads/Manuals/MT_Low-Level_Documentation.pdf (page 37)
//
// Format: $PSONCMS,Q.QQQQ,P.PPPP,R.RRRR,S.SSSS,XX.XXXX,YY.YYYY,ZZ.ZZZZ,
//			FF.FFFF,GG.GGGG,HH.HHHH,NN.NNNN,MM,MMMM,PP.PPPP,TT.T*hh<CR><LF>
// Example: $PSONCMS,0.0905,0.4217,0.9020,-0.0196,-1.7685,0.3861,-9.6648,-0.0116,0.0065,-0.0080,0.0581,0.3846,0.7421,33.1*76
type PSONCMS struct {
	BaseSentence
	Quaternion0       float64 // q0 from quaternions
	Quaternion1       float64 // q1 from quaternions
	Quaternion2       float64 // q2 from quaternions
	Quaternion3       float64 // q3 from quaternions
	AccelerationX     float64 // acceleration X in m/s2
	AccelerationY     float64 // acceleration Y in m/s2
	AccelerationZ     float64 // acceleration Z in m/s2
	RateOfTurnX       float64 // rate of turn X in rad/s
	RateOfTurnY       float64 // rate of turn Y in rad/s
	RateOfTurnZ       float64 // rate of turn Z in rad/s
	MagneticFieldX    float64 // magnetic field X in a.u.
	MagneticFieldY    float64 // magnetic field Y in a.u.
	MagneticFieldZ    float64 // magnetic field Z in a.u.
	SensorTemperature float64 // sensor temperature in degrees Celsius
}

// newPSONCMS constructor
func newPSONCMS(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypePSONCMS)
	m := PSONCMS{
		BaseSentence:      s,
		Quaternion0:       p.Float64(0, "q0 from quaternions"),
		Quaternion1:       p.Float64(1, "q1 from quaternions"),
		Quaternion2:       p.Float64(2, "q2 from quaternions"),
		Quaternion3:       p.Float64(3, "q3 from quaternions"),
		AccelerationX:     p.Float64(4, "acceleration X"),
		AccelerationY:     p.Float64(5, "acceleration Y"),
		AccelerationZ:     p.Float64(6, "acceleration Z"),
		RateOfTurnX:       p.Float64(7, "rate of turn X"),
		RateOfTurnY:       p.Float64(8, "rate of turn Y"),
		RateOfTurnZ:       p.Float64(9, "rate of turn Z"),
		MagneticFieldX:    p.Float64(10, "magnetic field X"),
		MagneticFieldY:    p.Float64(11, "magnetic field Y"),
		MagneticFieldZ:    p.Float64(12, "magnetic field Z"),
		SensorTemperature: p.Float64(13, "sensor temperature"),
	}
	return m, p.Err()
}
