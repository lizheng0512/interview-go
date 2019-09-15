package sort

import (
	"fmt"
)

// 快速排序，是一种交换排序，平均时间复杂度为O(nlogn)，最坏的时间复杂度为O(n2)，最好的时间复杂度为O(nlogn)，
// 空间复杂度为O(nlogn)，是一种较为复杂的不稳定排序算法
// 时间复杂度：当数据有序时，以第一个元素为基准分为两个序列，前一个序列为空，此时执行效率最差，当数据随机分布时，
// 分成两个序列长度接近相等，此时执行效率最好。
// 空间复杂度：每次分割时需要一个空间存储基准值，而快速排序大概需要O(nlogn)次分割，所以空间复杂度为O(nlogn)
// 稳定性：在快速排序中，相等元素可能会因为分区而交换顺序，所以它是不稳定的算法。
func QuickSort(arr []int, start, end int) {
	if start < end {
		baseIndex := Partition(arr, start, end)
		QuickSort(arr, start, baseIndex-1)
		QuickSort(arr, baseIndex+1, end)
	}
}

func Partition(arr []int, start, end int) int {
	//baseIndex := rand.Intn(end-start) + start
	//base := arr[baseIndex]
	//Swap(arr, start, baseIndex)
	//取第一个元素作为基准
	baseIndex := start
	base := arr[baseIndex]
	fmt.Printf("start:%d, end:%d, baseIndex:%d, base:%d, arr:%v\n", start, end, baseIndex, base, arr)

	for start < end {
		for arr[end] > base && end > start {
			end--
		}
		if arr[end] <= base {
			arr[start] = arr[end]
		}
		for arr[start] <= base && start < end {
			start++
		}
		if arr[start] > base {
			arr[end] = arr[start]
		}
	}
	arr[start] = base
	return start
}

// 交换下表为a和b的元素
func Swap(arr []int, a, b int) {
	if !withinRange(arr, a) || !withinRange(arr, b) || a == b {
		return
	}
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

func withinRange(arr []int, index int) bool {
	return index >= 0 && index < len(arr)
}
