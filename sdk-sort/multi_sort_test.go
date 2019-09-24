package sdk_sort

import (
	"fmt"
	"sort"
	"testing"
)

type AccessLog struct {
	Username   string
	AccessTime string
	ExistTime  string
}

type lessFunc func(a, b *AccessLog) bool

type accessLogSorter struct {
	logs []AccessLog
	less []lessFunc
}

func (als *accessLogSorter) Sort(logs []AccessLog) {
	als.logs = logs
	sort.Sort(als)
}

func (als *accessLogSorter) Len() int {
	return len(als.logs)
}

func (als *accessLogSorter) Swap(i, j int) {
	als.logs[i], als.logs[j] = als.logs[j], als.logs[i]
}

func (als *accessLogSorter) Less(i, j int) bool {
	x, y := &als.logs[i], &als.logs[j]
	k := 0
	for k < len(als.less)-1 {
		switch {
		case als.less[k](x, y):
			return true
		case als.less[k](y, x):
			return false
		}
		k++
	}
	return als.less[k](x, y)
}

func OrderBy(less ...lessFunc) *accessLogSorter {
	return &accessLogSorter{less: less}
}

func TestMultiSort(t *testing.T) {
	usernameSort := func(a, b *AccessLog) bool {
		return a.Username < b.Username
	}
	accessTimeSort := func(a, b *AccessLog) bool {
		return a.AccessTime < b.AccessTime
	}
	existTimeSort := func(a, b *AccessLog) bool {
		return a.ExistTime > b.ExistTime
	}
	logs := []AccessLog{
		{"lizheng", "11:11:11", "11:12:11"},
		{"lizheng", "12:11:11", "12:12:00"},
		{"lizheng", "12:11:11", "12:12:12"},
		{"zhangsan", "11:11:11", "12:12:12"},
	}
	OrderBy(usernameSort, accessTimeSort, existTimeSort).Sort(logs)
	fmt.Println(logs)
}
