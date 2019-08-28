package _033_search_in_rotated_sorted_array

func search(nums []int, target int) int {
	size := len(nums)
	left, right := 0, size-1
	for left <= right {
		mid := left + ((right - left) >> 2)
		if nums[mid] == target {
			return mid
		}
		if nums[mid] <= nums[right] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
