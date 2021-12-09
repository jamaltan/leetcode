package main

// 找到重复就移动left
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	cmap := make(map[rune]int)
	var left int = -1
	var max int

	for i, c := range s {
		if j, exists := cmap[c]; exists { //有重复
			if j > left {
				left = j
			}
		}
		cmap[c] = i
		if size := i - left; size > max {
			max = size
		}
	}

	return max
}
