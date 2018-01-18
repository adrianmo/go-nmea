package nmea

import (
	"fmt"
	"strconv"
)

type parser struct {
	Sentence
	prefix string
	err    error
}

func newParser(s Sentence, prefix string) *parser {
	p := &parser{Sentence: s, prefix: prefix}
	if p.Type != prefix {
		p.SetErr("prefix", p.Type)
	}
	return p
}

func (p *parser) Err() error {
	return p.err
}

func (p *parser) SetErr(context, value string) {
	if p.err == nil {
		p.err = fmt.Errorf("%s invalid %s: %s", p.prefix, context, value)
	}
}

func (p *parser) Empty(i int, context string) bool {
	return p.String(i, context) == ""
}

func (p *parser) String(i int, context string) string {
	if p.err != nil {
		return ""
	}
	if i < 0 || i >= len(p.Fields) {
		p.SetErr(context, strconv.Itoa(i))
	}
	return p.Fields[i]
}

func (p *parser) EnumString(i int, context string, options ...string) string {
	s := p.String(i, context)
	if p.err != nil {
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

func (p *parser) Int64(i int, context string) int64 {
	s := p.String(i, context)
	if p.err != nil {
		return 0
	}
	if s == "" {
		return 0
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		p.SetErr(context, s)
	}
	return v
}

func (p *parser) Float64(i int, context string) float64 {
	s := p.String(i, context)
	if p.err != nil {
		return 0
	}
	if s == "" {
		return 0
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		p.SetErr(context, s)
	}
	return v
}

func (p *parser) Time(i int, context string) Time {
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

func (p *parser) LatLong(i, j int, context string) LatLong {
	a := p.String(i, context)
	b := p.String(j, context)
	if p.err != nil {
		return 0
	}
	s := fmt.Sprintf("%s %s", a, b)
	v, err := NewLatLong(s)
	if err != nil {
		p.SetErr(context, err.Error())
	}
	return v
}
