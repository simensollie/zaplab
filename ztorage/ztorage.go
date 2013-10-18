package ztorage

import (
	"zaplab/zapevent"
)

type Zaps []zapevent.ZapEvent

func NewZapStore() *Zaps {
	zs := make(Zaps, 0)
	return &zs
}

func (zs *Zaps) StoreZap(z zapevent.ZapEvent) {
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

