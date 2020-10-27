package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode, k *int) int {

	if root == nil {
		return -1
	}
	result := 0
	if root.Left != nil {
		result = dfs(root.Left, k)
	}
	if *k == 0 {
		return result
	}
	*k--

	if *k == 0 {
		return root.Val
	}

	return dfs(root.Right, k)
}

func kthSmallest(root *TreeNode, k int) int {

	return dfs(root, &k)
}
func main() {
	v := &TreeNode{5, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, nil}}

	r := kthSmallest(v, 3)
	fmt.Println(r)
}
