package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// 最基础的推导 位置 position p
// p[i][j] = p[i-1][j] + p[i][j-1]
// p[0][j] = p[i][0] = 1

// 时间复杂度 time O(m*n) cost space  O(m*n)
func uniquePaths(m int, n int) int {
	res := make([][]int, m)
	for idx := range res {
		res[idx] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		res[i][0] = 1
	}

	for j := 0; j < n; j++ {
		res[0][j] = 1
	}

	for i := 1; i < m; i++ { //i,j >= 1 开始
		for j := 1; j < n; j++ {
			res[i][j] = res[i][j-1] + res[i-1][j]
		}
	}

	return res[m-1][n-1]
}

// 优化空间复杂度
func uniquePaths(m int, n int) int {

}

// 事件复杂度 O(m*n) cost space  O(n)
func uniquePaths(m int, n int) int {
	res := make(map[int]int, 0)
	res[0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j] += res[j-1]
		}
	}

	return res[n-1]
}
