package _26_remove_duplicates

func removeDuplicates(nums []int) int {
	count, i, size := 0, 0, len(nums)
	for ; i < size; i++ {
		if nums[count] != nums[i] {
			nums[count+1] = nums[i]
			count++
		}
	}

	return count+1
}
