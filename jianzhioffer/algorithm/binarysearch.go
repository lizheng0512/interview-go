package algorithm

// 二分查找法，要求输入的序列是有序的
func BinarySearchRecursively(arr []int, search int) (result bool) {
	if len(arr) == 0 || search < arr[0] || search > arr[len(arr)-1] {
		return
	}
	i := len(arr) / 2
	if arr[i] < search {
		return BinarySearchRecursively(arr[i+1:], search)
	}
	if arr[i] > search {
		return BinarySearchRecursively(arr[:i], search)
	}
	if arr[i] == search {
		result = true
	}
	return
}

func BinarySearchIteratively(arr []int, search int) (result bool) {
	if len(arr) == 0 || search < arr[0] || search > arr[len(arr)-1] {
		return
	}
	start, end := 0, len(arr)
	for start >= 0 && end <= len(arr) && start < end {
		i := (end-start)/2 + start
		if arr[i] == search {
			result = true
			break
		}
		if arr[i] > search {
			end = i
		} else {
			start = i + 1
		}
	}
	return
}
