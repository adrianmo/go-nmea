package nmea

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestZZZ struct {
	BaseSentence
	NumberValue int
	StringValue string
}

var customparsetests = []struct {
	name string
	raw  string
	err  string
	msg  interface{}
}{
	{
		name: "yyy sentence",
		raw:  "$AAYYY,20,one,*13",
		msg: TestZZZ{
			BaseSentence: BaseSentence{
				Talker:   "AA",
				Type:     "YYY",
				Fields:   []string{"20", "one", ""},
				Checksum: "13",
				Raw:      "$AAYYY,20,one,*13",
			},
			NumberValue: 20,
			StringValue: "one",
		},
	},
	{
		name: "zzz sentence",
		raw:  "$AAZZZ,30,two,*19",
		msg: TestZZZ{
			BaseSentence: BaseSentence{
				Talker:   "AA",
				Type:     "ZZZ",
				Fields:   []string{"30", "two", ""},
				Checksum: "19",
				Raw:      "$AAZZZ,30,two,*19",
			},
			NumberValue: 30,
			StringValue: "two",
		},
	},
	{
		name: "zzz sentence type",
		raw:  "$INVALID,123,123,*7D",
		err:  "nmea: sentence prefix 'INVALID' not supported",
	},
	{
		name: "still works",
		raw:  "$GPZDA,172809.456,12,07,1996,00,00*57",
		msg: ZDA{
			BaseSentence: BaseSentence{
				Talker:   "GP",
				Type:     "ZDA",
				Fields:   []string{"172809.456", "12", "07", "1996", "00", "00"},
				Checksum: "57",
				Raw:      "$GPZDA,172809.456,12,07,1996,00,00*57",
			},
			Time:          Time{Valid: true, Hour: 17, Minute: 28, Second: 9, Millisecond: 456},
			Day:           12,
			Month:         7,
			Year:          1996,
			OffsetHours:   0,
			OffsetMinutes: 0,
		},
	},
}

func init() {

	// Register some custom parsers
	MustRegisterParser("YYY", func(s BaseSentence) (Sentence, error) {
		// Somewhat error prone parser without deps
		fields := strings.Split(s.Raw, ",")
		checksum := Checksum(s.Raw[1 : len(s.Raw)-3])
		checksumRaw := s.Raw[len(s.Raw)-2:]

		if checksum != checksumRaw {
			return nil, fmt.Errorf("nmea: sentence checksum mismatch [%s != %s]", checksum, checksumRaw)
		}

		nummericValue, _ := strconv.Atoi(fields[1])
		return TestZZZ{
			BaseSentence: s,
			NumberValue:  nummericValue,
			StringValue:  fields[2],
		}, nil
	})

	MustRegisterParser("ZZZ", func(s BaseSentence) (Sentence, error) {
		// Somewhat error prone parser

		p := NewParser(s)
		numberVal := int(p.Int64(0, "number"))
		stringVal := p.String(1, "str")
		return TestZZZ{
			BaseSentence: s,
			NumberValue:  numberVal,
			StringValue:  stringVal,
		}, p.Err()
	})
}

func TestCustomParser(t *testing.T) {
	for _, tt := range customparsetests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.msg, m)
			}
		})
	}
}

func TestWillReturnErrorOnDuplicateRegistration(t *testing.T) {
	err := RegisterParser("XXX", func(s BaseSentence) (Sentence, error) {
		return BaseSentence{}, nil
	})
	assert.NoError(t, err)

	err = RegisterParser("XXX", func(s BaseSentence) (Sentence, error) {
		return BaseSentence{}, nil
	})
	assert.Error(t, err)
}

func TestWillPanicOnDuplicateMustRegister(t *testing.T) {
	MustRegisterParser("AAA", func(s BaseSentence) (Sentence, error) {
		return BaseSentence{}, nil
	})

	assert.PanicsWithError(t, "nmea: parser for prefix 'AAA' already exists", func() {
		MustRegisterParser("AAA", func(s BaseSentence) (Sentence, error) {
			return BaseSentence{}, nil
		})
	})

}
