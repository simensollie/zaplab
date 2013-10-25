package topchan

import (
//	"zaplab/mapsort"
	"zaplab/ztorage"
//	"fmt"
//	"time"
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

/*func TopCh(m map[string]int){
	for{
		time.Sleep(1 * time.Second)
		fmt.Printf("%v\n", len(m))
		ms := mapsort.NewMapSort(m)
		fmt.Printf("%v\n", *ms)
		ms.Sort()
		fmt.Printf("%v\n", *ms)
		//print top10
		fmt.Println("Top 10 channels:")
		for i := 0; i < 10; i++ {
			fmt.Println(ms[i])
		}
	}
}*/

