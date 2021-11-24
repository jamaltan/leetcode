package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// 额外开辟一份空间res， 不符合题目要求
func moveZeroes(nums []int) {
	right := len(nums)

	res := make([]int, right)
	var j int
	for i := 0; i < right; i++ {
		if nums[i] != 0 {
			res[j] = nums[i]
			j++
		}
	}

	nums = res
}

// 记录最后有一个非0元素的下标j 遇到下一个非0的元素就房子 j+1的位置上
// 后面再来一层for循环 重置后面元素为0
// 两个单层 for 循环
func moveZeroes(nums []int) {
	var j int

	for _, num := range nums {
		if num != 0 {
			nums[j] = num
			j++
		}
	}

	for k := j; k < len(nums); k++ {
		nums[k] = 0
	}
}

// 可不可以使用单层迭代呢
// snowball 滚雪球理论
func moveZeroes(nums []int) {
	var snowBarSize int

	// 遇到非0怎么办， 交换
	// 和谁交换 和雪球的第一个元素0交换 k =
	for i, num := range nums {
		if num == 0 {
			snowBarSize++
		} else {
			nums[i], nums[i-snowBarSize] = nums[i-snowBarSize], nums[i] //记录了雪球的大小， 就可以知道雪球的第一个元素的位置
		}
	}
}
