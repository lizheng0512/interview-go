package sdk_sort

import (
	"sort"
	"testing"
)

type myInt []int

func (mi myInt) Len() int {
	return len(mi)
}

func (mi myInt) Less(i, j int) (result bool) {
	if mi[i] < mi[j] {
		result = true
	}
	return
}

func (mi myInt) Swap(i, j int) {
	mi[i], mi[j] = mi[j], mi[i]
}

func TestSort(t *testing.T) {
	mi := make(myInt, 0, 1<<32-1)
	for i := 0; i < 1<<63-1; i++ {
		mi[i] = i
	}
	sort.Sort(mi)
}
