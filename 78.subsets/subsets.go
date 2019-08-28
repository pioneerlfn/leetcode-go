package main

import (
	"fmt"
	// "math"
)
func subsets(nums []int) [][]int {
	var result [][]int
	helper(nums, &result, []int{}, 0)
	return result
}

func helper(nums []int, result *[][]int, temp []int, index int) {
	t := make([]int, len(temp))
	copy(t, temp)
	*result = append(*result, t)
	for i := index; i < len(nums); i++ {
		temp = append(temp, nums[i])
		// 注意，要传temp的拷贝
		helper(nums, result, temp, i+1)
		// 回退
		temp = temp[:len(temp)-1]
	}
}



func main() {
	nums := []int{1,2,3}
	res := subsets(nums)
	fmt.Println(res)
}