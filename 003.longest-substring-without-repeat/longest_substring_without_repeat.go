package _03_longest_substring_without_repeat

// 解法1：滑动窗口
func lengthOfLongestSubstring(s string) int {
	// 处理边界情况
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	// 利用dict
	// 不用[26]int数组的原因是
	// string可能是"    "这样的，
	// 用' ' - 'a'会数组越界
	// ' '是32,'a'是97, ' '-'a' = 0xbf(ie:-65)
	// 但是如果将其置于数组下标，则会被解释为无符号数191
	charPos := make(map[byte]int)
	res, left := 0, -1
	for i := 0; i < len(s); i++ { // i:滑动窗口右边界移动
		if idx, ok := charPos[s[i]]; ok {
			if idx >= left {
				// 注意,实际窗口不包含左边界
				// 所以这里不是left=idx+1
				left = idx // left:滑动窗口坐边界移动
			}
		}
		charPos[s[i]] = i // insert or update
		if i-left > res {
			res = i - left
		}
	}
	return res
}

// 解法2：dp
// dp[i]定义为以第i个字符结尾的不包含重复字符的子字符串的最大长度
func dpApproch(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	dp := make([]int, len(s))
	dp[0] = 1
	max := 1
	m := make(map[byte]int)
	m[s[0]] = 0
	for i := 1; i < len(s); i++ {
		if idx, ok := m[s[i]]; !ok {
			dp[i] = dp[i-1] + 1
		} else {
			if i-idx <= dp[i-1] {
				dp[i] = i - idx
			} else {
				dp[i] = dp[i-1] + 1
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
		m[s[i]] = i
	}

	return max
}
