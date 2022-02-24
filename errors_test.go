package nmea

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseError_As(t *testing.T) {
	_, err := Parse("$HEROT,-11.23,X*1E")

	assert.Equal(t, &ParseError{Errors: []error{&InvalidEnumStringError{
		Prefix:  "HEROT",
		Context: "status valid",
		Value:   "X",
	}}}, err)

	var expect *InvalidEnumStringError
	as := errors.As(err, &expect)
	assert.True(t, as)
	assert.Equal(t, "HEROT", expect.Prefix)
	assert.Equal(t, "status valid", expect.Context)
	assert.Equal(t, "X", expect.Value)

	var expect2 *NotSupportedError
	asFalse := errors.As(err, &expect2)
	assert.False(t, asFalse)
}

func TestParseError_Is(t *testing.T) {
	_, err := Parse("$HEROT,-11.23,X*1E")

	errInvalidEnum := &InvalidEnumStringError{
		Prefix:  "HEROT",
		Context: "status valid",
		Value:   "X",
	}
	assert.Equal(t, &ParseError{Errors: []error{errInvalidEnum}}, err)

	isTrue := errors.Is(err, errInvalidEnum)
	assert.False(t, isTrue)

	var expectNot *NotSupportedError
	isFalse := errors.Is(err, expectNot)
	assert.False(t, isFalse)
}

func TestParseError_WithMultipleErrors(t *testing.T) {
	m, err := Parse("$HEROT,nope,X*08")

	assert.Equal(t, ROT{
		BaseSentence: BaseSentence{
			Talker:   "HE",
			Type:     "ROT",
			Fields:   []string{"nope", "X"},
			Checksum: "08",
			Raw:      "$HEROT,nope,X*08",
			TagBlock: TagBlock{},
		},
		RateOfTurn: 0,
		Valid:      false,
	}, m)

	assert.Equal(t, &ParseError{
		Errors: []error{
			&FieldError{
				Prefix:  "HEROT",
				Context: "rate of turn",
				Value:   "nope",
			},
			&InvalidEnumStringError{
				Prefix:  "HEROT",
				Context: "status valid",
				Value:   "X",
			},
		},
	}, err)

	var expect *InvalidEnumStringError
	assert.True(t, errors.As(err, &expect))
	assert.Equal(t,
		&InvalidEnumStringError{
			Prefix:  "HEROT",
			Context: "status valid",
			Value:   "X",
		},
		expect,
	)

	assert.True(t, errors.Is(err, errors.New("nmea: HEROT invalid rate of turn: nope")))
	assert.True(t, errors.Is(err, &InvalidEnumStringError{
		Prefix:  "HEROT",
		Context: "rate of turn",
		Value:   "nope",
	}))
}

func TestFieldError_Is(t *testing.T) {
	var testCases = []struct {
		name   string
		when   error
		expect bool
	}{
		{
			name: "is same",
			when: &FieldError{
				Prefix:  "HEROT",
				Context: "rate of turn",
				Value:   "nope",
			},
			expect: true,
		},
		{
			name: "is not exactly same",
			when: &FieldError{
				Prefix:  "XXX",
				Context: "rate of turn",
				Value:   "nope",
			},
			expect: false,
		},
		{
			name:   "is not same",
			when:   &NotSupportedError{Prefix: "HEROT"},
			expect: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := &FieldError{
				Prefix:  "HEROT",
				Context: "rate of turn",
				Value:   "nope",
			}
			is := errors.Is(err, tc.when)
			assert.Equal(t, tc.expect, is)
		})
	}
}
