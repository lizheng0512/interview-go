package array

import "container/list"

func LexicalOrder(n int) []int {
	l := list.New()
	lexicalOrderCore(l, 0, n)
	arr := make([]int, 0, l.Len())
	for front := l.Front(); front != nil; front = front.Next() {
		arr = append(arr, front.Value.(int))
	}
	return arr
}

func lexicalOrderCore(result *list.List, current, max int) {
	if current > max {
		return
	}
	if current != 0 {
		result.PushBack(current)
	}
	result.Front()
	for nextBit := 0; nextBit <= 9; nextBit++ {
		if current == 0 && nextBit == 0 {
			continue
		}
		lexicalOrderCore(result, current*10+nextBit, max)
	}
}
