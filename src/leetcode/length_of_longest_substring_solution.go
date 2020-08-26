package leetcode

func LengthOfLongestSubstring(s string) int {
	window := make(map[rune]int)
	runeS := []rune(s)
	length := len(runeS)

	left, right, res := 0, 0, 0
	var c rune
	for right < length {
		c = runeS[right]
		right++
		window[c]++
		for window[c] > 1 {
			window[runeS[left]]--
			left++
		}
		res = max(res, right-left)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
