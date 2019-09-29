package string

import "fmt"

// n为数字的长度，输出1~n位的所有数字
func PrintOneToNDigitNumber(n int) {
	if n <= 0 {
		return
	}
	// 初始化长度为n的字节数组，每一位置为字符'0'
	number := make([]byte, n)
	for i := 0; i < n; i++ {
		number[i] = '0'
	}
	for !increment(number) {
		printNumber(number)
	}
}

func increment(number []byte) bool {
	// 用于返回给调用方，表示是否打印完毕
	isOverflow := false
	nTakeOver := uint8(0)
	// 从低位到高位遍历，nTakeOver表示进位到i-1位的值
	for i := len(number) - 1; i >= 0; i-- {
		nSum := number[i] - '0' + nTakeOver
		if i == len(number)-1 {
			nSum++
		}
		if nSum >= 10 {
			if i == 0 {
				// 如果遍历到最高位，还发生了进位说明已经打印完毕，返回true中断PrintOneToNDigitNumber的循环
				isOverflow = true
			} else {
				// 进位之后当前位置为0，并且进位1
				number[i] = '0'
				nTakeOver = 1
			}
		} else {
			number[i] = '0' + nSum
			// 没有发生进位直接跳出循环
			break
		}
	}
	return isOverflow
}

// 打印数字，不打印前面的0
func printNumber(number []byte) {
	isBeginning0 := true
	for i := 0; i < len(number); i++ {
		if isBeginning0 && number[i] != '0' {
			isBeginning0 = false
		}
		if !isBeginning0 {
			fmt.Printf("%c", number[i])
		}
	}
	fmt.Print("\t")
}
