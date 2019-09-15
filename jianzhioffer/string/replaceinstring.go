package string

import "strings"

// 题目：替换字符串中的空格为%20
// 思路：替换后的字符串变长，需要重新分配空间，先计算出所需的空间，然后通过两个指针，一个指向原始字符串的头部，
// 另一个指向新分配的空间的头部，从前往后复制，遇到空格放置新字符
func replaceBlank(str, newStr string) string {
	if str == "" {
		return str
	}
	count := strings.Count(str, " ")
	if count == 0 {
		return str
	}
	b := make([]byte, len(str)+count*(len(newStr)-1))
	j := 0
	for i := 0; i < len(str); i++ {
		if str[i] != 32 {
			b[j] = str[i]
			j++
		} else {
			copy(b[j:], newStr)
			j += len(newStr)
		}
	}
	return string(b)
}
