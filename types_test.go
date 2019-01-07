package nmea

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var nearDistance = 0.001

func TestParseLatLong(t *testing.T) {
	var tests = []struct {
		value    string
		expected float64
		err      bool
	}{
		{"33\u00B0 12' 34.3423\"", 33.209540, false}, // dms
		{"3345.1232 N", 33.752054, false},            // gps
		{"151.234532", 151.234532, false},            // decimal
		{"200.000", 0, true},                         // out of range
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			l, err := ParseLatLong(tt.value)
			if tt.err {
				assert.Error(t, err)
			} else {
				assert.InDelta(t, tt.expected, l, nearDistance)
			}
		})
	}
}

func TestParseGPS(t *testing.T) {
	var tests = []struct {
		value    string
		expected float64
		err      bool
	}{
		{"3345.1232 N", 33.752054, false},
		{"15145.9877 S", -151.76646, false},
		{"12345.1234 X", 0, true},
		{"1234.1234", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			l, err := ParseGPS(tt.value)
			if tt.err {
				assert.Error(t, err)
			} else {
				assert.InDelta(t, tt.expected, l, nearDistance)
			}
		})
	}
}

func TestParseDMS(t *testing.T) {
	var tests = []struct {
		value    string
		expected float64
		err      bool
	}{
		{"33\u00B0 12' 34.3423\"", 33.209540, false},
		{"33\u00B0 1.1' 34.3423\"", 0, true},
		{"3.3\u00B0 1' 34.3423\"", 0, true},
		{"33\u00B0 1' 34.34.23\"", 0, true},
		{"33 1 3434.23", 0, true},
		{"123", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			l, err := ParseDMS(tt.value)
			if tt.err {
				assert.Error(t, err)
			} else {
				assert.InDelta(t, tt.expected, l, nearDistance)
			}
		})
	}
}

func TestParseDecimal(t *testing.T) {
	var tests = []struct {
		value    string
		expected float64
		err      bool
	}{
		{"151.234532", 151.234532, false},
		{"-151.234532", -151.234532, false},
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			l, err := ParseDecimal(tt.value)
			if tt.err {
				assert.Error(t, err)
			} else {
				assert.InDelta(t, tt.expected, l, nearDistance)
			}
		})
	}
}

func TestLatLongPrint(t *testing.T) {
	var tests = []struct {
		value float64
		dms   string
		gps   string
	}{
		{
			value: 151.434367,
			gps:   "15126.0620",
			dms:   "151° 26' 3.721200\"",
		},
		{
			value: 33.94057166666666,
			gps:   "3356.4343",
			dms:   "33° 56' 26.058000\"",
		},
		{
			value: 45.0,
			dms:   "45° 0' 0.000000\"",
			gps:   "4500.0000",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%f", tt.value), func(t *testing.T) {
			assert.Equal(t, tt.dms, FormatDMS(tt.value))
			assert.Equal(t, tt.gps, FormatGPS(tt.value))
		})
	}
}

func TestTimeParse(t *testing.T) {
	timetests := []struct {
		value    string
		expected Time
		ok       bool
	}{
		{"123456", Time{true, 12, 34, 56, 0}, true},
		{"", Time{}, true},
		{"112233.123", Time{true, 11, 22, 33, 123}, true},
		{"010203.04", Time{true, 1, 2, 3, 40}, true},
		{"10203.04", Time{}, false},
		{"x0u2xd", Time{}, false},
		{"xx2233.123", Time{}, false},
		{"11xx33.123", Time{}, false},
		{"1122xx.123", Time{}, false},
		{"112233.xxx", Time{}, false},
	}
	for _, tt := range timetests {
		actual, err := ParseTime(tt.value)
		if !tt.ok {
			if err == nil {
				t.Errorf("ParseTime(%s) expected error", tt.value)
			}
		} else {
			if err != nil {
				t.Errorf("ParseTime(%s) %s", tt.value, err)
			}
			if actual != tt.expected {
				t.Errorf("ParseTime(%s) got %s expected %s", tt.value, actual, tt.expected)
			}
		}
	}
}

func TestTimeString(t *testing.T) {
	d := Time{
		Hour:        1,
		Minute:      2,
		Second:      3,
		Millisecond: 4,
	}
	expected := "01:02:03.0040"
	if s := d.String(); s != expected {
		t.Fatalf("got %s, expected %s", s, expected)
	}
}

func TestDateParse(t *testing.T) {
	datetests := []struct {
		value    string
		expected Date
		ok       bool
	}{
		{"010203", Date{true, 1, 2, 3}, true},
		{"01003", Date{}, false},
		{"", Date{}, true},
		{"xx0203", Date{}, false},
		{"01xx03", Date{}, false},
		{"0102xx", Date{}, false},
	}
	for _, tt := range datetests {
		actual, err := ParseDate(tt.value)
		if !tt.ok {
			if err == nil {
				t.Errorf("ParseDate(%s) expected error", tt.value)
			}
		} else {
			if err != nil {
				t.Errorf("ParseDate(%s) %s", tt.value, err)
			}
			if actual != tt.expected {
				t.Errorf("ParseDate(%s) got %s expected %s", tt.value, actual, tt.expected)
			}
		}
	}
}

func TestDateString(t *testing.T) {
	d := Date{
		DD: 1,
		MM: 2,
		YY: 3,
	}
	expected := "01/02/03"
	if s := d.String(); s != expected {
		t.Fatalf("got %s expected %s", s, expected)
	}
}
