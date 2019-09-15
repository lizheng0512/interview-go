package numeric

// 计算x的y次幂，不需要考虑大数问题
func Power(x float64, y int) float64 {
	if x == 0 || x == 1 || y == 0 {
		return 1
	}
	if y == 1 {
		return x
	}
	// 考虑指数为负数的情况
	unsignedY := uint(0)
	if y < 0 {
		unsignedY = uint(-y)
	} else {
		unsignedY = uint(y)
	}
	power := powerWithUnsignedExponent(x, unsignedY)
	// 指数为负数时，取倒数
	if y < 0 {
		power = 1 / power
	}
	return power
}

func powerWithUnsignedExponent(x float64, exponent uint) float64 {
	switch exponent {
	case 0:
		return 1
	case 1:
		return x
	}
	// x的y次幂等于x的y/2次幂的平方
	power := powerWithUnsignedExponent(x, exponent>>1)
	power *= power
	// 如果指数是奇数，还需要再乘以1次底数
	if exponent&0x01 != 0 {
		power *= x
	}
	return power
}
