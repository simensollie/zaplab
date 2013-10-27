package ztorage

import (
	"zaplab/zapevent"
	"fmt"
	"time"
)

type Zaps []zapevent.ZapEvent
var m = make(map[string]zapevent.ZapEvent)

func NewZapStore() *Zaps {
	zs := make(Zaps, 0)
	return &zs
}

func (zs *Zaps) StoreZap(z zapevent.ZapEvent) {
	zs.PrevZap(z)
	*zs = append(*zs, z)
}

func (zs *Zaps) ComputeViewers(chName string) int {
	viewers := 0
	for _, v := range *zs {
		if v.ToCh == chName {
		viewers++
		}
		if v.FromCh == chName {
		viewers--
		}
	}
	return viewers
}

func (zs *Zaps) PrevZap(zap zapevent.ZapEvent) {
	m[zap.Ip] = zap
}

func ComputeDuration(zap zapevent.ZapEvent) {
	value, exists := m[zap.Ip]
	var dur time.Duration
	if exists {
		dur = zap.Duration(&value)
		if dur != 0 {
			fmt.Println(dur)

		}
	}
}

