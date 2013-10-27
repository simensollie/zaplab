package topchan

import (
	"zaplab/mapsort"
	"zaplab/ztorage"
	"time"
)

func ChCount(zaps *ztorage.Zaps, m map[string]int) {
	for _, v := range *zaps {
		if m[v.FromCh] > 0 {
			m[v.FromCh] -= 1
		} else {
			m[v.FromCh] = 0
		}

		if m[v.ToCh] >= 0 {
			m[v.ToCh] += 1
		} else  {
			m[v.ToCh] = 1
		}
	}
}

func TopTen(zapstore *ztorage.Zaps, m map[string]int) string {
	        i := 0
		s := "Top 10 channels: "
		for {
			ChCount(zapstore, m)

			//time.Sleep(1*time.Second)
			sm := mapsort.SortedKeys(m)


			if i > 0 {
				for i := 1; i < 11; i++ {
					s += "\n" + sm[i]
				}
				return s
			}
			i++
		}
		return ""
}


