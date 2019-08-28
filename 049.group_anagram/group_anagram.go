package main

import "fmt"

// Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
//Output:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	ans := make(map[[26]byte][]string)
	for _, item := range strs {
		var count [26]byte
		for _, char := range item {
			count[char-'a'] += 1
		}
		ans[count] = append(ans[count], item)
	}

	var res [][]string
	for _, value := range ans {
		res = append(res, value)
	}

	return res
}

func main() {
	for _, char := range "eat" {
		fmt.Println(char - 'a')
	}

	s :=  []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(s)
	fmt.Printf("res: %+q\n", res)
}