package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var parsertests = []struct {
	name     string
	fields   []string
	expected interface{}
	hasErr   bool
	parse    func(p *Parser) interface{}
}{
	{
		name:   "Bad Type",
		fields: []string{},
		hasErr: true,
		parse: func(p *Parser) interface{} {
			p.AssertType("WRONG_TYPE")
			return nil
		},
	},
	{
		name:     "String",
		fields:   []string{"foo", "bar"},
		expected: "bar",
		parse: func(p *Parser) interface{} {
			return p.String(1, "")
		},
	},
	{
		name:     "String out of range",
		fields:   []string{"wot"},
		expected: "",
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			return p.String(5, "thing")
		},
	},
	{
		name:     "ListString",
		fields:   []string{"wot", "foo", "bar"},
		expected: []string{"foo", "bar"},
		parse: func(p *Parser) interface{} {
			return p.ListString(1, "thing")
		},
	},
	{
		name:     "ListString out of range",
		fields:   []string{"wot"},
		expected: []string{},
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			return p.ListString(10, "thing")
		},
	},
	{
		name:     "String with existing error",
		expected: "",
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			p.SetErr("context", "value")
			return p.String(123, "blah")
		},
	},
	{
		name:     "EnumString",
		fields:   []string{"a", "b", "c"},
		expected: "b",
		parse: func(p *Parser) interface{} {
			return p.EnumString(1, "context", "b", "d")
		},
	},
	{
		name:     "EnumString invalid",
		fields:   []string{"a", "b", "c"},
		expected: "",
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			return p.EnumString(1, "context", "x", "y")
		},
	},
	{
		name:     "EnumString with existing error",
		fields:   []string{"a", "b", "c"},
		expected: "",
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			p.SetErr("context", "value")
			return p.EnumString(1, "context", "a", "b")
		},
	},
	{
		name:     "EnumChars",
		fields:   []string{"AA", "AB", "BA", "BB"},
		expected: []string{"A", "B"},
		parse: func(p *Parser) interface{} {
			return p.EnumChars(1, "context", "A", "B")
		},
	},
	{
		name:     "EnumChars invalid",
		fields:   []string{"a", "AB", "c"},
		expected: []string{},
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			return p.EnumChars(1, "context", "X", "Y")
		},
	},
	{
		name:     "EnumChars with existing error",
		fields:   []string{"a", "AB", "c"},
		expected: []string{},
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			p.SetErr("context", "value")
			return p.EnumChars(1, "context", "A", "B")
		},
	},
	{
		name:     "Int64",
		fields:   []string{"123"},
		expected: int64(123),
		parse: func(p *Parser) interface{} {
			return p.Int64(0, "context")
		},
	},
	{
		name:     "Int64 empty field is zero",
		fields:   []string{""},
		expected: int64(0),
		parse: func(p *Parser) interface{} {
			return p.Int64(0, "context")
		},
	},
	{
		name:     "Int64 invalid",
		fields:   []string{"abc"},
		expected: int64(0),
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			return p.Int64(0, "context")
		},
	},
	{
		name:     "Int64 with existing error",
		fields:   []string{"123"},
		expected: int64(0),
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			p.SetErr("context", "value")
			return p.Int64(0, "context")
		},
	},
	{
		name:     "Float64",
		fields:   []string{"123.123"},
		expected: float64(123.123),
		parse: func(p *Parser) interface{} {
			return p.Float64(0, "context")
		},
	},
	{
		name:     "Float64 empty field is zero",
		fields:   []string{""},
		expected: float64(0),
		parse: func(p *Parser) interface{} {
			return p.Float64(0, "context")
		},
	},
	{
		name:     "Float64 invalid",
		fields:   []string{"abc"},
		expected: float64(0),
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			return p.Float64(0, "context")
		},
	},
	{
		name:     "Float64 with existing error",
		fields:   []string{"123.123"},
		expected: float64(0),
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			p.SetErr("context", "value")
			return p.Float64(0, "context")
		},
	},
	{
		name:     "NullInt64",
		fields:   []string{"123"},
		expected: Int64{Value: 123, Valid: true},
		parse: func(p *Parser) interface{} {
			return p.NullInt64(0, "context")
		},
	},
	{
		name:     "NullInt64 empty field is invalid",
		fields:   []string{""},
		expected: Int64{},
		parse: func(p *Parser) interface{} {
			return p.NullInt64(0, "context")
		},
	},
	{
		name:     "NullInt64 invalid",
		fields:   []string{"abc"},
		expected: Int64{},
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			return p.NullInt64(0, "context")
		},
	},
	{
		name:     "NullInt64 with existing error",
		fields:   []string{"123"},
		expected: Int64{},
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			p.SetErr("context", "value")
			return p.NullInt64(0, "context")
		},
	},
	{
		name:     "NullFloat64",
		fields:   []string{"123.123"},
		expected: Float64{Value: 123.123, Valid: true},
		parse: func(p *Parser) interface{} {
			return p.NullFloat64(0, "context")
		},
	},
	{
		name:     "NullFloat64 empty field is invalid",
		fields:   []string{""},
		expected: Float64{},
		parse: func(p *Parser) interface{} {
			return p.NullFloat64(0, "context")
		},
	},
	{
		name:     "NullFloat64 invalid",
		fields:   []string{"abc"},
		expected: Float64{},
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			return p.NullFloat64(0, "context")
		},
	},
	{
		name:     "NullFloat64 with existing error",
		fields:   []string{"123.123"},
		expected: Float64{},
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			p.SetErr("context", "value")
			return p.NullFloat64(0, "context")
		},
	},
	{
		name:     "Time",
		fields:   []string{"123456"},
		expected: Time{true, 12, 34, 56, 0},
		parse: func(p *Parser) interface{} {
			return p.Time(0, "context")
		},
	},
	{
		name:     "Time empty field is zero",
		fields:   []string{""},
		expected: Time{},
		parse: func(p *Parser) interface{} {
			return p.Time(0, "context")
		},
	},
	{
		name:     "Time with existing error",
		fields:   []string{"123456"},
		expected: Time{},
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			p.SetErr("context", "value")
			return p.Time(0, "context")
		},
	},
	{
		name:     "Time invalid",
		fields:   []string{"wrong"},
		expected: Time{},
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			return p.Time(0, "context")
		},
	},
	{
		name:     "Date",
		fields:   []string{"010203"},
		expected: Date{true, 1, 2, 3},
		parse: func(p *Parser) interface{} {
			return p.Date(0, "context")
		},
	},
	{
		name:     "Date empty field is zero",
		fields:   []string{""},
		expected: Date{},
		parse: func(p *Parser) interface{} {
			return p.Date(0, "context")
		},
	},
	{
		name:     "Date invalid",
		fields:   []string{"Hello"},
		expected: Date{},
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			return p.Date(0, "context")
		},
	},
	{
		name:     "Date with existing error",
		fields:   []string{"010203"},
		expected: Date{},
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			p.SetErr("context", "value")
			return p.Date(0, "context")
		},
	},
	{
		name:     "LatLong",
		fields:   []string{"5000.0000", "N"},
		expected: 50.0,
		parse: func(p *Parser) interface{} {
			return p.LatLong(0, 1, "context")
		},
	},
	{
		name:     "LatLong - latitude out of range",
		fields:   []string{"9100.0000", "N"},
		expected: 0.0,
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			return p.LatLong(0, 1, "context")
		},
	},
	{
		name:     "LatLong - longitude out of range",
		fields:   []string{"18100.0000", "W"},
		expected: 0.0,
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			return p.LatLong(0, 1, "context")
		},
	},
	{
		name:     "LatLong with existing error",
		fields:   []string{"5000.0000", "W"},
		expected: 0.0,
		hasErr:   true,
		parse: func(p *Parser) interface{} {
			p.SetErr("context", "value")
			return p.LatLong(0, 1, "context")
		},
	},
}

func TestParser(t *testing.T) {
	for _, tt := range parsertests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(BaseSentence{
				Talker: "talker",
				Type:   "type",
				Fields: tt.fields,
			})
			assert.Equal(t, tt.expected, tt.parse(p))
			if tt.hasErr {
				assert.Error(t, p.Err())
			} else {
				assert.NoError(t, p.Err())
			}
		})
	}
}
