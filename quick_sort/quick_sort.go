package main

import "fmt"

const critical int = 10

func Qsort(nums []int)  {
	length :=  len(nums)
	qsort(nums, 0, length-1)
}

func qsort(nums []int, start, end int) {
	length := end -start + 1
	if length < critical {
		insertSort(nums[start:], length)
		return
	}
	pivot := getPivot(nums, start, end)
	i, j := start, end-1
	for {
		for i++; nums[i] < pivot; {
			i++
		}
		for j--; nums[j] > pivot; {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			// swap(&nums[i], &nums[j])
		} else {
			break
		}
	}
	nums[i], nums[end-1] = nums[end-1], nums[i]
	//swap(&nums[i], &nums[end-1])
	qsort(nums, start, i-1)
	qsort(nums, i+1, end)
}

func getPivot(nums []int, left, right int) int {
	center := (left + right) / 2
	if nums[left] > nums[center] {
		swap(&nums[left], &nums[center])
	}
	if nums[center] > nums[right] {
		swap(&nums[center], &nums[right])
	}
	if nums[left] > nums[center] {
		swap(&nums[left], &nums[center])
	}

	swap(&nums[center], &nums[right-1])
	return nums[right-1]
}

func swap(a, b *int)  {
	tmp := *a
	*a = *b
	*b = tmp
}

func insertSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1

		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = value
	}
}

func main() {
	nums := []int{85, 2, 4, 8, 1, 3, 11, 7, 22, 45, 11, 24, 93}
	fmt.Println(nums)
	Qsort(nums)
	fmt.Println(nums)
}