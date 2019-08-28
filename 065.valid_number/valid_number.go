package main

import "fmt"

func isNumber(s string) bool {
	if s == "" {
		return false
	}
	var i int
	for ; i < len(s); i++ {
		// 跳过空字符
		if s[i] == ' ' {
			continue
		}
		// 第一个非空字符必须是[0-9], '+', '-', '.' 中的一个.
		if !(s[i] >= '0' && s[i] <= '9' || s[i] == '+' || s[i] == '-' || s[i] == '.') {
			return false
		}
		break
	}

	// 全空
	if i == len(s) {
		return false
	}

	var numExist, eExist, dotExist, spaceExist bool
	if s[i] == '.' {
		dotExist = true
		if i == len(s)-1 || !(s[i+1] >= '0' && s[i+1] <= '9') {
			return false
		}
	}

	if !(s[i] == '+' || s[i] == '-' || s[i] == '.') {
		numExist = true
	}

	for i++; i < len(s); i++ {
		// fmt.Printf("i: %d  s[i]: %q\n", i, s[i])
		if !(s[i] >= '0' && s[i] <= '9' || s[i] == 'e' || s[i] == '.' || s[i] == '-' || s[i] == '+' || s[i] == ' ') {
			return false
		}
		if spaceExist && s[i] != ' ' {
			return false
		}

		if s[i] >= '0' && s[i] <= '9' {
			numExist = true
			continue
		}
		switch s[i] {
		case 'e':
			if !numExist || eExist {
				return false
			}
			eExist = true
			if i == len(s)-1 || !(s[i+1] >= '0' && s[i+1] <= '9' || s[i+1] == '-' || s[i+1] == '+') {
				return false
			}
		case '-': {
			if s[i-1] != 'e'  {
				return false
			} else {
				if i == len(s)-1 || !(s[i+1] >= '0' && s[i+1] <= '9') {
					return false
				}
			}
		}

		case '+': {
			if s[i-1] != 'e'  {
				return false
			} else {
				if i == len(s)-1 || !(s[i+1] >= '0' && s[i+1] <= '9') {
					return false
				}
			}
		}


		case '.':
			if eExist || dotExist {
				return false
			}
			if (i == len(s)-1 || !(s[i+1] >= '0' && s[i+1] <= '9')) && !(s[i-1] >= '0' && s[i-1] <= '9') {
				return false
			}
			dotExist = true
		case ' ':
			spaceExist = true
		}
	}

	return true
}

func main() {
	//s := " 99e2.5"
	//s := "0"
	//s := " 0.1 "
	// s := "abc"
	// s := "1 a"
	// s := "2e10"
	// s := " -90e3   "
	//s := " 1e"
	//s := "e3"
	//s := " 6e-1"
	//s := " 99e2.5"
	//s := "53.5e93"
	//s := " --6 "
	//s := "-+3"
	//s := " +0e-"
	s := " 005047e+6"
	fmt.Println(isNumber(s))
}


