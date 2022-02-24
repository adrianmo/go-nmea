package nmea

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_AssertType(t *testing.T) {
	var testCases = []struct {
		name        string
		givenType   string
		whenType    string
		expectError string
	}{
		{
			name:        "ok type",
			givenType:   "RMC",
			whenType:    "RMC",
			expectError: "",
		},
		{
			name:        "bad type",
			givenType:   "RMC",
			whenType:    "RMB",
			expectError: "nmea: talkerRMC invalid type: RMC",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewParser(BaseSentence{
				Talker: "talker",
				Type:   tc.givenType,
			})

			p.AssertType(tc.whenType)

			err := p.Err()
			if tc.expectError != "" {
				assert.EqualError(t, err, tc.expectError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParser_String(t *testing.T) {
	var testCases = []struct {
		name        string
		givenFields []string
		givenError  error
		expect      string
		expectError error
	}{
		{
			name:        "ok, string",
			givenFields: []string{"foo", "bar"},
			expect:      "bar",
		},
		{
			name:        "nok, index out of range",
			givenFields: []string{"wot"},
			expect:      "",
			expectError: &ParseError{Errors: []error{errors.New("nmea: talkertype invalid context: index out of range")}},
		},
		{
			name:        "ok, string + already existing error",
			givenFields: []string{"foo", "bar"},
			givenError:  errors.New("existing"),
			expect:      "bar",
			expectError: &ParseError{Errors: []error{errors.New("existing")}},
		},
		{
			name:        "nok, string + already existing error",
			givenFields: []string{"bar"},
			givenError:  errors.New("existing"),
			expect:      "",
			expectError: &ParseError{Errors: []error{
				errors.New("existing"),
				errors.New("nmea: talkertype invalid context: index out of range"),
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewParser(BaseSentence{
				Talker: "talker",
				Type:   "type",
				Fields: tc.givenFields,
			})
			if tc.givenError != nil {
				p.setError(tc.givenError)
			}

			result := p.String(1, "context")

			assert.Equal(t, tc.expect, result)
			assert.Equal(t, tc.expectError, p.Err())
		})
	}
}

func TestParser_ListString(t *testing.T) {
	var testCases = []struct {
		name        string
		givenFields []string
		givenError  error
		expect      []string
		expectError error
	}{
		{
			name:        "ok, ListString",
			givenFields: []string{"wot", "foo", "bar"},
			expect:      []string{"foo", "bar"},
		},
		{
			name:        "nok, index out of range",
			givenFields: []string{"wot"},
			expect:      []string{},
			expectError: &ParseError{Errors: []error{errors.New("nmea: talkertype invalid context: index out of range")}},
		},
		{
			name:        "ok, ListString + already existing error",
			givenFields: []string{"foo", "bar"},
			givenError:  errors.New("existing"),
			expect:      []string{"bar"},
			expectError: &ParseError{Errors: []error{errors.New("existing")}},
		},
		{
			name:        "nok, ListString + already existing error",
			givenFields: []string{"bar"},
			givenError:  errors.New("existing"),
			expect:      []string{},
			expectError: &ParseError{Errors: []error{
				errors.New("existing"),
				errors.New("nmea: talkertype invalid context: index out of range"),
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewParser(BaseSentence{
				Talker: "talker",
				Type:   "type",
				Fields: tc.givenFields,
			})
			if tc.givenError != nil {
				p.setError(tc.givenError)
			}

			result := p.ListString(1, "context")

			assert.Equal(t, tc.expect, result)
			assert.Equal(t, tc.expectError, p.Err())
		})
	}
}

func TestParser_EnumString(t *testing.T) {
	var testCases = []struct {
		name        string
		givenFields []string
		givenError  error
		whenOptions []string
		expect      string
		expectError error
	}{
		{
			name:        "ok, EnumString",
			givenFields: []string{"180", HeadingMagnetic},
			whenOptions: []string{HeadingTrue, HeadingMagnetic},
			expect:      HeadingMagnetic,
		},
		{
			name:        "nok, index out of range",
			givenFields: []string{"180"},
			whenOptions: []string{HeadingTrue, HeadingMagnetic},
			expect:      "",
			expectError: &ParseError{Errors: []error{errors.New("nmea: talkertype invalid context: index out of range")}},
		},
		{
			name:        "ok, EnumString + already existing error",
			givenFields: []string{"180", HeadingMagnetic},
			givenError:  errors.New("existing"),
			whenOptions: []string{HeadingTrue, HeadingMagnetic},
			expect:      HeadingMagnetic,
			expectError: &ParseError{Errors: []error{errors.New("existing")}},
		},
		{
			name:        "nok, EnumString + already existing error",
			givenFields: []string{"180"},
			givenError:  errors.New("existing"),
			whenOptions: []string{HeadingTrue, HeadingMagnetic},
			expect:      "",
			expectError: &ParseError{Errors: []error{
				errors.New("existing"),
				errors.New("nmea: talkertype invalid context: index out of range"),
			}},
		},
		{
			name:        "nok, invalid option + already existing error",
			givenFields: []string{"180", HeadingMagnetic},
			givenError:  errors.New("existing"),
			whenOptions: []string{Left, Right},
			expect:      HeadingMagnetic,
			expectError: &ParseError{Errors: []error{
				errors.New("existing"),
				&InvalidEnumStringError{Prefix: "talkertype", Context: "context", Value: "M"},
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewParser(BaseSentence{
				Talker: "talker",
				Type:   "type",
				Fields: tc.givenFields,
			})
			if tc.givenError != nil {
				p.setError(tc.givenError)
			}

			result := p.EnumString(1, "context", tc.whenOptions...)

			assert.Equal(t, tc.expect, result)
			assert.Equal(t, tc.expectError, p.Err())
		})
	}
}

func TestParser_EnumChars(t *testing.T) {
	var testCases = []struct {
		name        string
		givenFields []string
		givenError  error
		whenOptions []string
		expect      []string
		expectError error
	}{
		{
			name:        "ok, EnumChars",
			givenFields: []string{"AA", "AB", "BA", "BB"},
			whenOptions: []string{"A", "B"},
			expect:      []string{"A", "B"},
		},
		{
			name:        "nok, index out of range",
			givenFields: []string{"AA"},
			whenOptions: []string{"A", "B"},
			expect:      []string{},
			expectError: &ParseError{Errors: []error{errors.New("nmea: talkertype invalid context: index out of range")}},
		},
		{
			name:        "ok, EnumChars + already existing error",
			givenFields: []string{"AA", "AB", "BA", "BB"},
			givenError:  errors.New("existing"),
			whenOptions: []string{"A", "B"},
			expect:      []string{"A", "B"},
			expectError: &ParseError{Errors: []error{errors.New("existing")}},
		},
		{
			name:        "nok, EnumChars + already existing error",
			givenFields: []string{"AA"},
			givenError:  errors.New("existing"),
			whenOptions: []string{"A", "B"},
			expect:      []string{},
			expectError: &ParseError{Errors: []error{
				errors.New("existing"),
				errors.New("nmea: talkertype invalid context: index out of range"),
			}},
		},
		{
			name:        "nok, invalid option + already existing error",
			givenFields: []string{"AA", "AB"},
			givenError:  errors.New("existing"),
			whenOptions: []string{"A", "X"},
			expect:      []string{"A"},
			expectError: &ParseError{Errors: []error{
				errors.New("existing"),
				&InvalidEnumCharsError{Prefix: "talkertype", Context: "context", Value: "AB"},
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewParser(BaseSentence{
				Talker: "talker",
				Type:   "type",
				Fields: tc.givenFields,
			})
			if tc.givenError != nil {
				p.setError(tc.givenError)
			}

			result := p.EnumChars(1, "context", tc.whenOptions...)

			assert.Equal(t, tc.expect, result)
			assert.Equal(t, tc.expectError, p.Err())
		})
	}
}

func TestParser_Int64(t *testing.T) {
	var testCases = []struct {
		name          string
		givenFields   []string
		givenError    error
		expect        int64
		expectError   error
		expectErrorIs error
	}{
		{
			name:        "ok, Int64",
			givenFields: []string{"1", "2", "3"},
			expect:      2,
		},
		{
			name:        "ok, field is empty",
			givenFields: []string{"11", "", "22"},
			expect:      0,
		},
		{
			name:        "nok, index out of range",
			givenFields: []string{"1"},
			expect:      0,
			expectError: &ParseError{Errors: []error{errors.New("nmea: talkertype invalid context: index out of range")}},
		},
		{
			name:        "nok, can not be parsed to Int64",
			givenFields: []string{"1", "ABC"},
			expect:      0,
			expectError: &ParseError{
				Errors: []error{
					&FieldError{Prefix: "talkertype", Context: "context", Value: "ABC"},
				},
			},
			expectErrorIs: errors.New("nmea: talkertype invalid context: ABC"),
		},
		{
			name:        "ok, Int64 + already existing error",
			givenFields: []string{"1", "2"},
			givenError:  errors.New("existing"),
			expect:      2,
			expectError: &ParseError{Errors: []error{errors.New("existing")}},
		},
		{
			name:        "nok, Int64 + already existing error",
			givenFields: []string{"1"},
			givenError:  errors.New("existing"),
			expect:      0,
			expectError: &ParseError{Errors: []error{
				errors.New("existing"),
				errors.New("nmea: talkertype invalid context: index out of range"),
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewParser(BaseSentence{
				Talker: "talker",
				Type:   "type",
				Fields: tc.givenFields,
			})
			if tc.givenError != nil {
				p.setError(tc.givenError)
			}

			result := p.Int64(1, "context")

			err := p.Err()
			assert.Equal(t, tc.expectError, err)
			if tc.expectErrorIs != nil {
				assert.True(t, errors.Is(err, tc.expectErrorIs))
			}
			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestParser_Float64(t *testing.T) {
	var testCases = []struct {
		name          string
		givenFields   []string
		givenError    error
		expect        float64
		expectError   error
		expectErrorIs error
	}{
		{
			name:        "ok, Float64",
			givenFields: []string{"1.2", "2.3", "3.4"},
			expect:      2.3,
		},
		{
			name:        "ok, field is empty",
			givenFields: []string{"1.2", "", "3.4"},
			expect:      0,
		},
		{
			name:        "nok, index out of range",
			givenFields: []string{"1.2"},
			expect:      0,
			expectError: &ParseError{Errors: []error{errors.New("nmea: talkertype invalid context: index out of range")}},
		},
		{
			name:        "nok, can not be parsed to Float64",
			givenFields: []string{"1.2", "ABC"},
			expect:      0,
			expectError: &ParseError{
				Errors: []error{
					&FieldError{Prefix: "talkertype", Context: "context", Value: "ABC"},
				},
			},
			expectErrorIs: errors.New("nmea: talkertype invalid context: ABC"),
		},
		{
			name:        "ok, Float64 + already existing error",
			givenFields: []string{"1.2", "2.3"},
			givenError:  errors.New("existing"),
			expect:      2.3,
			expectError: &ParseError{Errors: []error{errors.New("existing")}},
		},
		{
			name:        "nok, Float64 + already existing error",
			givenFields: []string{"1.2"},
			givenError:  errors.New("existing"),
			expect:      0,
			expectError: &ParseError{Errors: []error{
				errors.New("existing"),
				errors.New("nmea: talkertype invalid context: index out of range"),
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewParser(BaseSentence{
				Talker: "talker",
				Type:   "type",
				Fields: tc.givenFields,
			})
			if tc.givenError != nil {
				p.setError(tc.givenError)
			}

			result := p.Float64(1, "context")

			assert.Equal(t, tc.expect, result)
			assert.Equal(t, tc.expectError, p.Err())
		})
	}
}

func TestParser_Time(t *testing.T) {
	var testCases = []struct {
		name          string
		givenFields   []string
		givenError    error
		expect        Time
		expectError   error
		expectErrorIs error
	}{
		{
			name:        "ok, Time",
			givenFields: []string{"x", "123456", "x"},
			expect:      Time{true, 12, 34, 56, 0},
		},
		{
			name:        "ok, field is empty",
			givenFields: []string{"1.2", "", "3.4"},
			expect:      Time{false, 0, 0, 0, 0},
		},
		{
			name:        "nok, index out of range",
			givenFields: []string{"1.2"},
			expect:      Time{false, 0, 0, 0, 0},
			expectError: &ParseError{Errors: []error{errors.New("nmea: talkertype invalid context: index out of range")}},
		},
		{
			name:        "nok, can not be parsed to Time",
			givenFields: []string{"1.2", "ABC"},
			expect:      Time{false, 0, 0, 0, 0},
			expectError: &ParseError{
				Errors: []error{
					&FieldError{Prefix: "talkertype", Context: "context", Value: "ABC"},
				},
			},
			expectErrorIs: errors.New("nmea: talkertype invalid context: ABC"),
		},
		{
			name:        "ok, Time + already existing error",
			givenFields: []string{"x", "123456", "x"},
			givenError:  errors.New("existing"),
			expect:      Time{true, 12, 34, 56, 0},
			expectError: &ParseError{Errors: []error{errors.New("existing")}},
		},
		{
			name:        "nok, Time + already existing error",
			givenFields: []string{"1.2"},
			givenError:  errors.New("existing"),
			expect:      Time{false, 0, 0, 0, 0},
			expectError: &ParseError{Errors: []error{
				errors.New("existing"),
				errors.New("nmea: talkertype invalid context: index out of range"),
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewParser(BaseSentence{
				Talker: "talker",
				Type:   "type",
				Fields: tc.givenFields,
			})
			if tc.givenError != nil {
				p.setError(tc.givenError)
			}

			result := p.Time(1, "context")

			assert.Equal(t, tc.expect, result)
			assert.Equal(t, tc.expectError, p.Err())
		})
	}
}

func TestParser_Date(t *testing.T) {
	var testCases = []struct {
		name          string
		givenFields   []string
		givenError    error
		expect        Date
		expectError   error
		expectErrorIs error
	}{
		{
			name:        "ok, Date",
			givenFields: []string{"x", "010203", "x"},
			expect:      Date{true, 1, 2, 3},
		},
		{
			name:        "ok, field is empty",
			givenFields: []string{"1.2", "", "3.4"},
			expect:      Date{},
		},
		{
			name:        "nok, index out of range",
			givenFields: []string{"1.2"},
			expect:      Date{},
			expectError: &ParseError{Errors: []error{errors.New("nmea: talkertype invalid context: index out of range")}},
		},
		{
			name:        "nok, can not be parsed to Date",
			givenFields: []string{"1.2", "ABC"},
			expect:      Date{},
			expectError: &ParseError{
				Errors: []error{
					&FieldError{Prefix: "talkertype", Context: "context", Value: "ABC"},
				},
			},
			expectErrorIs: errors.New("nmea: talkertype invalid context: ABC"),
		},
		{
			name:        "ok, Date + already existing error",
			givenFields: []string{"x", "010203", "x"},
			givenError:  errors.New("existing"),
			expect:      Date{true, 1, 2, 3},
			expectError: &ParseError{Errors: []error{errors.New("existing")}},
		},
		{
			name:        "nok, Date + already existing error",
			givenFields: []string{"1.2"},
			givenError:  errors.New("existing"),
			expect:      Date{},
			expectError: &ParseError{Errors: []error{
				errors.New("existing"),
				errors.New("nmea: talkertype invalid context: index out of range"),
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewParser(BaseSentence{
				Talker: "talker",
				Type:   "type",
				Fields: tc.givenFields,
			})
			if tc.givenError != nil {
				p.setError(tc.givenError)
			}

			result := p.Date(1, "context")

			assert.Equal(t, tc.expect, result)
			assert.Equal(t, tc.expectError, p.Err())
		})
	}
}

func TestParser_LatLong(t *testing.T) {
	var testCases = []struct {
		name          string
		givenFields   []string
		givenError    error
		expect        float64
		expectError   error
		expectErrorIs error
	}{
		{
			name:        "ok, LatLong",
			givenFields: []string{"5000.0000", "N"},
			expect:      50.0,
		},
		{
			name:        "ok, field is empty",
			givenFields: []string{"", "", "3.4"},
			expect:      0,
		},
		{
			name:        "nok, index out of range",
			givenFields: []string{"5000.0000"},
			expect:      0,
			expectError: &ParseError{Errors: []error{
				errors.New("nmea: talkertype invalid context: index out of range"),
				&FieldError{Prefix: "talkertype", Context: "context", Value: "cannot parse [5000.0000 ], unknown format"},
			}},
		},
		{
			name:        "nok, can not be parsed to LatLong",
			givenFields: []string{"5000.0XX", "N"},
			expect:      0,
			expectError: &ParseError{Errors: []error{
				&FieldError{Prefix: "talkertype", Context: "context", Value: "cannot parse [5000.0XX N], unknown format"},
			}},
		},
		{
			name:        "nok, latitude out of range",
			givenFields: []string{"9100.0000", "N"},
			expect:      0,
			expectError: &ParseError{Errors: []error{
				&FieldError{Prefix: "talkertype", Context: "context", Value: "latitude is not in range (-90, 90)"},
			}},
		},
		{
			name:        "nok, longitude out of range",
			givenFields: []string{"18100.0000", "W"},
			expect:      0,
			expectError: &ParseError{Errors: []error{
				&FieldError{Prefix: "talkertype", Context: "context", Value: "longitude is not in range (-180, 180)"},
			}},
		},
		{
			name:        "ok, LatLong + already existing error",
			givenFields: []string{"5000.0000", "W"},
			givenError:  errors.New("existing"),
			expect:      -50.0,
			expectError: &ParseError{Errors: []error{errors.New("existing")}},
		},
		{
			name:        "nok, LatLong + already existing error",
			givenFields: []string{"5000.0000", "X"},
			givenError:  errors.New("existing"),
			expect:      0,
			expectError: &ParseError{Errors: []error{
				errors.New("existing"),
				&FieldError{Prefix: "talkertype", Context: "context", Value: "cannot parse [5000.0000 X], unknown format"},
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewParser(BaseSentence{
				Talker: "talker",
				Type:   "type",
				Fields: tc.givenFields,
			})
			if tc.givenError != nil {
				p.setError(tc.givenError)
			}

			result := p.LatLong(0, 1, "context")

			assert.Equal(t, tc.expect, result)
			assert.Equal(t, tc.expectError, p.Err())
		})
	}
}

func TestParser_SixBitASCIIArmour(t *testing.T) {
	var testCases = []struct {
		name         string
		givenFields  []string
		givenError   error
		whenFillBits int
		expect       []byte
		expectError  error
	}{
		{
			name:         "ok, SixBitASCIIArmour",
			givenFields:  []string{"13aGt0PP0jPN@9fMPKVDJgwfR>`<"},
			whenFillBits: 0,
			expect: []byte{
				0, 0, 0, 0, 0, 1, 0, 0, 0, 0,
				1, 1, 1, 0, 1, 0, 0, 1, 0, 1,
				0, 1, 1, 1, 1, 1, 1, 1, 0, 0,
				0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
				0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 1, 1, 0, 0, 1, 0,
				1, 0, 0, 0, 0, 0, 0, 1, 1, 1,
				1, 0, 0, 1, 0, 0, 0, 0, 0, 0,
				1, 0, 0, 1, 1, 0, 1, 1, 1, 0,
				0, 1, 1, 1, 0, 1, 1, 0, 0, 0,
				0, 0, 0, 1, 1, 0, 1, 1, 1, 0,
				0, 1, 1, 0, 0, 1, 0, 1, 0, 0,
				0, 1, 1, 0, 1, 0, 1, 0, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
				1, 1, 1, 0, 1, 0, 0, 0, 1, 0,
				0, 0, 1, 1, 1, 0, 1, 0, 1, 0,
				0, 0, 0, 0, 1, 1, 0, 0,
			},
		},
		{
			name:         "ok, SixBitASCIIArmour with fillbits = 2",
			givenFields:  []string{"H77nSfPh4U=<E`H4U8G;:222220"},
			whenFillBits: 2,
			expect: []byte{
				0, 1, 1, 0, 0, 0, 0, 0, 0, 1,
				1, 1, 0, 0, 0, 1, 1, 1, 1, 1,
				0, 1, 1, 0, 1, 0, 0, 0, 1, 1,
				1, 0, 1, 1, 1, 0, 1, 0, 0, 0,
				0, 0, 1, 1, 0, 0, 0, 0, 0, 0,
				0, 1, 0, 0, 1, 0, 0, 1, 0, 1,
				0, 0, 1, 1, 0, 1, 0, 0, 1, 1,
				0, 0, 0, 1, 0, 1, 0, 1, 1, 0,
				1, 0, 0, 0, 0, 1, 1, 0, 0, 0,
				0, 0, 0, 1, 0, 0, 1, 0, 0, 1,
				0, 1, 0, 0, 1, 0, 0, 0, 0, 1,
				0, 1, 1, 1, 0, 0, 1, 0, 1, 1,
				0, 0, 1, 0, 1, 0, 0, 0, 0, 0,
				1, 0, 0, 0, 0, 0, 1, 0, 0, 0,
				0, 0, 1, 0, 0, 0, 0, 0, 1, 0,
				0, 0, 0, 0, 1, 0, 0, 0, 0, 0,
			},
		},
		{
			name:         "ok, empty is ok",
			givenFields:  []string{""},
			whenFillBits: 0,
			expect:       []byte{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewParser(BaseSentence{
				Talker: "talker",
				Type:   "type",
				Fields: tc.givenFields,
			})
			if tc.givenError != nil {
				p.setError(tc.givenError)
			}

			result := p.SixBitASCIIArmour(0, tc.whenFillBits, "context")

			assert.Equal(t, tc.expect, result)
			assert.Equal(t, tc.expectError, p.Err())
		})
	}
}
