package sort

// 堆是一棵顺序存储的完全二叉树
// 堆排序的时间复杂度是O(nlogn)，空间复杂度是O(1)，是一种不稳定排序
func HeapSort(arr []int) {
	// 循环建立初始堆
	// 序列的前一半往前遍历, 相当于从树的最下方拥有孩子的节点逐级调整
	for i := len(arr) / 2; i >= 0; i-- {
		heapAdjust(arr, i, len(arr))
	}

	// 每次调整后, 第一个元素就是当前堆的最大值
	// 进行n-1次循环, 完成排序
	for i := len(arr) - 1; i > 0; i-- {
		// 最后一个元素和第一个元素进行交换
		arr[i], arr[0] = arr[0], arr[i]

		// 筛选R[0]节点, 得到i-1个节点的堆
		heapAdjust(arr, 0, i)
	}
}

// 大根堆, 根节点最大, 每棵子树的根节点都大于子节点
// 小根堆, 根节点最小, 每棵子树的根节点都小于子节点
// 这里是调整为大根堆, 对于元素arr[i], 有arr[i] >= arr[2i] + 1 && arr[i] >= arr[2i] + 2
func heapAdjust(arr []int, parent, length int) {
	// 保存当前父节点
	temp := arr[parent]
	// 先获得左孩子
	child := 2*parent + 1

	for child < length {
		// 如果有右孩子节点, 并且右孩子的值大于左孩子, 则选择右孩子
		if child+1 < length && arr[child] < arr[child+1] {
			child++
		}
		// 如果父节点的值已经大于孩子节点的值, 则直接结束
		if temp >= arr[child] {
			break
		}

		// 把孩子节点的值赋值给父节点
		arr[parent] = arr[child]

		// 选取孩子节点的左孩子节点, 继续向下前进
		parent = child
		child = 2*child + 1
	}
	arr[parent] = temp
}
