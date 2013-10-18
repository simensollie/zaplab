package ztorage

type Zaps []zapEvent

func NewZapStore() *Zaps {
	zs := make(Zaps, 0)
	return &zs
}

func (zs *Zaps) StoreZap(z zapEvent) {
	*zs = append(*zs, z)
}

func (zs *Zaps) ComputeViewers(chName string) int {
	viewers := 0
	for _, v := range *zs {
		if v.ToChan == chName {
		viewers++
		}
		if v.FromChan == chName {
		viewers--
		}
	}
	return viewers
}

