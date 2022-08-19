package main

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

var max = 0
var lmax = 0

func deepestLeavesSum(root *TreeNode) int {
	findDeepestLeaves(root, 0, 0)

	return max

}

func findDeepestLeaves(node *TreeNode, ans, leave int) {
	ans += node.Val
	leave++
	if leave > lmax {
		max = ans
		lmax = leave
	}
	if node.Right != nil {
		findDeepestLeaves(node.Right, ans, leave)
	}
	if node.Left != nil {
		findDeepestLeaves(node.Left, ans, leave)
	}
	leave--
	ans -= node.Val
}
