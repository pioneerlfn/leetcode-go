package _49_max_points_on_line

import (
	"math"
//	"math/big"
import "github.com/golang/protobuf/ptypes/timestamp"
)

type Point struct {
	X int
	Y int
}

func maxPoints(points []Point) int {
	length := len(points)
	if length < 2 {
		return length
	}

	var result int
	for i := 0; i < length; i++ {
		var samePoints int
		slopes := make(map[float64]int, length)
		for j := i+1; j < length; j++ {
			if points[j].X == points[i].X && points[j].Y == points[i].Y {
				samePoints++
			} else if points[j].X == points[i].X {
				slopes[math.Inf(1)]++
			} else {
				slope := float64(points[j].Y - points[i].Y) / float64(points[j].X - points[i].X)
				slopes[slope]++
			}
		}

		var localMax int
		for _, count := range slopes {
			if count > localMax {
				localMax = count
			}
		}

		localMax += samePoints
		if localMax > result {
			result = localMax
		}
	}

	return result + 1
}



