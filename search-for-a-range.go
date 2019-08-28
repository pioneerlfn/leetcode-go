package _034_search_range

func searchRange(nums []int, target int) []int {
	index := search(nums, target)
	if index == -1 {
		return []int{-1, -1}
	}
	// find the left boundary.
	first := index
	for {
		f := search(nums[:first], target)
		if f == -1 {
			break
		}
		first = f
	}

	last := index
	for {
		l := search(nums[last+1:], target)
		if l == -1 {
			break
		}
		last += l + 1
	}
	return []int{first, last}
}


func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	var mid int
	for low <= high {
		mid = low + ((high - low) >> 2)
		switch {
		case nums[mid] < target:
			low = mid + 1
		case nums[mid] > target:
			high = mid - 1
		default:
			return mid
		}
		return -1
	}
}
