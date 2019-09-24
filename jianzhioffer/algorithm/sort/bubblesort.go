package sort

import "fmt"

// 冒泡排序
// 平均时间复杂度O(n^2), 最好时间复杂度O(n), 空间复杂度, 属于稳定排序算法
func BubbleSort(arr []int) {
	// 为了在原本排序好的序列上时间复杂度将为O(n)
	sorted := true
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				sorted = false
			}
		}
		fmt.Println("扫描次数: ", i+1)
		if sorted {
			break
		}
	}
}
