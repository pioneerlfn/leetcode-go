package _034_search_range

func SearchRange(nums []int, target int) []int {
	length := len(nums)
	left, right := 0, length - 1
	for left <= right  {
		mid := left + ((right - left) >> 2)
		if nums[mid] == target {
			start, end := mid, mid
			for start >= 0 && nums[start] == target {
				start--
			}
			start += 1
			for  end <= length - 1 && nums[end] == target {
				end++
			}
			end -= 1
			return []int{start, end}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return []int{-1, -1}
}
