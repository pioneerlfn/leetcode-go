package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	nums = nums[:1]
	fmt.Println(nums)
}
