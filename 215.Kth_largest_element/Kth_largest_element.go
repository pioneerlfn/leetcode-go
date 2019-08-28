package _15_Kth_largest_element

const cutoff  = 10


// quickselect.
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, k, 0, len(nums)-1)
}

func quickSelect(nums []int, k, left, right int) int {
	if left + cutoff <= right {

	} else {
		insertSort(nums, left, )
	}










}
