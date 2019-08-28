package main

import "fmt"

func numTrees(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	nums := make([]int, n+1, n+1)
	nums[0], nums[1] = 1, 1
	for i := 2; i < n+1; i++ {
		for j := 1; j <= i; j++ {
			nums[i] += nums[j-1] * nums[i-j]
		}
	}
	return nums[n]
}

func main() {
	res :=numTrees(3)
	fmt.Println(res)
}