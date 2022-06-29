// Package util go实用库
// RSA的签名及验签
// SHA256 With RSA
package util

import "math"

// Round 四舍五入
// decimals 精确到几位小数
func Round(v float64, decimals int) float64 {
	if math.IsNaN(v) {
		return math.NaN()
	}
	// 根据精确几位小数decimals, 先把小数部分前decimals变为整数,后面再取一位小数
	var pow float64 = 1
	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	vpow := v * pow

	// math.Modf函数返回vpow的整数部分和小数部分，结果的正负号和都x相同，比如1.1返回1.0和0.1，-1.1返回-1.0和-0.1
	intPart, fracPart := math.Modf(vpow)

	// If we're at the midpoint.
	m := 0.5
	if v < 0 {
		m = -0.5
	}
	if isWithin(fracPart, m) {
		if int64(intPart)%2 == 0 {
			// It's even, so round to it.
			return intPart / pow
		}
	}
	return float64(int64(vpow+(m*1.001))) / pow
}

func isWithin(a float64, b float64) bool {
	if a > b {
		return a-b <= 0.01
	}
	return b-a <= 0.01
}

// Floor 舍弃法
func Floor(v float64) float64 {
	i, f := math.Modf(v)
	if f >= 0.1 {
		return i - f
	}
	return i
}

// Ceil 进一法
func Ceil(v float64) float64 {
	i, f := math.Modf(v)
	if f >= 0.1 {
		return i + 1.0
	}
	return i
}
