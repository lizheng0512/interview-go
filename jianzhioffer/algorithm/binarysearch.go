package algorithm

// 二分查找法，要求输入的序列是有序的
func BinarySearchRecursively(arr []int, search int) (int, bool) {
	if len(arr) == 0 || search < arr[0] || search > arr[len(arr)-1] {
		return 0, false
	}
	i := len(arr) / 2
	if arr[i] == search {
		return i, true
	} else if arr[i] < search {
		return BinarySearchRecursively(arr[i+1:], search)
	} else {
		return BinarySearchRecursively(arr[:i], search)
	}
}

func BinarySearchIteratively(arr []int, search int) (int, bool) {
	if len(arr) == 0 || search < arr[0] || search > arr[len(arr)-1] {
		return 0, false
	}
	start, end := 0, len(arr)
	for start >= 0 && end <= len(arr) && start < end {
		i := (end-start)/2 + start
		if arr[i] == search {
			return i, true
		}
		if arr[i] > search {
			end = i
		} else {
			start = i + 1
		}
	}
	return 0, false
}

// 如果有重复元素则返回第一个
func BinarySearchFirst(arr []int, target int) (int, bool) {
	if len(arr) <= 0 || target < arr[0] || target > arr[len(arr)-1] {
		return 0, false
	}
	start, end, mid := 0, len(arr)-1, 0
	for start < end {
		mid = (start + end) >> 1
		if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	if arr[start] == target {
		return start, true
	}
	return 0, false
}

// 有果有重复元素则返回最后一个
func BinarySearchLast(arr []int, target int) (int, bool) {
	if len(arr) <= 0 || target < arr[0] || target > arr[len(arr)-1] {
		return 0, false
	}
	start, end, mid := 0, len(arr), 0
	for start < end {
		mid = (start + end) >> 1
		if arr[mid] <= target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	if arr[start-1] == target {
		return start - 1, true
	}
	return 0, false
}
