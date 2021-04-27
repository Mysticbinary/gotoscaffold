package gtsmath

// 切片求和
func SumSlice(s []int) int {
	var sum int
	for _, value := range s {
		sum += value
	}
	return sum
}

// 求切片平均值
func AverageSlice(s []int) float64 {
	sum := SumSlice(s)
	return float64(sum) / float64(len(s))
}


