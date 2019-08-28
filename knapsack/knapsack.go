package main

import "fmt"

func main() {
	val := []int{60, 100, 120}
	wt := []int{10, 20, 30}
	W := 50
	maxVal := knapsack(W, wt, val, 3)
	fmt.Println(maxVal)
}



func knapsack(W int, wt, val []int, n int)  int {

	dp := make([][]int, n+1)
	fmt.Println(len(dp))
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 0; i <= n; i++ {
		for w := 0; w <= W; w++ {
			if i == 0 || w == 0 {
				dp[i][w] = 0
			} else if wt[i-1] <= w {
				// 注意dp[i]对应与val[i-1]和wt[i-1].
				dp[i][w] = max(val[i-1] + dp[i-1][w-wt[i-1]], dp[i][w])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}
	return dp[n][W]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
