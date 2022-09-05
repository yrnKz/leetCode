package main

var max2 int

func longestUnivaluePath(root *TreeNode) int {
	longestUnivaluePath2(root, 0)
	return max2
}
func longestUnivaluePath2(root *TreeNode, long int) {
	if root == nil {
		return
	}
	if max2 < long {
		max2 = long
	}
	if root.Left != nil && root.Val == root.Left.Val {
		long++
		longestUnivaluePath2(root.Left, long)
	} else {
		longestUnivaluePath2(root.Left, 0)
	}

	if root.Right != nil && root.Val == root.Right.Val {
		long++
		longestUnivaluePath2(root.Right, long)
	} else {
		longestUnivaluePath2(root.Right, 0)
	}
}
