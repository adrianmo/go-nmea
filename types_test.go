package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var nearDistance = 0.001

func TestParseGPS(t *testing.T) {
	var tests = []struct {
		value    string
		expected LatLong
		err      bool
	}{
		{"3345.1232 N", 33.752054, false},
		{"15145.9877 S", -151.76646, false},
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			l, err := ParseGPS(tt.value)
			if tt.err {
				assert.Error(t, err)
			} else {
				if !l.IsNear(tt.expected, nearDistance) {
					t.Errorf("ParseGPS got %f, expected %f", l, tt.expected)
				}
			}
		})
	}
}

func TestParseDMS(t *testing.T) {
	var tests = []struct {
		value    string
		expected LatLong
		err      bool
	}{
		{"33\u00B0 12' 34.3423\"", 33.209540, false},
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			l, err := ParseDMS(tt.value)
			if tt.err {
				assert.Error(t, err)
			} else {
				if !l.IsNear(tt.expected, nearDistance) {
					t.Errorf("ParseDMS got %f, expected %f", l, tt.expected)
				}
			}
		})
	}
}

func TestParseDecimal(t *testing.T) {
	var tests = []struct {
		value    string
		expected LatLong
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
				if !l.IsNear(tt.expected, nearDistance) {
					t.Errorf("ParseDecimal got %f, expected %f", l, tt.expected)
				}
			}
		})
	}
}

func TestLatLongPrint(t *testing.T) {
	l, _ := ParseDecimal("151.434367")
	exp := "15126.0620"
	if s := l.PrintGPS(); s != exp {
		t.Errorf("PrintGPS() got %s expected %s", s, exp)
	}

	l, _ = ParseGPS("3356.4343 N")
	exp = "3356.4343"
	if s := l.PrintGPS(); s != exp {
		t.Errorf("PrintGPS() got %s expected %s", s, exp)
	}

	exp = "33° 56' 26.058000\""
	if s := l.PrintDMS(); s != exp {
		t.Errorf("PrintDMS() got %s expected %s", s, exp)
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
		{"010203.04", Time{true, 1, 2, 3, 4}, true},
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
	expected := "01:02:03.0004"
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
