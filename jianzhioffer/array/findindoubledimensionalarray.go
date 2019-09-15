package array

import (
	"fmt"
)

// 题目：在二维数组中查找数字，二维数组每一行从左到右递增，每一列从上到下递增
// 思路：从右上角/左下角开始找，如果该元素小于要查找的元素，则排除掉这一行，要查找的元素可能出现在下面的所有行，
// 如果大于要查找的元素，则排除这一列，要查找的元素可能出现在左侧的所有列，排除掉后，继续按照此规则在剩余的二维数组中查找，
// 直到找到该元素，或者遍历完毕
func findInDoubleDimensionalArray(arr [][]int, goal int) {
	if len(arr) <= 0 || len(arr[0]) <= 0 {
		panic("请输入至少包含一个元素的二维数组")
	}
	x, y, found := len(arr[0])-1, 0, false
	for x >= 0 && y < len(arr) {
		n := arr[y][x]
		if n < goal {
			y++
		} else if n > goal {
			x--
		} else {
			fmt.Printf("数字%d在第%d行第%d列\n", goal, y+1, x+1)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("没有找到数字%d\n", goal)
	}
}
