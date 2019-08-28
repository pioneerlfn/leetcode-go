package main

import (
	"fmt"
	"strings"
)

// "/a/./b/../../c/"

func simplifyPath(path string) string {
	parts := strings.Split(path, "/")
	fmt.Printf("parts: %+q\n", parts)
	var stack []string
	for i := 0; i < len(parts); i++ {
		fmt.Printf("stack: %+q\n", stack)
		if parts[i] == "" || parts[i] == "." {
			continue
		} else if parts[i] == ".." {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, parts[i])
		}
	}

	var res string
	for i := 0; i < len(stack); i++ {
		res += "/" + stack[i]
	}
	if len(res) == 0 {
		return "/"
	}
	return res
}

func main() {
//	s := "/a/./b/../../c/"
//	s2 := "/a//b////c/d//././/.."
	s3 := "/../"
// 	simplePath := simplifyPath(s)
//	simplePath2 := simplifyPath(s2)
	simplePath3 := simplifyPath(s3)
//	fmt.Println(simplePath)
//	fmt.Println(simplePath2)
	fmt.Println(simplePath3)
}
