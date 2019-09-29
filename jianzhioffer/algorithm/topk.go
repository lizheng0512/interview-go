package algorithm

// 求海量数据中的最大的topK
// 思路：找最大的topK用小根堆实现，先把海量数据中的前K个节点构建为小根堆，再把剩余的数据挨个与根节点比较，
// 因为小根堆的根节点是最小的，如果比根节点小，则直接舍弃，如果比根节点大，则替换根节点，并调整堆。
func TopK(arr []int, k int) []int {
	if arr == nil || len(arr) == 0 || k <= 0 {
		return []int{}
	}
	if len(arr) < k {
		k = len(arr)
	}
	top := make([]int, 0, k)
	for i := 0; i < k; i++ {
		top = append(top, arr[i])
	}
	// 构建初始堆
	for i := k / 2; i >= 0; i-- {
		adjustMinHeap(top, i)
	}
	for i := k; i < len(arr); i++ {
		if arr[i] > top[0] {
			top[0] = arr[i]
			adjustMinHeap(top, 0)
		}
	}
	// 给top排一下序
	sortHeap(top)
	return top
}

// 调整小根堆
func adjustMinHeap(arr []int, pIndex int) {
	parent := arr[pIndex]
	cIndex := 2*pIndex + 1
	for cIndex < len(arr) {
		if cIndex+1 < len(arr) && arr[cIndex+1] < arr[cIndex] {
			cIndex++
		}
		if parent < arr[cIndex] {
			break
		}
		arr[pIndex] = arr[cIndex]

		pIndex = cIndex
		cIndex = 2*pIndex + 1
	}
	arr[pIndex] = parent
}

// 给topK排序，采用快排
func sortHeap(arr []int) {
	if len(arr) <= 1 {
		return
	}
	base := arr[0]
	start, end := 0, len(arr)-1
	for start < end {
		for arr[end] <= base && end > start {
			end--
		}
		arr[start] = arr[end]
		for arr[start] >= base && start < end {
			start++
		}
		arr[end] = arr[start]
	}
	arr[start] = base
	sortHeap(arr[:start])
	sortHeap(arr[start+1:])
}
