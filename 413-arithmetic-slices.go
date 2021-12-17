package main

// a 先排序 排序里面找 len >= 3 的组合
// b dp[i] 当前节点x 能不能组成 a b x  a < b < x 依次递增
func numberOfArithmeticSlices(nums []int) int {
	//return sort(nums)

	var cnt int
	for i, _ := range nums {
		cnt += dp(i, nums)
	}

	return cnt
}

func sort(nums []int) int {
	lenght := len(nums)
	if lenght <= 2 {
		return 0
	}
	bubbleSort(nums, lenght)

	var cnt int = 1

	for left, right := 0, 2; right < lenght; right++ {
		cnt++
		if right == lenght-1 {
			left++
			right = left + 2
		}
	}

	return cnt
}

func bubbleSort(nums []int, lenght int) {
	for i := 0; i < lenght-1; i++ {
		for j := i + 1; j < lenght; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func dp(i int, nums []int) int {
	if i < 2 {
		return 0
	}

	if nums[i] > nums[i-1] && nums[i-1] > nums[i-2] {
		return 1 + dp(i-1, nums)
	}

	return 0
}
