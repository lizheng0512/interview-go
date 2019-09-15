package array

import "fmt"

// 题目：在长度为n的数组中找到重复的数字，数组中每个元素值的范围为0~n-1
// 思路：将每个数组元素移动到和值相同的下标位置，如果移动时发现目标位置的元素值和自己相同，则为重复元素,
// 如果不同，则将两值交换
func find1(arr []int) map[int]int8 {
	duplicateNumbersSet := make(map[int]int8)
	for i := 0; i < len(arr); {
		if arr[i] == -1 {
			i++
			continue
		}
		if arr[i] != i {
			if arr[arr[i]] == arr[i] {
				//fmt.Println("发现重复元素：", arr[i])
				duplicateNumbersSet[arr[i]] = 0
				arr[i] = -1
				i++
			} else {
				temp := arr[arr[i]]
				arr[arr[i]] = arr[i]
				arr[i] = temp
			}
		} else {
			i++
		}
	}
	//duplicateNumbers := make([]int, 0, len(set))
	//for k := range set {
	//	duplicateNumbers = append(duplicateNumbers, k)
	//}
	return duplicateNumbersSet
}

// 题目: 不修改原数组找出一个重复的数字, 数组长度是n+1, 其中的元素都在1~n的范围内, 所以数组中至少有一个元素是重复的
// 思路: 二分法, 即在每一个二分区间里去判断原数组中包含多少个区间内的元素, 如果大于二分区间长度, 则表示重复元素在这个区间内,
// 依次循环, 直到找到最小区间里的那个元素
// 特点: 此方法无法找到全部元素, 并且输入长度为n的数组, 时间复杂度为O(nlogn), countRange的执行次数为logn, 每次执行是n
// 空间复杂度为O(1)
func find2(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	start := 1
	end := len(arr) - 1
	// 计算原始数组中在start和end值之间的元素个数
	countRange := func(originArr []int, s, e int) (count int) {
		for _, v := range originArr {
			if v >= s && v <= e {
				count++
			}
		}
		return
	}
	for start <= end {
		middle := (end-start)>>1 + start
		count := countRange(arr, start, middle)
		if end == start {
			if count > 1 {
				return start
			} else {
				fmt.Println("无重复数字")
				break
			}
		}
		if count > (middle - start + 1) {
			end = middle
		} else {
			start = middle + 1
		}
	}
	return -1
}

// 题目同find2
// 此解法类似于find1, 开辟一个长度为n+1的数组, 用于放每个1~n的元素
// 特点: 此解法用空间复杂度替换时间复杂度, 空间复杂度为O(n), 并可以找出全部的重复元素
func find3(arr []int) (re []int) {
	if len(arr) <= 0 {
		panic("请输入长度大于0的切片")
	}
	tempArr := make([]int, 0, len(arr))
	for i := 0; i < cap(tempArr); i++ {
		tempArr = append(tempArr, -1)
	}
	duplicateNumbersSet := make(map[int]int8)
	for _, v := range arr {
		if tempArr[v] == v {
			duplicateNumbersSet[v] = 0
		} else {
			tempArr[v] = v
		}
	}
	for k := range duplicateNumbersSet {
		re = append(re, k)
	}
	return
}
