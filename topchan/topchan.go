package topchan

import (
	"zaplab/mapsort"
	"zaplab/ztorage"
	"fmt"
	"time"
)

func ChCount(zaps *ztorage.Zaps, m map[string]int) {
	//m = make(map[string]int)
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

func TopTen(zapstore *ztorage.Zaps, m map[string]int){
	        i := 0
		for {
			ChCount(zapstore, m)

			time.Sleep(1*time.Second)
			sm := mapsort.SortedKeys(m)

			if i > 0 {
				 fmt.Printf("Top 10 channels: %v\n", sm[0:10])
			}

			i++
		}
}


