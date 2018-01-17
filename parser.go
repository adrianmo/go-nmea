package nmea

import (
	"fmt"
	"strconv"
)

type parser struct {
	Sentence
	context string
	err     error
}

func newParser(s Sentence) *parser {
	return &parser{Sentence: s}
}

func (p *parser) Err() error {
	if p.err != nil {
		return fmt.Errorf("%s %s", p.context, p.err)
	}
	return nil
}

func (p *parser) String(i int, context string) string {
	if p.err != nil {
		return ""
	}
	if i < 0 || i >= len(p.Fields) {
		p.context = context
		p.err = fmt.Errorf("index out of range %d", i)
	}
	return p.Fields[i]
}

func (p *parser) Int64(i int, context string) int64 {
	s := p.String(i, context)
	if p.err != nil {
		return 0
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		p.context = context
		p.err = err
	}
	return v
}

func (p *parser) Float64(i int, context string) float64 {
	s := p.String(i, context)
	if p.err != nil {
		return 0
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		p.context = context
		p.err = err
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
		p.context = context
		p.err = err
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
		p.context = context
		p.err = err
	}
	return v
}
