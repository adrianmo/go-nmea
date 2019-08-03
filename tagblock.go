package nmea

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	// TypeUnixTime unix timestamp, parameter: -c
	TypeUnixTime = "c"
	// TypeDestinationID destination identification 15char max, parameter: -d
	TypeDestinationID = "d"
	// TypeGrouping sentence grouping, parameter: -g
	TypeGrouping = "g"
	// TypeLineCount linecount, parameter: -n
	TypeLineCount = "n"
	// TypeRelativeTime relative time time, paremeter: -r
	TypeRelativeTime = "r"
	// TypeSourceID source identification 15char max, paremter: -s
	TypeSourceID = "s"
	// TypeTextString valid character string, parameter -t
	TypeTextString = "t"
)

var (
	// TagBlockRegexp matches nmea tag blocks
	TagBlockRegexp = regexp.MustCompile(`^(.*)\\(\S+)\\(.*)`)
)

// TagBlock struct
type TagBlock struct {
	Head         string // *
	Time         uint32 // -c
	RelativeTime uint32 // -r
	Destination  string // -d 15 char max
	Grouping     string // -g nummeric string
	LineCount    uint32 // -n int
	Source       string // -s 15 char max
	Text         string // -t Variable length text
}

func parseUint(raw string) (uint32, error) {
	i, err := strconv.ParseUint(raw[2:], 10, 32)
	if err != nil {
		return 0, fmt.Errorf("nmea: tagblock unable to parse uint32 [%s]", raw)
	}
	return uint32(i), nil
}

// parseTagBlock adds support for tagblocks
// https://rietman.wordpress.com/2016/09/17/nemastudio-now-supports-the-nmea-0183-tag-block/
func parseTagBlock(raw string) (TagBlock, string, error) {
	matches := TagBlockRegexp.FindStringSubmatch(raw)
	var tagBlock TagBlock
	if matches != nil {
		raw = matches[3]
		tags := matches[2]
		tagBlock.Head = matches[1]

		sumSepIndex := strings.Index(tags, ChecksumSep)
		if sumSepIndex == -1 {
			return tagBlock, raw, fmt.Errorf("nmea: tagblock does not contain checksum separator")
		}

		var (
			fieldsRaw   = tags[0:sumSepIndex]
			checksumRaw = strings.ToUpper(tags[sumSepIndex+1:])
			checksum    = xorChecksum(fieldsRaw)
			err         error
		)

		// Validate the checksum
		if checksum != checksumRaw {
			return tagBlock, raw, fmt.Errorf("nmea: tagblock checksum mismatch [%s != %s]", checksum, checksumRaw)
		}

		data := strings.Split(tags[:sumSepIndex], ",")
		for _, item := range data {
			switch item[0:1] {
			case TypeUnixTime:
				tagBlock.Time, err = parseUint(item)
				if err != nil {
					return tagBlock, raw, err
				}
			case TypeDestinationID:
				tagBlock.Destination = item[2:]
			case TypeGrouping:
				tagBlock.Grouping = item[2:]
			case TypeLineCount:
				tagBlock.LineCount, err = parseUint(item)
				if err != nil {
					return tagBlock, raw, err
				}
			case TypeRelativeTime:
				tagBlock.RelativeTime, err = parseUint(item)
				if err != nil {
					return tagBlock, raw, err
				}
			case TypeSourceID:
				tagBlock.Source = item[2:]
			case TypeTextString:
				tagBlock.Text = item[2:]
			}
		}

	}

	return tagBlock, raw, nil
}
