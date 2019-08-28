package _035_search_insert_position

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + ((right - left) >> 2)
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			if mid == len(nums)-1 || nums[mid+1] >= target {
				return mid + 1
			} else {
				left = mid + 1
			}
		} else {
			return mid
		}
	}
	return left

}

