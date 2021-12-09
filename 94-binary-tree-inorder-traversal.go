package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var (
	res = make([]int, 0)
)

// 给定的函数申明参数中没有 返回值[]int,所以不能直接使用recursion
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return res
	}

	inorderTraversal(root.Left)
	res = append(res, root.Val)
	inorderTraversal(root.Right)

	return res
}

// 使用迭代
func inorderTraversal(root *TreeNode) []int {
}
