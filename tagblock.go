package nmea

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	// tagBlockRegexp matches nmea tag blocks
	tagBlockRegexp = regexp.MustCompile(`^(.*)\\(\S+)\\(.*)`)
)

// TagBlock struct
type TagBlock struct {
	Head         string // *
	Time         int64  // TypeUnixTime unix timestamp, parameter: -c
	RelativeTime int64  // TypeRelativeTime relative time time, parameter: -r
	Destination  string // TypeDestinationID destination identification 15 char max, parameter: -d
	Grouping     string // TypeGrouping sentence grouping, parameter: -g
	LineCount    int64  // TypeLineCount line count, parameter: -n
	Source       string // TypeSourceID source identification 15 char max, parameter: -s
	Text         string // TypeTextString valid character string, parameter -t
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
		checksum    = Checksum(fieldsRaw)
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
		case "c": // UNIX timestamp
			tagBlock.Time, err = parseInt64(item)
			if err != nil {
				return tagBlock, raw, err
			}
			tagBlock.Time, err = validUnixTimestamp(tagBlock.Time)
			if err != nil {
				return tagBlock, raw, err
			}
		case "d": // Destination ID
			tagBlock.Destination = item[2:]
		case "g": // Grouping
			tagBlock.Grouping = item[2:]
		case "n": // Line count
			tagBlock.LineCount, err = parseInt64(item)
			if err != nil {
				return tagBlock, raw, err
			}
		case "r": // Relative time
			tagBlock.RelativeTime, err = parseInt64(item)
			if err != nil {
				return tagBlock, raw, err
			}
		case "s": // Source ID
			tagBlock.Source = item[2:]
		case "t": // Text string
			tagBlock.Text = item[2:]
		}
	}
	return tagBlock, raw, nil
}
