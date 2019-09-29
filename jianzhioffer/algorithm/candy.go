package algorithm

func Candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	if len(ratings) == 1 {
		return 1
	}
	candis := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		num := 1
		for j := i + 1; j < len(ratings) && ratings[j] < ratings[j-1]; j++ {
			num++
		}
		if i-1 >= 0 && ratings[i] > ratings[i-1] && num <= candis[i-1] {
			candis[i] = candis[i-1] + 1
		} else {
			candis[i] = num
		}
		num--
		i++
		for num >= 1 {
			candis[i] = num
			i++
			num--
		}
		i--
	}
	count := 0
	for i := 0; i < len(candis); i++ {
		count += candis[i]
	}
	return count
}
