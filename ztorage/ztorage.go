package ztorage

import (
	"zaplab/zapevent"
)

type Zaps []zapevent.ZapEvent
var m = make(map[string]ZapEvent)

func NewZapStore() *Zaps {
	zs := make(Zaps, 0)
	return &zs
}

func (zs *Zaps) StoreZap(z zapevent.ZapEvent) {
	prevZap(z)
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

func (zs *Zaps) prevZap(zap zapevent.ZapEvent) {
	m[zap.Ip] = zap.Dt
}

func (zs *Zaps) Duration(zap *zapevent.ZapEvent) string {
	v, exists := zs.prevZap[zap.Ip]
	dur := ""
	if exists {
		dur += zapevent.Duration(&v).String()
	}
	return dur
}
