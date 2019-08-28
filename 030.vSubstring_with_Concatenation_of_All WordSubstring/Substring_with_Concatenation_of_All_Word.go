package main

import "fmt"

func findSubstring(s string, words []string) []int {
	if s == "" || words == nil || len(words) == 0 {
		return []int{}
	}
	n := len(words)
	m := len(words[0])
	wordsMap := make(map[string]int)

	var count int // count定义放到for循环外面
	var res []int
	for _, word := range words {
		count = wordsMap[word]
		count++
		wordsMap[word] = count
	}

	start, temp, end := 0, 0, m
	var wordsInWin int
	copyMap := make(map[string]int)
	// = 可以取
	for start <= len(s) - m*n {
		fmt.Println("\nstart: ", start, "temp:", temp,  "end: ", end, "copyMap:", copyMap, "wordInWind: ", wordsInWin)
		subStr := s[temp:end]
		count := wordsMap[subStr]
		copyCount := copyMap[subStr]
		fmt.Println("subStr: ", subStr, "\tcount: ", count, "copyCount: ", copyCount)


		if count != 0 && copyCount < count {
			copyMap[subStr]++
			temp = end
			end = temp + m
			wordsInWin++
		} else {
			start++
			copyMap = map[string]int{}
			wordsInWin = 0
			temp = start
			end = temp + m
			continue
		}

		if wordsInWin == n {
			fmt.Println("start: ", start, "temp:", temp,  "end: ", end, "copyMap:", copyMap, "wordInWind: ", wordsInWin)
			copyMap = map[string]int{}
			wordsInWin = 0
			res = append(res, start)
			start++
			temp = start
			end = temp + m
		}
	}

	return res
}

func main()  {
	// s := "barfoothefoobarman"
	// words := []string{"foo", "bar"}
//	 s := "wordgoodgoodgoodbestword"
//	words := []string{"word","good","best","word"}
	//s := "barfoofoobarthefoobarman"
	//words := []string{"bar","foo","the"}
	//
	 s := "lingmindraboofooowingdingbarrwingmonkeypoundcake"
	words := []string{"fooo","barr","wing","ding","wing"}
	// s := "ababaab"
	//words := []string{"ab","ba","ba"}
	///s := "aaaaaaaa"
	 //words := []string{"aa","aa","aa"}
	//s := "abaababbaba"
	//words := []string{"ab","ba","ab","ba"}
	res := findSubstring(s, words)
	fmt.Println("res: " ,res)
}
