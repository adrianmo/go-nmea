package nmea

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
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
	// tagBlockRegexp matches nmea tag blocks
	tagBlockRegexp = regexp.MustCompile(`^(.*)\\(\S+)\\(.*)`)
)

// TagBlock struct
type TagBlock struct {
	Head         string // *
	Time         int64  // -c
	RelativeTime int64  // -r
	Destination  string // -d 15 char max
	Grouping     string // -g nummeric string
	LineCount    int64  // -n int
	Source       string // -s 15 char max
	Text         string // -t Variable length text
}

func parseInt64(raw string) (int64, error) {
	i, err := strconv.ParseInt(raw[2:], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("nmea: tagblock unable to parse uint32 [%s]", raw)
	}
	return i, nil
}

// Timestamp can come as milliseconds or seconds
func validUnixTimestamp(timestamp int64) (int64, error) {
	if timestamp < 0 {
		return 0, errors.New("nmea: Tagblock timestamp is not valid must be between 0 and now + 24h")
	}
	now := time.Now()
	unix := now.Unix() + 24*3600
	if timestamp > unix {
		if timestamp > unix*1000 {
			return 0, errors.New("nmea: Tagblock timestamp is not valid")
		}
		return timestamp / 1000, nil
	}

	return timestamp, nil
}

// parseTagBlock adds support for tagblocks
// https://rietman.wordpress.com/2016/09/17/nemastudio-now-supports-the-nmea-0183-tag-block/
func parseTagBlock(raw string) (TagBlock, string, error) {
	matches := tagBlockRegexp.FindStringSubmatch(raw)
	if matches == nil {
		return TagBlock{}, raw, nil
	}

	tagBlock := TagBlock{}
	raw = matches[3]
	tags := matches[2]
	tagBlock.Head = matches[1]

	sumSepIndex := strings.Index(tags, ChecksumSep)
	if sumSepIndex == -1 {
		return tagBlock, "", fmt.Errorf("nmea: tagblock does not contain checksum separator")
	}

	var (
		fieldsRaw   = tags[0:sumSepIndex]
		checksumRaw = strings.ToUpper(tags[sumSepIndex+1:])
		checksum    = xorChecksum(fieldsRaw)
		err         error
	)

	// Validate the checksum
	if checksum != checksumRaw {
		return tagBlock, "", fmt.Errorf("nmea: tagblock checksum mismatch [%s != %s]", checksum, checksumRaw)
	}

	items := strings.Split(tags[:sumSepIndex], ",")
	for _, item := range items {
		if len(item) == 0 {
			continue
		}
		switch item[:1] {
		case TypeUnixTime:
			tagBlock.Time, err = parseInt64(item)
			if err != nil {
				return tagBlock, raw, err
			}
			tagBlock.Time, err = validUnixTimestamp(tagBlock.Time)
			if err != nil {
				return tagBlock, raw, err
			}
		case TypeDestinationID:
			tagBlock.Destination = item[2:]
		case TypeGrouping:
			tagBlock.Grouping = item[2:]
		case TypeLineCount:
			tagBlock.LineCount, err = parseInt64(item)
			if err != nil {
				return tagBlock, raw, err
			}
		case TypeRelativeTime:
			tagBlock.RelativeTime, err = parseInt64(item)
			if err != nil {
				return tagBlock, raw, err
			}
		case TypeSourceID:
			tagBlock.Source = item[2:]
		case TypeTextString:
			tagBlock.Text = item[2:]
		}
	}
	return tagBlock, raw, nil
}
