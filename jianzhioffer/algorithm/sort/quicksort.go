package sort

import (
	"fmt"
)

// 快速排序，是一种交换排序，平均时间复杂度为O(nlogn)，最坏的时间复杂度为O(n^2)，最好的时间复杂度为O(nlogn)，
// 空间复杂度为O(nlogn)，是一种较为复杂的不稳定排序算法
// 时间复杂度：当数据有序时，以第一个元素为基准分为两个序列，前一个序列为空，此时执行效率最差，当数据随机分布时，
// 分成两个序列长度接近相等，此时执行效率最好。
// 空间复杂度：每次分割时需要一个空间存储基准值，而快速排序大概需要O(nlogn)次分割，所以空间复杂度为O(nlogn)
// 稳定性：在快速排序中，相等元素可能会因为分区而交换顺序，所以它是不稳定的算法。
func QuickSort(arr []int) {
	// 大于一个元素才进行排序
	if len(arr) > 1 {
		start, end := 0, len(arr)-1
		// 取第一个元素作为基准
		base := arr[start]
		fmt.Printf("start:%d, end:%d, base:%d, arr:%v\n", start, end, base, arr)

		for start < end {
			// 从尾部开始往前遍历, 直到遇到小于base, 或者到达start
			for arr[end] >= base && end > start {
				end--
			}
			// 判断当前到达的元素是否小于base, 如果是的话, 将当前元素放到start位置
			//if arr[end] < base {
			arr[start] = arr[end]
			//}
			// 从头部开始往后遍历, 直到遇到大于base的值, 或者到达end
			for arr[start] <= base && start < end {
				start++
			}
			// 判断当前到达的元素是否大于base, 如果是的话, 将当前元素放到end的位置
			//if arr[start] > base {
			arr[end] = arr[start]
			//}
		}
		arr[start] = base
		// 到这里, base左边都是比base小的数, base右边都是比base大的数, 接着把左边的数和右边的数进行相同的操作, 直到len(arr) <= 1
		QuickSort(arr[:start])
		QuickSort(arr[start+1:])
	}
}

// 快排的平均复杂度为O(logn),
func QuickSort2(arr []int) {
	if len(arr) > 1 {
		start, end := 0, len(arr)-1
		base := arr[start]
		for start < end {
			for arr[end] >= base && end > start {
				end--
			}
			arr[start] = arr[end]
			for arr[start] <= base && start < end {
				start++
			}
			arr[end] = arr[start]
		}
		arr[start] = base
		QuickSort2(arr[:start])
		QuickSort2(arr[start+1:])
	}
}
