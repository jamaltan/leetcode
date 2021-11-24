package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// fibonacci
/**
   n <= 0 -> 0
   n = 1 -> 1
   n = 2 -> 1 + 1

   f(n) = f(n-1) + f(n-2)
**/

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	var (
		one int = 2
		two int = 1
		all int
	)

	for i := 2; i < n; i++ {
		all = one + two
		two = one
		one = all
	}

	return all
}
