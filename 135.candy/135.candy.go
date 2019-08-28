package _35_candy

func candy(ratings []int) int {
	candys := make([]int, len(ratings))
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candys[i] = candys[i-1] + 1
		}
	}

	for i := len(ratings)-2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candys[i] = max(candys[i+1] + 1, candys[i])
		}
	}

	var res int
	for i := 0; i < len(candys); i++ {
		res += candys[i] + 1
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}