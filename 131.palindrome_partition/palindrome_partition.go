package main

import "fmt"

func isPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if !(s[i] == s[len(s)-1-i]) {
			return false
		}
	}
	fmt.Printf("%s is palindrome.\n", s)
	return true
}

// aab
// a, a, b
// [["a", "a". "b"], ["aa", "b"]]
// i = 1, partition("ab") = [["a", "b"]]
//
func main() {
	s := "aabbaa"
	res := partition(s)
	fmt.Println(res)
}


// cdd

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	var res [][]string
	for i := 1; i < len(s)+1; i++ {
		// 所有对res的append操作，必须以s[:i]回文为前提
		if isPalindrome(s[:i]) {
			// 检察边界条件，如果此时字符串遍历结束了，那么直接将s[:i] append到res中，
			// 且不再递归调用partition.否则tmp必为空，没有意义
			if i == len(s) {
				res = append(res, []string{s})
				break
			}
			// tmp不为空的情况
			tmp := partition(s[i:])
			for _, item := range tmp {
				current := append([]string{s[:i]}, item...)
				res = append(res, current)
			}
		}
	}
	return res
}

