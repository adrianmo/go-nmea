package nmea

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	// tagBlockRegexp matches nmea tag blocks
	tagBlockRegexp = regexp.MustCompile(`^(.*)\\(\S+)\\(.*)`)
)

// TagBlock struct
type TagBlock struct {
	Head         string // *
	Time         int64  // TypeUnixTime unix timestamp (unit is likely to be s, but might be ms, YMMV), parameter: -c
	RelativeTime int64  // TypeRelativeTime relative time, parameter: -r
	Destination  string // TypeDestinationID destination identification 15 char max, parameter: -d
	Grouping     string // TypeGrouping sentence grouping, parameter: -g
	LineCount    int64  // TypeLineCount line count, parameter: -n
	Source       string // TypeSourceID source identification 15 char max, parameter: -s
	Text         string // TypeTextString valid character string, parameter -t
}

func parseInt64(raw string) (int64, error) {
	i, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("nmea: tagblock unable to parse uint64 [%s]", raw)
	}
	return i, nil
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
		return TagBlock{}, "", fmt.Errorf("nmea: tagblock does not contain checksum separator")
	}

	var (
		fieldsRaw   = tags[0:sumSepIndex]
		checksumRaw = strings.ToUpper(tags[sumSepIndex+1:])
		checksum    = Checksum(fieldsRaw)
		err         error
	)

	// Validate the checksum
	if checksum != checksumRaw {
		return TagBlock{}, "", fmt.Errorf("nmea: tagblock checksum mismatch [%s != %s]", checksum, checksumRaw)
	}

	items := strings.Split(tags[:sumSepIndex], ",")
	for _, item := range items {
		parts := strings.SplitN(item, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key, value := parts[0], parts[1]
		switch key {
		case "c": // UNIX timestamp
			tagBlock.Time, err = parseInt64(value)
			if err != nil {
				return TagBlock{}, "", err
			}
		case "d": // Destination ID
			tagBlock.Destination = value
		case "g": // Grouping
			tagBlock.Grouping = value
		case "n": // Line count
			tagBlock.LineCount, err = parseInt64(value)
			if err != nil {
				return TagBlock{}, "", err
			}
		case "r": // Relative time
			tagBlock.RelativeTime, err = parseInt64(value)
			if err != nil {
				return TagBlock{}, "", err
			}
		case "s": // Source ID
			tagBlock.Source = value
		case "t": // Text string
			tagBlock.Text = value
		}
	}
	return tagBlock, raw, nil
}
