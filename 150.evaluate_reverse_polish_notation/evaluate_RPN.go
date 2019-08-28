package main

import (
	"log"
	"strconv"
)



func main() {
	tokens  := []string{"2","1","+","3","*"}
	result := evalRPN(tokens)
	log.Println(result)
}
func evalRPN(tokens []string) int {
	// remember to set len = 0.
	stack := make([]int, 0, len(tokens))
	// log.Println(len(tokens))
	for _, token := range tokens {
		if token == "+" ||
			token == "-" ||
			token == "*" ||
			token == "/" {
				b, a := stack[len(stack)-1], stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				tmp := compute(a, b, token)
				stack = append(stack, tmp)
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
			log.Println(stack)
		}
	}
	return stack[len(stack)-1] // or stack[0]
}

func compute(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		return a / b
	}
}



