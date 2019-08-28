package main

import "fmt"

var target  = "catsanddog"
var dict  = []string{"cat", "cats", "dog", "and", "sand"}

func main() {
	res := wordBreak(target, dict)
	fmt.Printf("%q\n", res)
}

func wordBreak(s string, wordDict []string) []string {
	return dfs(s, wordDict, 0)
}

func dfs(s string, wordDict []string, start int) []string {
	if s == "" {
		return []string{""}
	}
	cache := make(map[int][]string)
	var res []string
	for i := start; i < len(s)+1; i++ {
		if words, ok := cache[start]; ok {
			fmt.Println("exists...")
			return words
		}
		var words []string
		if inDict(s[:i], wordDict) {
			partialStrs := dfs(s, wordDict, i)
			for _, partialStr := range partialStrs {
				if partialStr != "" {
					partialStr = " " + partialStr
				}
				words = append(words, s[:i] + partialStr)
			}
			// fmt.Println("s[:i]: ", s[:i])
			// fmt.Println("words: ", words)


			res = append(res, words...)
		}
	}
	cache[start] = res
	return res
}

func inDict(substr string, wordDict []string) bool {
	for _, word := range wordDict {
		if substr == word {
			return true
		}
	}
	return false
}
