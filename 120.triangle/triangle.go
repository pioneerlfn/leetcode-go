package _20_triangle



// bottom-up DP.
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	dp := make([]int, len(triangle))
	for i := len(triangle)-1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if i == len(triangle)-1 {
				dp[j] = triangle[i][j]
				continue
			}
			dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
		}
	}
	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
