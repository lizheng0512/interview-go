package algorithm

import "math"

// 动态规划
// 如果一个面试题是求一个问题的最优解（通常是最大值或者最小值），而且该问题能够分解为若干个子问题，并且子问题之间还有重叠的小问题，
// 这种情况下往往可以通过动态规划解决。
// 特点：
// 1.求一个问题的最优解
// 2.整体问题的最优解依赖各个子问题的最优解
// 3.把大问题拆分为小问题，小问题之间还有相互重叠的更小的问题
// 4.为了避免重复子问题的求解，通常从下往上解决问题，把小问题的最优解存储下来，特点归纳为：从上到下分析问题，从下到上解决问题

// 题目14：剪绳子，给一段长度为n的绳子，把绳子剪成m段（m、n都是整数，并且m>1），求解每段绳子的最大乘机是多少。
func MaxProductAfterCutting(length int) int {
	switch length {
	case 0, 1:
		return 0
	case 2:
		return 1
	case 3:
		return 2
	}

	// 用于存放长度为下标的绳子的最优解
	products := make([]int, length+1)
	// 下列几个为特殊情况，不满足长度为下标，值为最优解
	products[0] = 0
	products[1] = 1
	products[2] = 2
	products[3] = 3

	// 从下到上求解
	for i := 4; i <= length; i++ {
		max := 0
		// 计算长度为i的每种情况
		for j := 0; j <= i/2; j++ {
			// 每次计算的因子都已存在于products中
			product := products[j] * products[i-j]
			if product > max {
				max = product
			}
		}
		products[i] = max
	}

	return products[length]
}

// 剪绳子的问题可以用贪婪算法来解决，绳子长度大于等于5时，尽可能剪出更多的长度为3的绳子，当绳子长度等于4时，剪成长度为2的绳子
// 这种规律是经过数学推导出来的，不能通用
func MaxProductAfterCutting2(length int) int {
	switch length {
	case 0, 1:
		return 0
	case 2:
		return 1
	case 3:
		return 2
	}

	// 尽可能多地剪出长度为3的绳子
	timesOf3 := length / 3

	// 当最后剩长度为4时，剪成两段长度为2的，不剪成长度为1和3的
	if length-timesOf3*3 == 1 {
		timesOf3--
	}

	timesOf2 := (length - timesOf3*3) / 2

	return int(math.Pow(3, float64(timesOf3)) * math.Pow(2, float64(timesOf2)))
}
