package nmea

import "strings"

const (
	// TypeDSC type of DSC sentence for Digital Selective Calling Information
	TypeDSC = "DSC"

	// AcknowledgementRequestDSC is type for Acknowledge request
	AcknowledgementRequestDSC = "R"
	// AcknowledgementDSC is type for Acknowledgement
	AcknowledgementDSC = "B"
	// AcknowledgementNeitherDSC is type for Neither (end of sequence)
	AcknowledgementNeitherDSC = "S"
)

// DSC – Digital Selective Calling Information
// https://opencpn.org/wiki/dokuwiki/doku.php?id=opencpn:opencpn_user_manual:advanced_features:nmea_sentences
// https://web.archive.org/web/20190303170916/http://continuouswave.com/whaler/reference/DSC_Datagrams.html
// http://www.busse-yachtshop.de/pdf/icom-GM600-handbuch.pdf
// https://github.com/mariokonrad/marnav/blob/master/src/marnav/nmea/dsc.cpp (marnav has interesting enums worth checking)
//
// Note: many fields of DSC are conditional with double meaning and we only map raw sentence to fields without any
// logic/checking of those conditions. We could have specific fields if we only knew the rules to populate them.
//
// Format: $--DSC,xx,xxxxxxxxxx,xx,xx,xx,x.x, x.x,xxxxxxxxxx,xx, a,a*hh<CR><LF>
// Example: $CDDSC,20,3380400790,00,21,26,1423108312,2021,,,B, E*73
type DSC struct {
	BaseSentence
	// Note: all fields are strings even if specified as digits as int can not express "00" and would be 0 which is different
	// Source of quotes: https://web.archive.org/web/20190303170916/http://continuouswave.com/whaler/reference/DSC_Datagrams.html

	// FormatSpecifier is Format specifier (2 digits)
	// > The call content is first described by a "format specifier" element. The format specifier is explained in
	// > ITU-Rec. M.493-13 Section 4, with various symbol codes in the "service command" range of symbols representing
	// > various message formats, as shown in Table 3 (by symbol number, then meaning of symbol) as follows:
	// > * 102 = selective call to a group of ships in particular geographic area
	// > * 112 = distress alert call
	// > * 114 = selective call to a group of ships having common interest
	// > * 116 = all ships call
	// > * 120 = selective call to particular individual station
	// > * 123 = selective call to a particular individual using automatic service
	FormatSpecifier string

	// Address (10 digits)
	Address string

	// Category (2 digits or empty)
	// > The call content is next described by a "category element" in Section 6. Again, various symbol codes in the
	// > "service command" range of symbols represent various categories, as follows from Table 3 (by symbol number,
	// > then meaning of symbol):
	// > * 100 = routine
	// > * 108 = safety
	// > * 110 = urgency
	// > * 112 = distress
	Category string

	// DistressCauseOrTeleCommand1 is The cause of the distress or first telecommand (2 digits or empty)
	// > Nature of Distress is to be encoded, again using Table 3, as follows
	// > * 100 = Fire, explosion
	// > * 101 = Flooding
	// > * 102 = Collision
	// > * 103 = Grounding
	// > * 104 = Listing, in danger of capsize
	// > * 105 = Sinking
	// > * 106 = Disabled and adrift
	// > * 107 = Undesignated distres
	// > * 108 = Abandoning ship
	// > * 109 = Piracy/armed robbery attack
	// > * 110 = Man overboard
	// > * 111 = unassigned symbol; take no action
	// > * 112 = EPRIB emission
	// > * 113 through 27 = unassigned symbol; take no action
	DistressCauseOrTeleCommand1 string

	// CommandTypeOrTeleCommand2 is Type of communication or second telecommand (2 digits)
	CommandTypeOrTeleCommand2 string

	// PositionOrCanal is Position (lat+lon) or Canal/frequency (Maximum 16 digits)
	// > Distress coordinates are to be encoded five parts, sent as a string of ten digits. The first digit indicates
	// > the direction of the latitude and longitude, with "0" for North and East, "1" for North and West,
	// > "2" for South and East, and "3" for South and West. The next two digits are the latitude in degrees.
	// > The next two digits are the latitude in whole minutes. The next three digits are the longitude in degrees.
	// > The next two digits are longitude in whole minutes.
	PositionOrCanal string // Position (lat+lon) or Canal/frequency (Maximum 16 digits)

	// TimeOrTelephoneNumber is Time or Telephone Number (Maximum 16 digits)
	// > The time in universal coordinated time is to be sent in 24-hour format in two parts, a total of four digits.
	// > The first two digits are the hours. The next two are the minutes.
	TimeOrTelephoneNumber string

	// MMSI of ship in distress (10 digits or empty)
	// > The call content is next described as having a "self-identification" element. This is simply the sending
	// > station's MMSI, encoded like the address element. This identifies who sent the message.
	MMSI string

	// DistressCause is The cause of the distress (2 digits or empty)
	DistressCause string

	// Acknowledgement (R=Acknowledge request, B=Acknowledgement, S=Neither (end of sequence))
	Acknowledgement string

	// Expansion indicator (E or empty)
	ExpansionIndicator string
}

// newDSC constructor
func newDSC(s BaseSentence) (DSC, error) {
	p := NewParser(s)
	p.AssertType(TypeDSC)
	return DSC{
		BaseSentence:                s,
		FormatSpecifier:             p.String(0, "format specifier"),
		Address:                     p.String(1, "address"),
		Category:                    p.String(2, "category"),
		DistressCauseOrTeleCommand1: p.String(3, "cause of the distress or first telecommand"),
		CommandTypeOrTeleCommand2:   p.String(4, "type of communication or second telecommand"),
		PositionOrCanal:             p.String(5, "position or canal"),
		TimeOrTelephoneNumber:       p.String(6, "time or telephone"),
		MMSI:                        p.String(7, "MMSI"),
		DistressCause:               p.String(8, "distress cause"),
		Acknowledgement: strings.TrimSpace(p.EnumString(
			9,
			"acknowledgement",
			AcknowledgementRequestDSC,
			" "+AcknowledgementRequestDSC,
			AcknowledgementDSC,
			" "+AcknowledgementDSC,
			AcknowledgementNeitherDSC,
			" "+AcknowledgementNeitherDSC,
		)),
		ExpansionIndicator: p.String(10, "expansion indicator"),
	}, p.Err()
}
