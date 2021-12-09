package main

// ( )
// backtrace 回溯法
func generateParenthesis(n int) []string {
	result := make([]string, 0)

	generator(&result, n, 0, "")

	return result
}

func generator(res *[]string, left, right int, s string) {
	//终止条件
	if left == 0 && right == 0 {
		*res = append(*res, s)
		return
	}

	if left > 0 {
		generator(res, left-1, right+1, s+"(")
	}

	if right > 0 {
		generator(res, left, right-1, s+")")
	}
}
