package sort

// 对公司全部员工年龄进行排序，要求时间复杂度为O(n)，可以使用辅助空间，但只能使用常量大小的辅助空间且不能超过O(n)
func Case1(ages []int) {
	if len(ages) <= 1 {
		return
	}
	// 只计算100岁以下的年龄
	oldestAge := 100
	// 下标代表年龄，元素值代表统计次数
	timesOfAges := make([]int, oldestAge, oldestAge)

	for _, age := range ages {
		timesOfAges[age]++
	}

	index := 0
	for age, times := range timesOfAges {
		for i := 0; i < times; i++ {
			ages[index] = age
			index++
		}
	}
}
