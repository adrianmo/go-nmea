package nmea

import "errors"

const (
	// TypeDSE type of DSE sentence for Expanded digital selective calling
	TypeDSE = "DSE"

	// AcknowledgementAutomaticDSE is type for automatic
	AcknowledgementAutomaticDSE = "A"
	// AcknowledgementRequestDSE is type for request
	AcknowledgementRequestDSE = "R"
	// AcknowledgementQueryDSE is type for query
	AcknowledgementQueryDSE = "Q"
)

// DSE – Expanded digital selective calling. Is sentence that follows DSC sentence to provide additional (extended) data.
// https://opencpn.org/wiki/dokuwiki/doku.php?id=opencpn:opencpn_user_manual:advanced_features:nmea_sentences
// http://www.busse-yachtshop.de/pdf/icom-GM600-handbuch.pdf
//
// Format: $CDDSE, x, x, a, xxxxxxxxxx, xx, c--c, .........., xx, c--c*hh<CR><LF>
// Example: $CDDSE,1,1,A,3380400790,00,46504437*15
type DSE struct {
	BaseSentence
	TotalNumber     int64  // total number of sentences, 01 to 99
	Number          int64  // number of current sentence, 01 to 99
	Acknowledgement string // Acknowledgement (R=Acknowledge request, B=Acknowledgement, S=Neither (end of sequence))
	MMSI            string // MMSI of vessel (10 digits)
	DataSets        []DSEDataSet
}

// DSEDataSet is pair of DSE sets of data containing code + its data
type DSEDataSet struct {
	// Code is code field, 2 digits
	// From OpenCPN wiki:
	// > 00–this field of two-digits appears to be the expansion data specifier described in Table 1 of ITU-Rec.M821-1,
	// > but with the symbol representation in two-digits instead of three-digits. The leading “1” seems to not be used.
	// > (See modified table, above.) This field identifies the data that will follow in the next field. In this message,
	// > the data will be “enhanced position resolution.”
	Code string
	// Data is data field, Enhanced position resolution, Maximum 8 characters, could be empty
	// From OpenCPN wiki:
	// > 45894494–the data payload, which is eight digits. The first four are the decimal portion of the latitude
	// > minutes; the last four are the decimal portion of the longitude minutes. The latitude and longitude whole
	// > minutes were sent in the immediately preceding datagram. This is as specified in the ITU-Rec. M.821-1 in
	// > section 2.1.2.1
	Data string
}

// newDSE constructor
func newDSE(s BaseSentence) (Sentence, error) {
	p := NewParser(s)
	p.AssertType(TypeDSE)
	dse := DSE{
		BaseSentence:    s,
		TotalNumber:     p.Int64(0, "total number of sentences"),
		Number:          p.Int64(1, "sentence number"),
		Acknowledgement: p.EnumString(2, "acknowledgement", AcknowledgementAutomaticDSE, AcknowledgementRequestDSE, AcknowledgementQueryDSE),
		MMSI:            p.String(3, "MMSI"),
		DataSets:        nil,
	}
	datasetFieldCount := len(p.Fields) - 4
	if datasetFieldCount < 2 {
		return dse, errors.New("DSE is missing fields for parsing data sets")
	}
	if datasetFieldCount%2 != 0 {
		return dse, errors.New("DSE data set field count is not exactly dividable by 2")
	}
	dse.DataSets = make([]DSEDataSet, 0, datasetFieldCount/2)
	for i := 0; i < datasetFieldCount; i = i + 2 {
		tmp := DSEDataSet{
			Code: p.String(4+i, "data set code"),
			Data: p.String(5+i, "data set data"),
		}
		dse.DataSets = append(dse.DataSets, tmp)
	}
	return dse, p.Err()
}
