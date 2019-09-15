package array

import (
	"fmt"
	"testing"
)

var tests = []struct {
	in  []int
	out []int
}{
	{[]int{2, 3, 1, 6, 4, 0, 5, 7}, []int{}},
	{[]int{7, 5, 2, 5, 1, 2, 4, 3, 3, 2}, []int{2, 3, 5}},
	{nil, []int{}},
}

func TestMain(m *testing.M) {
	fmt.Println("initializing...")
	m.Run()
	fmt.Println("Exiting...")
}

func TestFind1(t *testing.T) {
	for _, tt := range tests {
		re := find1(tt.in)
		for _, o := range tt.out {
			if _, ok := re[o]; !ok {
				t.Errorf("没有找到 %d", o)
			}
		}
		t.Logf("在%v中找到了重复数字%v", tt.in, tt.out)
	}
}

func BenchmarkFind1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		find1(tests[1].in)
		//fmt.Println(tests[1].out)
	}
}
