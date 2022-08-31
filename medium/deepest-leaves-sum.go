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

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	t := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	if val > root.Val {
		t.Left = root
		return t
	}
	if root.Left != nil && val > root.Left.Val {
		t.Left = root.Left
		root.Left = t
		return t
	}
	if root.Right != nil && val > root.Right.Val {
		t.Right = root.Right
		root.Right = t
		return t
	}

	if insertIntoMaxTree(root.Left, val) != nil {
		return root
	} else if insertIntoMaxTree(root.Right, val) != nil {
		return root
	}
	return nil
}
