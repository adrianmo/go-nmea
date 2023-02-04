package nmea

import (
	"fmt"
	"strconv"
)

// Parser provides a simple way of accessing and parsing
// sentence fields
type Parser struct {
	BaseSentence
	err error
}

// NewParser constructor
func NewParser(s BaseSentence) *Parser {
	return &Parser{BaseSentence: s}
}

// AssertType makes sure the sentence's type matches the provided one.
func (p *Parser) AssertType(typ string) {
	if p.Type != typ {
		p.SetErr("type", p.Type)
	}
}

// Err returns the first error encountered during the parser's usage.
func (p *Parser) Err() error {
	return p.err
}

// SetErr assigns an error. Calling this method has no
// effect if there is already an error.
func (p *Parser) SetErr(context, value string) {
	if p.err == nil {
		p.err = fmt.Errorf("nmea: %s invalid %s: %s", p.Prefix(), context, value)
	}
}

// String returns the field value at the specified index.
func (p *Parser) String(i int, context string) string {
	if p.err != nil {
		return ""
	}
	if i < 0 || i >= len(p.Fields) {
		p.SetErr(context, "index out of range")
		return ""
	}
	return p.Fields[i]
}

// ListString returns a list of all fields from the given start index.
// An error occurs if there is no fields after the given start index.
func (p *Parser) ListString(from int, context string) (list []string) {
	if p.err != nil {
		return []string{}
	}
	if from < 0 || from >= len(p.Fields) {
		p.SetErr(context, "index out of range")
		return []string{}
	}
	return append(list, p.Fields[from:]...)
}

// EnumString returns the field value at the specified index.
// An error occurs if the value is not one of the options and not empty.
func (p *Parser) EnumString(i int, context string, options ...string) string {
	s := p.String(i, context)
	if p.err != nil || s == "" {
		return ""
	}
	for _, o := range options {
		if o == s {
			return s
		}
	}
	p.SetErr(context, s)
	return ""
}

// EnumChars returns an array of strings that are matched in the Mode field.
// It will only match the number of characters that are in the Mode field.
// If the value is empty, it will return an empty array
func (p *Parser) EnumChars(i int, context string, options ...string) []string {
	s := p.String(i, context)
	if p.err != nil || s == "" {
		return []string{}
	}
	strs := []string{}
	for _, r := range s {
		rs := string(r)
		for _, o := range options {
			if o == rs {
				strs = append(strs, o)
				break
			}
		}
	}
	if len(strs) != len(s) {

		p.SetErr(context, s)
		return []string{}
	}
	return strs
}

// HexInt64 returns the hex encoded int64 value at the specified index.
// If the value is an empty string, 0 is returned.
func (p *Parser) HexInt64(i int, context string) int64 {
	s := p.String(i, context)
	if p.err != nil {
		return 0
	}
	if s == "" {
		return 0
	}
	value, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		p.SetErr(context, s)
	}
	return value
}

// Int64 returns the int64 value at the specified index.
// If the value is an empty string, 0 is returned.
func (p *Parser) Int64(i int, context string) int64 {
	return p.NullInt64(i, context).Value
}

// NullInt64 returns the int64 value at the specified index.
// If the value is an empty string, Valid is set to false
func (p *Parser) NullInt64(i int, context string) Int64 {
	s := p.String(i, context)
	if p.err != nil {
		return Int64{}
	}
	if s == "" {
		return Int64{}
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		p.SetErr(context, s)
		return Int64{}
	}
	return Int64{Value: v, Valid: true}
}

// Float64 returns the float64 value at the specified index.
// If the value is an empty string, 0 is returned.
func (p *Parser) Float64(i int, context string) float64 {
	return p.NullFloat64(i, context).Value
}

// NullFloat64 returns the Float64 value at the specified index.
// If the value is an empty string, Valid is set to false.
func (p *Parser) NullFloat64(i int, context string) Float64 {
	s := p.String(i, context)
	if p.err != nil {
		return Float64{}
	}
	if s == "" {
		return Float64{}
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		p.SetErr(context, s)
		return Float64{}
	}
	return Float64{Value: v, Valid: true}
}

// Time returns the Time value at the specified index.
// If the value is empty, the Time is marked as invalid.
func (p *Parser) Time(i int, context string) Time {
	s := p.String(i, context)
	if p.err != nil {
		return Time{}
	}
	v, err := ParseTime(s)
	if err != nil {
		p.SetErr(context, s)
	}
	return v
}

// Date returns the Date value at the specified index.
// If the value is empty, the Date is marked as invalid.
func (p *Parser) Date(i int, context string) Date {
	s := p.String(i, context)
	if p.err != nil {
		return Date{}
	}
	v, err := ParseDate(s)
	if err != nil {
		p.SetErr(context, s)
	}
	return v
}

// LatLong returns the coordinate value of the specified fields.
func (p *Parser) LatLong(i, j int, context string) float64 {
	a := p.String(i, context)
	b := p.String(j, context)
	if p.err != nil {
		return 0
	}
	s := fmt.Sprintf("%s %s", a, b)
	v, err := ParseLatLong(s)
	if err != nil {
		p.SetErr(context, err.Error())
	}

	if (b == North || b == South) && (v < -90.0 || 90.0 < v) {
		p.SetErr(context, "latitude is not in range (-90, 90)")
		return 0
	} else if (b == West || b == East) && (v < -180.0 || 180.0 < v) {
		p.SetErr(context, "longitude is not in range (-180, 180)")
		return 0
	}

	return v
}

// SixBitASCIIArmour decodes the 6-bit ascii armor used for VDM and VDO messages
func (p *Parser) SixBitASCIIArmour(i int, fillBits int, context string) []byte {
	if p.err != nil {
		return nil
	}
	if fillBits < 0 || fillBits >= 6 {
		p.SetErr(context, "fill bits")
		return nil
	}

	payload := []byte(p.String(i, "encoded payload"))
	numBits := len(payload)*6 - fillBits

	if numBits < 0 {
		p.SetErr(context, "num bits")
		return nil
	}

	result := make([]byte, numBits)
	resultIndex := 0

	for _, v := range payload {
		if v < 48 || v >= 120 {
			p.SetErr(context, "data byte")
			return nil
		}

		d := v - 48
		if d > 40 {
			d -= 8
		}

		for i := 5; i >= 0 && resultIndex < len(result); i-- {
			result[resultIndex] = (d >> uint(i)) & 1
			resultIndex++
		}
	}

	return result
}
