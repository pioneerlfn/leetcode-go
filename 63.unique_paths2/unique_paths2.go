package _3_unique_paths2

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid[0]), len(obstacleGrid)
	if obstacleGrid[n-1][m-1] == 1 {
		return 0
	}

	d := constructMatrix(n, m)
	d[n-1][m-1] = 1 // obstacleGrid[n-1][m-1] != 1.
	// initialize the column m-1.
	for i := n-2; i >= 0; i-- {
		if obstacleGrid[i][m-1] == 1 {
			d[i][m-1] = 0
		} else {
			d[i][m-1] = d[i+1][m-1]
		}
	}

	for i := n-1; i >= 0; i--{
		for j := m-2; j >= 0; j-- {
			if i == n-1 {
				d[i][j] = d[i][j+1]
			} else {
				d[i][j] = d[i][j+1] + d[i+1][j]
			}

			if obstacleGrid[i][j] == 1 {
				d[i][j] = 0
			}
		}
	}
	return d[0][0]
}

type Matrix [][]int

func constructMatrix(rows, cols int) Matrix {
	m := make([][]int, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]int, cols)
	}
	return m
}
