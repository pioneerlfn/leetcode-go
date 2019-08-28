package leetcode

import (
	"math/rand"
	"testing"
)

const N = 1000

func BenchmarkTwoSum1(b *testing.B) {
	nums := make([]int, 12)
	for i := 0; i < N; i++ {
		nums = append(nums, rand.Int())
	}
	nums = append(nums, 7, 2)


	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum1(nums, 9)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	nums := make([]int, 12)
	for i := 0; i < N; i++ {
		nums = append(nums, rand.Int())
	}
	nums = append(nums, 7, 2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum2([]int{11, 15, 2, 7}, 9)
	}
}