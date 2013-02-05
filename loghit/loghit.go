package loghit

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const (
	COMBINED_REGEX = "^(?P<RemoteAddress>(?:\\d{1,3}\\.){3}\\d{1,3}) \\- (?P<RemoteUser>\\S+) \\[(?P<LocalTime>.+?)\\] \"(?P<Request>.+?)\" (?P<Status>\\d{3}) (?P<BodyBytesSent>\\d+) \"(?P<HttpReferer>.+?)\" \"(?P<HttpUserAgent>.+?)\"$"
)

var CombinedRegex = regexp.MustCompile(COMBINED_REGEX)

type LogHit struct {
	RemoteAddress string
	RemoteUser    string
	LocalTime     string
	Request       string
	Status        int
	BodyBytesSent int
	HttpReferer   string
	HttpUserAgent string
}

func (lh *LogHit) String() string {
	return fmt.Sprintf(
		"%s\t%s\t%s\t%s\t%d\t%d\t%s\t%s",
		lh.RemoteAddress,
		lh.RemoteUser,
		lh.LocalTime,
		lh.Request,
		lh.Status,
		lh.BodyBytesSent,
		lh.HttpReferer,
		lh.HttpUserAgent,
	)
}

func New(line string) (*LogHit, error) {
	var err error
	matches := CombinedRegex.FindStringSubmatch(line)
	if len(matches) != 9 {
		return nil, errors.New(fmt.Sprintf("Cannot parse: \"%s\"", line))
	}

	logHit := LogHit{}
	logHit.RemoteAddress = matches[1]
	logHit.RemoteUser = matches[2]
	logHit.LocalTime = matches[3]
	logHit.Request = matches[4]
	logHit.Status, err = strconv.Atoi(matches[5])
	if err != nil {
		return nil, err
	}

	logHit.BodyBytesSent, err = strconv.Atoi(matches[6])
	if err != nil {
		return nil, err
	}

	logHit.HttpReferer = matches[7]
	logHit.HttpUserAgent = matches[8]

	return &logHit, nil
}
