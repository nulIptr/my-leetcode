package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	return isValidBSTInner(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTInner(root *TreeNode, min, max int64) bool {
	if root == nil {
		return true
	}

	if int64(root.Val) >= max || int64(root.Val) <= min {
		return false
	}

	return isValidBSTInner(root.Left, min, int64(root.Val)) && isValidBSTInner(root.Right, int64(root.Val), max)

}
func main() {

	fmt.Println()
}
