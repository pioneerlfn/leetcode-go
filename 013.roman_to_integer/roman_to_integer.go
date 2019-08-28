package main

import "fmt"

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var res int
	var last int
	for i := len(s)-1; i >= 0; i-- {
		current := m[s[i]]
		sign := 1
		if current < last {
			sign = -sign
		}
		last = current
		res += sign * current
	}

	return res
}

func main() {
	s := "IV"
	res := romanToInt(s)
	fmt.Println("res: ", res)
}