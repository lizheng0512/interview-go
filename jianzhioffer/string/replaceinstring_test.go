package string

import "testing"

var replaceBlackTests = []struct {
	p1, p2, re string
}{
	{"1 2 3 4 ", "%20", "1%202%203%204%20"},
	{"1234", "%20", "1234"},
	{"", "%20", ""},
	{" 1 2 3", "", "123"},
	{" 你 我 2 他 ", "！", "！你！我！2！他！"},
}

func TestReplaceBlank(t *testing.T) {
	for _, tt := range replaceBlackTests {
		if re := replaceBlank(tt.p1, tt.p2); re != tt.re {
			t.Errorf("结果错误！结果应该为：\"%s\"，实为：\"%s\"", tt.re, re)
		} else {
			t.Logf("输入\"%s\"，\"%s\"，输出结果\"%s\"，正确", tt.p1, tt.p2, re)
		}
	}
}
