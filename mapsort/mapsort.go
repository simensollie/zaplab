package mapsort

import "sort"

type MapSorter struct {
	Keys []string
	Values []int
}

func NewMapSorter(m map[string]int) *MapSorter {
	ms := &MapSorter{
		Keys: make([]string, 0, len(m)),
		Values: make([]int, 0, len(m)),
	}
	for k, v := range m {
		ms.Keys = append(ms.Keys, k)
		ms.Values = append(ms.Values, v)
	}
	return ms
}

func (ms *MapSorter) Sort() {
	sort.Sort(ms)
}

func (ms *MapSorter) Len() int {
	return len(ms.Values)
}

func (ms *MapSorter) Less(i, j int) bool {
	return ms.Values[i] > ms.Values[j]
}

func (ms *MapSorter) swap(i, j int) {
	ms.Values[i], ms.Values[j] = ms.Values[j], ms.Values[i]
	ms.Keys[i], ms.Keys[j] = ms.Keys[j], ms.Keys[i]
}
