package leetcode

func TwoSum1(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TwoSum2(nums []int, target int) []int {
	n := len(nums)
	d := make(map[int]int, n)
	for i := 0; i < n; i++ {
		if index, present := d[target - nums[i]]; present {
			return []int{index, i}
		}
		d[nums[i]] = i
	}
	return nil
}

