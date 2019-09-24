package sort

// 简单选择排序, 每次从剩下的序列中选出最小的元素放到已排序的序列的末尾
func SelectionSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 0; i < len(arr); i++ {
		mIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[mIndex] {
				mIndex = j
			}
		}
		arr[i], arr[mIndex] = arr[mIndex], arr[i]
	}
}
