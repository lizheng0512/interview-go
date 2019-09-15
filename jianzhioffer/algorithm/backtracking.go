package algorithm

// 用回溯法解决问题

// 题目：判断矩阵中是否包含指定的路径，每个位置只能走一次
func FindPathInMatrix(matrix [][]string, path string) bool {
	if path == "" {
		return true
	}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rowLength := len(matrix)
	colLength := len(matrix[0])
	// 用于记录哪些位置已经走过
	visited := make([][]bool, 0, colLength)
	for i := 0; i < colLength; i++ {
		visited = append(visited, make([]bool, rowLength))
	}
	// 尝试每一个起点
	for row := 0; row < rowLength; row++ {
		for col := 0; col < colLength; col++ {
			if findPathInMatrixCore(matrix, path, 0, row, col, visited) {
				return true
			}
		}
	}
	return false
}

func findPathInMatrixCore(matrix [][]string, path string, pathIndex int, row, col int, visited [][]bool) bool {
	// 判断是否走完
	if len(path) == pathIndex {
		return true
	}

	hasPath := false
	if row < len(matrix) && row >= 0 && col < len(matrix[0]) && col >= 0 &&
		visited[row][col] == false && matrix[row][col] == string(path[pathIndex]) {
		visited[row][col] = true
		pathIndex++
		hasPath = findPathInMatrixCore(matrix, path, pathIndex, row-1, col, visited) ||
			findPathInMatrixCore(matrix, path, pathIndex, row, col+1, visited) ||
			findPathInMatrixCore(matrix, path, pathIndex, row+1, col, visited) ||
			findPathInMatrixCore(matrix, path, pathIndex, row, col-1, visited)
		// 如果下一步不符合要求的路径，则返回上一步
		if !hasPath {
			//pathIndex--
			visited[row][col] = false
		}
	}
	return hasPath
}

// 题目：地上有一个m行n列的方格，一个机器人从坐标（0,0）开始走，每次可以从上下左右方向移动一格，
// 但不能进入行坐标和列坐标的数位之和大于特定值的格子，如特定值为18，则机器人可以进入坐标为(35,37)，但不能进入(35,38)，
// 因为3+5+3+8=19，大于18，问机器人能够到达多少个格子
func MovingCount(rows, cols int, threshold int) int {
	if threshold <= 0 || rows <= 0 || cols <= 0 {
		return 0
	}

	visited := make([][]bool, 0, rows)
	for i := 0; i < rows; i++ {
		visited = append(visited, make([]bool, cols))
	}

	return movingCountCore(threshold, 0, 0, rows, cols, visited)
}

func movingCountCore(threshold, row, col, rows, cols int, visited [][]bool) int {
	count := 0
	if row >= 0 && row < rows && col >= 0 && col < cols &&
		visited[row][col] == false && getDigitSum(row)+getDigitSum(col) <= threshold {
		visited[row][col] = true
		count = 1 + movingCountCore(threshold, row-1, col, rows, cols, visited) +
			movingCountCore(threshold, row, col+1, rows, cols, visited) +
			movingCountCore(threshold, row+1, col, rows, cols, visited) +
			movingCountCore(threshold, row, col-1, rows, cols, visited)
	}
	return count
}

// 获取数字num的各个数位之和
func getDigitSum(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
