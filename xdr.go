package nmea

import "errors"

const (
	// TypeXDR type of XDR sentence for Transducer Measurement
	TypeXDR = "XDR"
)

const (
	// TransducerAngularDisplacementXDR is transducer type for Angular displacement
	TransducerAngularDisplacementXDR = "A"
	// TransducerTemperatureXDR is transducer type for Temperature
	TransducerTemperatureXDR = "C"
	// TransducerDepthXDR is transducer type for Depth
	TransducerDepthXDR = "D"
	// TransducerFrequencyXDR is transducer type for Frequency
	TransducerFrequencyXDR = "F"
	// TransducerHumidityXDR is transducer type for Humidity
	TransducerHumidityXDR = "H"
	// TransducerForceXDR is transducer type for Force
	TransducerForceXDR = "N"
	// TransducerPressureXDR is transducer type for Pressure
	TransducerPressureXDR = "P"
	// TransducerFlowXDR is transducer type for Flow
	TransducerFlowXDR = "R"
	// TransducerAbsoluteHumidityXDR is transducer type for Absolute humidity
	TransducerAbsoluteHumidityXDR = "B"
	// TransducerGenericXDR is transducer type for Generic
	TransducerGenericXDR = "G"
	// TransducerCurrentXDR is transducer type for Current
	TransducerCurrentXDR = "I"
	// TransducerSalinityXDR is transducer type for Salinity
	TransducerSalinityXDR = "L"
	// TransducerSwitchValveXDR is transducer type for Switch, valve
	TransducerSwitchValveXDR = "S"
	// TransducerTachometerXDR is transducer type for Tachometer
	TransducerTachometerXDR = "T"
	// TransducerVoltageXDR is transducer type for Voltage
	TransducerVoltageXDR = "U"
	// TransducerVolumeXDR is transducer type for Volume
	TransducerVolumeXDR = "V"
)

// XDR - Transducer Measurement
// https://gpsd.gitlab.io/gpsd/NMEA.html#_xdr_transducer_measurement
// https://www.eye4software.com/hydromagic/documentation/articles-and-howtos/handling-nmea0183-xdr/
//
// Format: $--XDR,a,x.x,a,c--c, ..... *hh<CR><LF>
// Example: $HCXDR,A,171,D,PITCH,A,-37,D,ROLL,G,367,,MAGX,G,2420,,MAGY,G,-8984,,MAGZ*41
//			$SDXDR,C,23.15,C,WTHI*70
type XDR struct {
	BaseSentence
	Measurements []XDRMeasurement
}

// XDRMeasurement is measurement recorded by transducer
type XDRMeasurement struct {
	// TransducerType is type of transducer
	// * A - Angular displacement
	// * C - Temperature
	// * D - Depth
	// * F - Frequency
	// * H - Humidity
	// * N - Force
	// * P - Pressure
	// * R - Flow
	// * B - Absolute humidity
	// * G - Generic
	// * I - Current
	// * L - Salinity
	// * S - Switch, valve
	// * T - Tachometer
	// * U - Voltage
	// * V - Volume
	// could be more
	TransducerType string

	// Value of measurement
	Value float64

	// Unit of measurement
	// * "" - could be empty!
	// * A - Amperes
	// * B - Bars | Binary
	// * C - Celsius
	// * D - Degrees
	// * H - Hertz
	// * I - liters/second
	// * K - Kelvin | Density, kg/m3 kilogram per cubic metre
	// * M - Meters | Cubic Meters (m3)
	// * N - Newton
	// * P - percent of full range | Pascal
	// * R - RPM
	// * S - Parts per thousand
	// * V - Volts
	// could be more
	Unit string

	// TransducerName is name of transducer where measurement was recorded
	TransducerName string
}

// newXDR constructor
func newXDR(s BaseSentence) (XDR, error) {
	p := NewParser(s)
	p.AssertType(TypeXDR)

	xdr := XDR{
		BaseSentence: s,
		Measurements: nil,
	}

	if len(p.Fields)%4 != 0 {
		return xdr, errors.New("XDR field count is not exactly dividable by 4")
	}

	xdr.Measurements = make([]XDRMeasurement, 0, len(s.Fields)/4)
	for i := 0; i < len(s.Fields); {
		tmp := XDRMeasurement{
			TransducerType: p.EnumString(
				i,
				"transducer type",
				TransducerAngularDisplacementXDR,
				TransducerTemperatureXDR,
				TransducerDepthXDR,
				TransducerFrequencyXDR,
				TransducerHumidityXDR,
				TransducerForceXDR,
				TransducerPressureXDR,
				TransducerFlowXDR,
				TransducerAbsoluteHumidityXDR,
				TransducerGenericXDR,
				TransducerCurrentXDR,
				TransducerSalinityXDR,
				TransducerSwitchValveXDR,
				TransducerTachometerXDR,
				TransducerVoltageXDR,
				TransducerVolumeXDR,
			),
			Value: p.Float64(i+1, "measurement value"),
			Unit: p.EnumString(
				i+2,
				"measurement unit",
				UnitAmpere,
				UnitBars,
				UnitBinary,
				UnitCelsius,
				UnitDegrees,
				UnitHertz,
				UnitLitresPerSecond,
				UnitKelvin,
				UnitKilogramPerCubicMetre,
				UnitMeters,
				UnitCubicMeters,
				UnitRevolutionsPerMinute,
				UnitPercent,
				UnitPascal,
				UnitPartsPerThousand,
				UnitVolts,
			),
			TransducerName: p.String(i+3, "transducer name"),
		}
		xdr.Measurements = append(xdr.Measurements, tmp)
		i += 4
	}
	return xdr, p.Err()
}
