package main

import "fmt"

// '?' Matches any single character.
// '*' Matches any sequence of characters (including the empty sequence).
// s could be empty and contains only lowercase letters a-z.
// p could be empty and contains only lowercase letters a-z, and characters like ? or *.

// s = "a d c e b"
//      0 1 2 3 4
// p = "* a * b", star = 2, match = 3, pp = 3, sp = 3
// Output: true

//    sp:  0 1 2
// match:    1 2
//        "a c d c b"
//         0 1 2 3 4
//        "a * c ? b"
//  star:    1
//    pp:  0 1 2

// For each element in s
// If *s==*p or *p == ? which means this is a match, then goes to next element s++ p++.
// If p=='*', this is also a match, but one or many chars may be available, so let us save this *'s position and the matched s position.
// If not match, then we check if there is a * previously showed up,
//       if there is no *,  return false;
//       if there is an *,  we set current p to the next element of *, and set current s to the next saved s position.

func isMatch(s string, p string) bool {
	var sp, pp int
	var match int
	star := -1

	// 只用sp循环，这样循环结束的时候，我们能知道s已经遍历完了
	for sp < len(s) {
		if pp < len(p) && (s[sp] == p[pp] || p[pp] == '?') {
			sp++
			pp++
			continue
		} else if pp < len(p) && p[pp] == '*' {
			star = pp
			match = sp
			pp++ // 注意，是pp++
		} else if  star != -1 {
			pp = star + 1 // 这句不能丢，pp和sp都可能会回溯
			match++
			sp = match
		} else {
			return false
		}
	}

	// 处理p未结束的情况
	for pp < len(p) && p[pp] == '*' {
		pp++
	}

	return pp == len(p)
}

func main() {
	s := "aa"
	p := "*"
	// s := "cb"
	// p := "?a"
	// s := "a d c e b"
	// p := "*a*b"
	match := isMatch(s, p)
	fmt.Println(match)

}