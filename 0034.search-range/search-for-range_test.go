package _034_search_range

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	nums2, target2 := []int{1}, 1

	target := 8
	indexs := SearchRange(nums, target)
	indexs2 := SearchRange(nums2, target2)
	fmt.Println(indexs, "\n", indexs2)
}
