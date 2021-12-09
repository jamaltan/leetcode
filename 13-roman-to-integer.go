package main

/**
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

示例 1:

输入: s = "III"
输出: 3
示例 2:

输入: s = "IV"
输出: 4
示例 3:

输入: s = "IX"
输出: 9
示例 4:

输入: s = "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.
示例 5:

输入: s = "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.
**/

func romanToInt(s string) int {
	smap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	kInts := make([]int, 7)
	for c, i := range kInts {
		kInts = append(kInts, i)
	}
	sort.Ints(kInts)

}
