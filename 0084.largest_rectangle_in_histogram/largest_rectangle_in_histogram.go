package _084_largest_rectangle_in_histogram

func largestRectangleArea(heights []int) int {
	heights = append([]int{-2}, heights...)
	heights = append(heights, -1)
	size := len(heights)
	stack := make([]int, 1, size)
	end := 1
	res := 0

	for end < size {
		if heights[stack[len(stack)-1]] < heights[end] {
			stack = append(stack, end)
			end++
			continue
		}
		begin := stack[len(stack)-2]
		index := stack[len(stack)-1]
		height := heights[index]
		area := (end - begin - 1) * height
		res = max(res, area)
		stack = stack[:len(stack)-1]
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
