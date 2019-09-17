package string

// 题目：实现一个函数用于正则表达式匹配，pattern中的'.'表示一个字符，'*'表示它前面的字符可以出现多次或者0次
//func Match(str, pattern string) (bool, error) {
//	if len(str) == 0 || len(pattern) == 0 {
//		return false, errors.New("请输入正确的参数")
//	}
//	if pattern[0] == '*' {
//		return false, errors.New("请输入正确的正则表达式")
//	}
//
//	strIndex, patternIndex := 0, 0
//	for strIndex < len(str) {
//		if patternIndex >= len(pattern) {
//			return false, nil
//		}
//		switch pattern[patternIndex] {
//		case '.':
//			// 判断后面是不是有'*'
//			if patternIndex+1 < len(pattern) && pattern[patternIndex+1] == '*' {
//				patternIndex += 2
//				if patternIndex >= len(pattern) {
//					return true, nil
//				} else {
//					for patternIndex < len(pattern) && pattern[patternIndex] != '*' && pattern[patternIndex] != '.'
//					for strIndex < len(str) && str[strIndex] != patter
//				}
//			}
//			strIndex++
//			patternIndex++
//		default:
//			if str[strIndex] == pattern[patternIndex] {
//				// 判断后面是不是有'*'
//				if patternIndex+1 < len(pattern) && pattern[patternIndex+1] == '*' {
//					patternIndex += 2
//					curValue := str[strIndex]
//					for strIndex < len(str) && str[strIndex] == curValue  {
//						strIndex ++
//					}
//				} else {
//					strIndex++
//					patternIndex++
//				}
//			} else {
//				// 判断后面是不是有'*'
//				if patternIndex+1 < len(pattern) && pattern[patternIndex+1] == '*' {
//					patternIndex += 2
//				} else {
//					return false, nil
//				}
//			}
//		}
//	}
//	if patternIndex < len(pattern) {
//		return false, nil
//	}
//	return true, nil
//}
func Match(str, pattern string) bool {
	if len(str) == 0 || len(pattern) == 0 {
		return false
	}
	return matchCore(str, pattern, 0, 0)
}

func matchCore(str, pattern string, s, p int) bool {
	if s == len(str) && p == len(pattern) {
		return true
	}

	// 模式先到末尾, 但是文本没有到
	if p == len(pattern) && s < len(str) {
		return false
	}

	// 第二个字符是'*'
	if p+1 < len(pattern) && pattern[p+1] == '*' {
		if s < len(str) && (str[s] == pattern[p] || pattern[p] == '.') {
			return matchCore(str, pattern, s, p+2) || matchCore(str, pattern, s+1, p+2) ||
				matchCore(str, pattern, s+1, p)
		} else {
			return matchCore(str, pattern, s, p+2)
		}
	}

	// 第二个字符不是'*'
	if s < len(str) && (str[s] == pattern[p] || pattern[p] == '.') {
		return matchCore(str, pattern, s+1, p+1)
	}
	return false
}
