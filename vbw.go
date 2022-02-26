package nmea

const (
	// TypeVBW type of VBW sentence for Dual Ground/Water Speed
	TypeVBW = "VBW"
)

// VBW - Dual Ground/Water Speed
// https://gpsd.gitlab.io/gpsd/NMEA.html#_vbw_dual_groundwater_speed
//
// Format: $--VBW,x.x,x.x,A,x.x,x.x,A,x.x,A,x.x,A*hh<CR><LF>
// Example: $VMVBW,-7.1,0.1,A,,,V,,V,,V*65
type VBW struct {
	BaseSentence
	LongitudinalWaterSpeedKnots float64 // longitudinal water speed, "-" means astern, knots
	TransverseWaterSpeedKnots   float64 // transverse water speed, "-" means port, knots
	WaterSpeedStatusValid       bool    // A = true
	WaterSpeedStatus            string  // A = valid, V = invalid

	LongitudinalGroundSpeedKnots float64 // longitudinal ground speed, "-" means astern, knots
	TransverseGroundSpeedKnots   float64 // transverse ground speed, "-" means port, knots
	GroundSpeedStatusValid       bool    // A = true
	GroundSpeedStatus            string  // A = valid, V = invalid

	SternTraverseWaterSpeedKnots       float64 // Stern traverse water speed, knots (NMEA 3 and above)
	SternTraverseWaterSpeedStatusValid bool    // A = true
	SternTraverseWaterSpeedStatus      string  // A = valid, V = invalid (NMEA 3 and above)

	SternTraverseGroundSpeedKnots       float64 // Stern traverse ground speed, knots (NMEA 3 and above)
	SternTraverseGroundSpeedStatusValid bool    // A = true
	SternTraverseGroundSpeedStatus      string  // A = valid, V = invalid (NMEA 3 and above)
}

// newVBW constructor
func newVBW(s BaseSentence) (VBW, error) {
	p := NewParser(s)
	p.AssertType(TypeVBW)

	m := VBW{
		BaseSentence:                s,
		LongitudinalWaterSpeedKnots: p.Float64(0, "longitudinal water speed"),
		TransverseWaterSpeedKnots:   p.Float64(1, "transverse water speed"),
		WaterSpeedStatusValid:       p.String(2, "water speed status valid") == StatusValid,
		WaterSpeedStatus:            p.EnumString(2, "water speed status", StatusValid, StatusInvalid),

		LongitudinalGroundSpeedKnots: p.Float64(3, "longitudinal ground speed"),
		TransverseGroundSpeedKnots:   p.Float64(4, "transverse ground speed"),
		GroundSpeedStatusValid:       p.String(5, "ground speed status valid") == StatusValid,
		GroundSpeedStatus:            p.EnumString(5, "ground speed status", StatusValid, StatusInvalid),
	}
	if len(p.Fields) > 6 {
		m.SternTraverseWaterSpeedKnots = p.Float64(6, "stern traverse water speed")
		m.SternTraverseWaterSpeedStatusValid = p.String(7, "stern water speed status valid") == StatusValid
		m.SternTraverseWaterSpeedStatus = p.EnumString(7, "stern water speed status", StatusValid, StatusInvalid)

		m.SternTraverseGroundSpeedKnots = p.Float64(8, "stern traverse ground speed")
		m.SternTraverseGroundSpeedStatusValid = p.String(9, "stern ground speed status valid") == StatusValid
		m.SternTraverseGroundSpeedStatus = p.EnumString(9, "stern ground speed status", StatusValid, StatusInvalid)
	}

	return m, p.Err()
}
