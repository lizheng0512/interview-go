package array

// 合并两个递增排序的数组为一个递增排序数组
func Merge(arr1, arr2 []int) []int {
	if len(arr1) == 0 {
		return arr2
	}
	if len(arr2) == 0 {
		return arr1
	}

	mergedArr := make([]int, len(arr1)+len(arr2))
	MergeCore(mergedArr, arr1, arr2, 0, 0)
	return mergedArr
}

func MergeCore(mergedArr, arr1, arr2 []int, index1, index2 int) {
	if index1 < len(arr1) && index2 < len(arr2) {
		if arr1[index1] < arr2[index2] {
			mergedArr[index1+index2] = arr1[index1]
			MergeCore(mergedArr, arr1, arr2, index1+1, index2)
		} else {
			mergedArr[index1+index2] = arr2[index2]
			MergeCore(mergedArr, arr1, arr2, index1, index2+1)
		}
	} else if index1 >= len(arr1) {
		// arr1提前遍历完了, arr2还没有遍历完
		for index2 < len(arr2) {
			mergedArr[index1+index2] = arr2[index2]
			index2++
		}
	} else if index2 >= len(arr2) {
		// arr2提前遍历完了, arr1还没有遍历完
		for index1 < len(arr1) {
			mergedArr[index1+index2] = arr1[index1]
			index1++
		}
	}
}
