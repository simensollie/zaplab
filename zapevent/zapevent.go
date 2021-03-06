package zapevent

import (
	"fmt"
	"strings"
	"time"
)

const timeLayout = "2006/01/02, 15:04:05"

var m = make(map[string]time.Time)

type ZapEvent struct {
	Dt     time.Time
	Ip     string
	FromCh string
	ToCh   string
}

func NewZapEvent(data string) *ZapEvent {
	s := strings.Split(data, ", ")
	t, _ := time.Parse(timeLayout, data[0:20])
	if len(s) < 5 {
		return &ZapEvent{Dt: t, Ip: s[2], FromCh: s[3]}
	} else {
		return &ZapEvent{Dt: t, Ip: s[2], FromCh: s[3], ToCh: s[4]}
	}

	return nil
}

func (ze *ZapEvent) String() string {
	return fmt.Sprintf("%s, %s, %s, %s", ze.Dt, ze.Ip, ze.FromCh, ze.ToCh)
}

func (ze *ZapEvent) Duration(newZap *ZapEvent) time.Duration {
	return ze.Dt.Sub(newZap.Dt)
}
