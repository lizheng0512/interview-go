package algorithm

import "fmt"

// 题目11：把一个数组最开始的若干个元素搬到末尾，称之为旋转数组，输入一个递增排序数组的旋转，找出这个数组的最小值
// 思路：利用两个指针指向旋转数组的开头和末尾，每次求两个指针中间的元素，若大于前面的指针，则表示最小值在中间元素的后面，
// 将前面的指针移动到中间元素，若小于前面的指针，则表示最小值在中间元素的前面或者中间元素就是最小值，将后面的指针移动到中间元素，
// 循环执行，直到两个指针相邻，此时后面的指针指向的元素是最小值。
// 通常这种前提下所有的旋转数组的第一个数字大于等于最后一个数字，当第一个数字小于最后一个数字的时候，说明旋转数组是完全升序的，
// 最小值为第一个数字，此外，不排除在移动过程中遇到两个指针和中间元素都相等的情况，这种情况下只能顺序遍历
func FindMinimumInRotateArray(arr []int) int {
	if len(arr) == 0 {
		panic("数组长度为0")
	}
	if len(arr) == 1 {
		return arr[0]
	}

	start, end := 0, len(arr)-1

	// 这种情况下为已排好序的数组，返回第一个
	if arr[start] < arr[end] {
		return arr[start]
	}

	for start+1 < end {
		middle := start + (end-start)/2
		if arr[start] == arr[end] && arr[start] == arr[middle] {
			fmt.Printf("need to iterate from %d to %d\n", start, end)
			// 这种情况下需要从start到end遍历数组
			min := arr[start]
			for start <= end {
				if arr[start] < min {
					min = arr[start]
				}
				start++
			}
			return min
		}
		if arr[start] <= arr[middle] {
			start = middle
		} else if arr[middle] <= arr[end] {
			end = middle
		}
	}
	return arr[end]
}
