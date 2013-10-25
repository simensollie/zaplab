package topchan

import "zaplab/mapsort"

func ChCount(zaps string) {
	m = make(map[*zaps]int)
	for _, v := range *zaps {
		if m[zaps.FromCh]{
			if m[zaps.FromCh] > 0 {
				m[zaps.FromCh] =- 1
			}
		} else {
			m[zaps.FromCh] = 0
		}

		if m[zaps.ToCh] {
			m[zaps.ToCh] += 1
		} else  {
			m[zaps.ToCh] = 1
		}
	}
	return m
}

func TopCh(){
	for{
		time.Sleep(1 * time.Second)
		ms := mapsort.NewMapSorter(m)
		ms.Sort()
		//print top10
	}
}

