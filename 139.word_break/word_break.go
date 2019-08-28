package main

import "fmt"

// leetcode, [leet,code]
func main() {
	dict := []string{"leet", "code"}
	canBreak := wordBreak("leetcode", dict)
	fmt.Println(canBreak)

}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && inDict(s[j:i], wordDict) {
				dp[i] = true
				break
			}
		}
	}
	fmt.Println(dp)
	return dp[len(dp)-1]
}

func inDict(substr string, wordDict []string) bool {
	for _, word := range wordDict {
		if substr == word {
			return true
		}
	}
	return false
}
