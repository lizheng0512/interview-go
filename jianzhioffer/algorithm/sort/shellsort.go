package sort

// 希尔排序, 是一种加强版的插入排序算法, 每次执行插入排序之前先定义步长, 每隔特定步长的元素为一组,
// 将同组内的元素进行插入排序, 直至步长为1
// 平均时间复杂度为O(nlogn), 最好的情况下为O(n^(1.3)), 空间复杂度为O(1), 是一种不稳定的排序算法
func ShellSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	// 定义初始步长为长度的一半, 之后每次除以2
	gap := len(arr) / 2
	for gap > 0 {
		for i := gap; i < len(arr); i++ {
			cur := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > cur {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = cur
		}
		gap /= 2
	}
}
