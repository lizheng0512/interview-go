package sort

import "fmt"

// 归并排序，归并排序是将n个长度为1的有序序列，合并为n/2个长度为2的有序序列，
// 再合并为n/4个长度为4的有序序列……直到合并为一个长度为n的有序序列，
// 所以归并排序需要先拆分再合并。
// 时间复杂度为O(nlogn)，空间复杂度O(n)，是一种稳定排序算法。
func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	// gap表示子序列长度，从1开始，每次×2
	for gap := 1; gap < len(arr); gap *= 2 {
		// 每两个相邻子序列进行合并
		for i := 0; i < len(arr); i += 2 * gap {
			merge(arr, i, i+gap, i+2*gap)
		}
	}
}

// 合并从arr[start:middle]和arr[middle:end]，合并后为递增序列
func merge(arr []int, start, middle, end int) {
	if middle > len(arr) {
		middle = len(arr)
	}
	if end > len(arr) {
		end = len(arr)
	}
	if middle == end {
		return
	}
	fmt.Printf("start: %d, middle: %d, end: %d\n", start, middle, end)
	mergedArr := make([]int, end-start)
	mergedIndex := 0
	originStart, originEnd, start2 := start, end, middle
	for start < middle && start2 < end {
		if arr[start] < arr[start2] {
			mergedArr[mergedIndex] = arr[start]
			start++
		} else {
			mergedArr[mergedIndex] = arr[start2]
			start2++
		}
		mergedIndex++
	}
	for start < middle {
		mergedArr[mergedIndex] = arr[start]
		start++
		mergedIndex++
	}
	for start2 < end {
		mergedArr[mergedIndex] = arr[start2]
		start2++
		mergedIndex++
	}
	copy(arr[originStart:originEnd], mergedArr)
}
