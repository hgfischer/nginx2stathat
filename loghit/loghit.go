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
	remoteAddress string
	remoteUser    string
	localTime     string
	request       string
	status        int
	bodyBytesSent int
	httpReferer   string
	httpUserAgent string
}

func (lh *LogHit) String() string {
	return fmt.Sprintf(
		"%s\t%s\t%s\t%s\t%d\t%d\t%s\t%s",
		lh.remoteAddress,
		lh.remoteUser,
		lh.localTime,
		lh.request,
		lh.status,
		lh.bodyBytesSent,
		lh.httpReferer,
		lh.httpUserAgent,
	)
}

func New(line string) (*LogHit, error) {
	var err error
	matches := CombinedRegex.FindStringSubmatch(line)
	if len(matches) != 9 {
		return nil, errors.New(fmt.Sprintf("Cannot parse: \"%s\"", line))
	}

	logHit := LogHit{}
	logHit.remoteAddress = matches[1]
	logHit.remoteUser = matches[2]
	logHit.localTime = matches[3]
	logHit.request = matches[4]
	logHit.status, err = strconv.Atoi(matches[5])
	if err != nil {
		return nil, err
	}

	logHit.bodyBytesSent, err = strconv.Atoi(matches[6])
	if err != nil {
		return nil, err
	}

	logHit.httpReferer = matches[7]
	logHit.httpUserAgent = matches[8]

	return &logHit, nil
}
