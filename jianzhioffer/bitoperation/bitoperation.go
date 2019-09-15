package bitoperation

var m = map[byte]int{
	'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5, 'F': 6, 'G': 7, 'H': 8, 'I': 9, 'J': 10, 'K': 11, 'L': 12, 'M': 13, 'N': 14,
	'O': 15, 'P': 16, 'Q': 17, 'R': 18, 'S': 19, 'T': 20, 'U': 21, 'V': 22, 'W': 23, 'X': 24, 'Y': 25, 'Z': 26}

// 题目：在微软EXCEL中，A表示第一列，B表示第二列……Z表示第26列，AA表示第27列，写一个函数，输入字母表示的列号，输出是第几列
func Cal26ToDecimal(num26 string) int {
	decimal := 0
	for i := 0; i < len(num26); i++ {
		decimal *= 26
		if d, ok := m[num26[i]]; ok {
			decimal += d
		}
	}
	return decimal
}

// 题目：计算整数的二进制表示中1的个数
func Count1(num int) int {
	count := 0
	if num < 0 {
		count = 1
	}
	checkNum := 1
	max := 1 << 62
	for {
		if num&checkNum != 0 {
			count++
		}
		if checkNum == max {
			break
		} else {
			checkNum = checkNum << 1
		}
	}
	return count
}

// 最佳的count1方法
// 一个数和它减去一的结果做与运算，相当于将它二进制表示的最右边的为1的数位变成0
// 如1100，减去1为1011，1100 & 1011 = 1000，去掉了最右边的1
func Count1Optimum(num int) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
