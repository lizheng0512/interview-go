package array

// 题目: 调整数组顺序, 一分为二使满足f的在数组后面, 不满足f的在数组前面
func Reorder(arr []int, f func(int) bool) {
	if len(arr) <= 1 {
		return
	}

	start, end := 0, len(arr)-1
	for start < end {
		for start < end && !f(arr[start]) {
			start++
		}
		for start < end && f(arr[end]) {
			end--
		}
		if start < end {
			temp := arr[start]
			arr[start] = arr[end]
			arr[end] = temp
		}
	}
}
