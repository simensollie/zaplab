package zapevent

import (
	"fmt"
	"time"
	"strings"
)

const timeLayout = "2006/01/02, 15:04:05"

type zapEvent struct {
	datetime time.Time
	ip string
	fromCh string
	toCh string
}

func newZapEvent(data string) *zapEvent {
	s := strings.Split(data, ", ")
	t, _ := time.Parse(timeLayout, data[0:20])
	if (len(s) < 5) {
		return &zapEvent{t, s[2], s[3], "TASTATUR"}
	} else {
		return &zapEvent{t, s[2], s[3], s[4]}
	}
}

func (ze *zapEvent) String() string {
	return fmt.Sprintf("%s, %s, %s, %s", ze.datetime, ze.ip, ze.fromCh, ze.toCh)
}

/*func (ze *zapEvent) Duration(provided ChZap) time.Duration {
	
	return time
}*/
