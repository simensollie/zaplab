package mapsort
 
import "sort"
 
/*func main() {
	m := map[string]int{
		"One": 1,
		"Two": 2,
		"Three": 3,
		"Ten": 10,
		"Fifty": 50,
	}
	vs := NewValSorter(m)
	fmt.Printf("%v\n", *vs)
	vs.Sort()
	fmt.Printf("%v\n", *vs)
}*/
 
type MapSorter struct {
	Keys []string
	Vals []int
}
 
func NewMapSorter(m map[string]int) *MapSorter {
	ms := &MapSorter{
		Keys: make([]string, 0, len(m)),
		Vals: make([]int, 0, len(m)),
	}
	for k, v := range m {
		ms.Keys = append(ms.Keys, k)
		ms.Vals = append(ms.Vals, v)
	}
	return ms
}
 
func (ms *MapSorter) Sort() {
	sort.Sort(ms)
}
 
func (ms *MapSorter) Len() int {
	return len(ms.Vals)
}

func (ms *MapSorter) Less(i, j int) bool {
	return ms.Vals[i] < ms.Vals[j]
}

func (ms *MapSorter) Swap(i, j int) {
	ms.Vals[i], ms.Vals[j] = ms.Vals[j], ms.Vals[i]
	ms.Keys[i], ms.Keys[j] = ms.Keys[j], ms.Keys[i]
}
