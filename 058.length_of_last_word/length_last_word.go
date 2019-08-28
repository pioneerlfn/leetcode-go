package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	words := strings.Split(strings.Trim(s, " "), " ")
	fmt.Printf("words: %+q\n", words)
	return len(words[len(words)-1])
}

func main() {
	s := " a "
	fmt.Println(lengthOfLastWord(s))
}
