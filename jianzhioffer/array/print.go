package array

import "fmt"

// 题目29: 输入一个矩阵, 按照从外向里顺时针的顺序依次打印每个数字
func PrintMatrixByCircle(matrix [][]int) {
	// 可以看成是从外向内一圈一圈地打印, 每圈打印的起点为(0,0),(1,1),(2,2)...
	// 循环打印, 循环继续的条件为X轴元素数>起点的x坐标*2, Y轴元素数>起点的Y坐标*2
	if matrix == nil {
		return
	}

	start := 0

	for len(matrix) > start*2 && len(matrix[0]) > start*2 {
		printMatrixByOneCircle(matrix, start)
		start++
	}
}

func printMatrixByOneCircle(matrix [][]int, start int) {
	x, y := start, start
	var step1, step2, step3 bool
	for x < len(matrix[0])-start {
		fmt.Println(matrix[y][x])
		x++
		step1 = true
	}
	if step1 {
		y++
		x--
		for y < len(matrix)-start {
			fmt.Println(matrix[y][x])
			y++
			step2 = true
		}
		if step2 {
			x--
			y--
			for x >= start {
				fmt.Println(matrix[y][x])
				x--
				step3 = true
			}
			if step3 {
				y--
				x++
				for y > start {
					fmt.Println(matrix[y][x])
					y--
				}
			}
		}
	}
}
