package nmea

const (
	// TypeVSD type of VSD sentence for AIS voyage static data.
	TypeVSD = "VSD"
)

// VSD is sentence for AIS voyage static data.
// https://fcc.report/FCC-ID/ADB9ZWRTR100/2768717.pdf (page 10) FURUNO MARINE RADAR, model FAR-15XX manual
// http://www.annoyingdesigns.com/vsd/VSDControl.pdf
//
// Format: $--VSD,x.x,x.x,x.x,c--c,hhmmss.ss,xx,xx,x.x,x.x*hh<CR><LF>
// Example: $RAVSD,0,4.5,6,@@@@@@@@@@@@@@@@@@@@,220516,01,02,8,*12<CR><LF>
type VSD struct {
	BaseSentence

	// From: https://www.itu.int/rec/R-REC-M.1371-5-201402-I/en (page 113)
	// * null means unchanged
	// * 0 = not available or no ship = default
	// * 1-99 = as defined in § 3.3.2
	// * 100-199 = reserved, for regional use
	// * 200-255 = reserved, for future use
	TypeOfShipAndCargo Int64 // 0

	// StaticDraughtMeters is maximum present static draught in meters (0 - 25.5, null). null means unchanged, 0 means not available
	StaticDraughtMeters Float64 // 1

	// Persons on-board (0 - 8191, null)
	PersonsOnBoard Int64 // 2

	// Destination (Alphanumeric character, null)
	Destination string // 3

	// NOTE: we are not combining time+day+month here to time.Time because - some of these fields can be empty.

	// Estimated UTC of arrival at destination (000000.00 - 246000.00*, null), null means unchanged
	EstimatedArrivalTime Int64 // 4

	// Estimated day of arrival at destination (00 - 31) (UTC), null means unchanged
	EstimatedArrivalDay Int64 // 5

	// Estimated month of arrival at destination (00 - 12) (UTC), null means unchanged
	EstimatedArrivalMonth Int64 // 6

	// Navigational status (0 - 15), null means unchanged. Reference ITU-R M.1371, Message 1, navigational status.
	// Source: https://www.itu.int/dms_pubrec/itu-r/rec/m/R-REC-M.1371-5-201402-I!!PDF-E.pdf (page 111)
	// null - unchanged
	// 0 - Under way using engine
	// 1 - At anchor
	// 2 - Not under command
	// 3 - Restricted maneuverability
	// 4 - Constrained by her draught
	// 5 - Moored
	// 6 - Aground
	// 7 - Engaged in Fishing
	// 8 - Under way sailing
	// 9 - HSC
	// 10 - WIG
	// 11 - Power-driven vessel towing astern
	// 12 - Power-driven vessel pushing ahead or towing alongside
	// 13 - Reserved for future use
	// 14 - AIS-SART (active), MOB-AIS, EPIRB-AIS
	// 15 - Undefined = default (also used by AIS-SART, MOB-AIS and EPIRB AIS under test)
	NavigationalStatus Int64 // 7

	// RegionalApplication is Regional application flags, null means unchanged. Reference ITU-R M.1371, Message 1, reserved for regional applications.
	RegionalApplication Int64 // 8
}

// newVSD constructor
func newVSD(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeVSD)
	m := VSD{
		BaseSentence:          s,
		TypeOfShipAndCargo:    p.NullInt64(0, "type of ship and cargo"),
		StaticDraughtMeters:   p.NullFloat64(1, "maximum present static draught"),
		PersonsOnBoard:        p.NullInt64(2, "persons on-board"),
		Destination:           p.String(3, "destination"),
		EstimatedArrivalTime:  p.NullInt64(4, "estimated arrival time"),
		EstimatedArrivalDay:   p.NullInt64(5, "estimated arrival day"),
		EstimatedArrivalMonth: p.NullInt64(6, "estimated arrival month"),
		NavigationalStatus:    p.NullInt64(7, "navigational status"),
		RegionalApplication:   p.NullInt64(8, "Regional application"),
	}
	return m, p.Err()
}
