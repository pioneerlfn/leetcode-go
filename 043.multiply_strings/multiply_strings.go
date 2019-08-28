package main

import (
	"fmt"
)

// Input: num1 = "123", num2 = "456"
// Output: "56088"

func multiply(num1 string, num2 string) string {
	if num1 == "" || num1 == "0" || num2 == "" || num2 == "0" {
		return "0"
	}

	digits := make([]int, len(num1) + len(num2))

	for i := len(num1)-1; i >= 0; i-- {
		for j := len(num2)-1; j >= 0; j-- {
			product := (num1[i] - '0') * (num2[j] - '0')
			p1, p2 := i + j, i + j + 1
			sum := int(product) + digits[p2]
			digits[p2] = sum % 10
			digits[p1] += sum / 10
		}
	}

	var res []byte
	for _, item := range digits {
		if !(item == 0 && len(res) == 0) {
			res = append(res, byte(item) + '0')
		}
	}

	return string(res)
}

func main() {
	num1 := "123456"
	num2 := "45678"
	fmt.Println("num1 * num2: ", multiply(num1, num2))
}
