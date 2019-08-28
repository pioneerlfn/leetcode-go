package _2_unique_paths

func uniquePaths(m int, n int) int {
	d := make([]int, m)
	for j := 0; j < m; j++ {
		d[j] = 1
	}

	for i := n-2; i >= 0; i--{
		for j := m-2; j >= 0; j-- {
			d[j] = d[j] + d[j+1]
		}
	}
	return d[0]
}
