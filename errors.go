package nmea

import (
	"errors"
	"fmt"
)

// NotSupportedError is returned when parsed sentence is not supported
type NotSupportedError struct {
	Prefix string
}

// Error returns error message
func (p *NotSupportedError) Error() string {
	return fmt.Sprintf("nmea: sentence prefix '%s' not supported", p.Prefix)
}

// ParseError is error container to hold multiple error encountered during sentence parsing
type ParseError struct {
	Errors []error
}

func (e *ParseError) Error() string {
	return e.Errors[0].Error()
}

// As implements errors.As by attempting to map to the current value.
func (e *ParseError) As(target interface{}) bool {
	for _, err := range e.Errors {
		if ok := errors.As(err, target); ok {
			return true
		}
	}
	return false
}

// Is implements errors.Is by comparing the current value directly
func (e *ParseError) Is(target error) bool {
	for _, err := range e.Errors {
		if ok := errors.Is(err, target); ok {
			return true
		}
	}
	return false
}

// FieldError error structure related to parsing invalid field values
type FieldError struct {
	Prefix  string
	Context string
	Value   string
}

// Error returns error text
func (e *FieldError) Error() string {
	return fmt.Sprintf("nmea: %s invalid %s: %s", e.Prefix, e.Context, e.Value)
}

// Is implements errors.Is by comparing the current value directly
func (e *FieldError) Is(target error) bool {
	if t, ok := target.(*FieldError); ok {
		return t.Prefix == e.Prefix && t.Context == e.Context && t.Value == e.Value
	}
	// we want FieldError to be equal to old errors created as `errors.New("nmea: HEROT invalid rate of turn: nope")` was
	return e.Error() == target.Error()
}

// InvalidEnumCharsError is error returned when sentence has invalid character value for enum field
type InvalidEnumCharsError FieldError

// Error returns error text
func (e *InvalidEnumCharsError) Error() string {
	return fmt.Sprintf("nmea: %s invalid %s: %s", e.Prefix, e.Context, e.Value)
}

// InvalidEnumStringError is error returned when enum field in sentence has invalid value
type InvalidEnumStringError FieldError

// Error returns error text
func (e *InvalidEnumStringError) Error() string {
	return fmt.Sprintf("nmea: %s invalid %s: %s", e.Prefix, e.Context, e.Value)
}
