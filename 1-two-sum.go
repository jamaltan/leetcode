package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// 蛮力枚举
func twoSum(nums []int, target int) []int {
	right := len(nums)
	res := make([]int, 0)

	for i := 0; i < right-1; i++ {
		for j := i + 1; j < right; j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i, j)

			}

		}

	}
	return res
}

// hash降低一层时间复杂度
func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	maps := make(map[int]int, 0)

	for i, num := range nums {
		left := target - num
		if idx, exists := maps[left]; exists {
			res = append(res, idx, i) // 把前面找到的 idx 放在前面 题目也是这么要求的
			return res                // 直接return
		}

		maps[num] = i
	}

	return res
}
