package zapevent

import (
	"fmt"
	"time"
	"strings"
)

const timeLayout = "2006/01/02, 15:04:05"

type ZapEvent struct {
	Dt time.Time
	Ip string
	FromCh string
	ToCh string
}

func NewZapEvent(data string) *ZapEvent {
	s := strings.Split(data, ", ")
	t, _ := time.Parse(timeLayout, data[0:20])
	if (len(s) < 5) {
		return &ZapEvent{t, s[2], s[3], "TASTATUR"}
	} else {
		return &ZapEvent{t, s[2], s[3], s[4]}
	}
}

func (ze *ZapEvent) String() string {
	return fmt.Sprintf("%s, %s, %s, %s", ze.Dt, ze.Ip, ze.FromCh, ze.ToCh)
}

/*func (ze *ZapEvent) Duration(provided ChZap) time.Duration {
	
	return time
}*/
