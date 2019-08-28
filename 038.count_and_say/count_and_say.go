package main

import (
	"fmt"
)

const start int = 1

func countAndSay(n int) string {
	newArray := []int{start}
	var old []int

	for i := n-1; i > 0; i-- {
		old = newArray
		newArray = nil
		count := 1
		for j := 0; j < len(old); j++ {
			value := old[j]
			if j < len(old) -1 && old[j] == old[j+1] {
					count++
					continue
			}
			newArray = append(newArray, count, value)
			count = 1
		}
	}

	var res []byte
	for _, item := range newArray {
		res = append(res, byte(item) + '0')
	}
	return string(res)
}

func main() {
	fmt.Println(countAndSay(5))
}