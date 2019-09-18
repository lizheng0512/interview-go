package stack

// 题目31: 输入两个整数序列, 一个为栈的压入序列, 判断另一个是否为栈的弹出序列, 假设压入栈的数字都不相等
func IsPopSequence(push, pop []int) bool {
	if (len(push) != len(pop)) || len(push) == 0 {
		return false
	}
	// 用一个辅助栈, 模拟压入弹出过程
	tempStack := New()
	pushIndex := 0
	for i := 0; i < len(pop); i++ {
		curPop := pop[i]
		top, ok := tempStack.Peek()
		if !ok || top.(int) != curPop {
			if pushIndex >= len(push) {
				return false
			}
			for push[pushIndex] != curPop {
				tempStack.Push(push[pushIndex])
				pushIndex++
			}
			// push序列已经走完了还没有找到curPop, 说明不是对应的弹出序列
			if push[pushIndex] != curPop {
				return false
			} else {
				pushIndex++
			}
		} else {
			tempStack.Pop()
		}
	}
	return true
}
