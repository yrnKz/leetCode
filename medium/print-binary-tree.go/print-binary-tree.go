package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func calDepth(node *TreeNode) int {
	h := 0
	if node.Left != nil {
		h = calDepth(node.Left) + 1
	}
	if node.Right != nil {
		h = max(h, calDepth(node.Right)+1)
	}
	return h
}

func printTree(root *TreeNode) [][]string {
	height := calDepth(root)
	m := height + 1
	n := 1<<m - 1
	ans := make([][]string, m)
	for i := range ans {
		ans[i] = make([]string, n)
	}
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, r, c int) {
		ans[r][c] = strconv.Itoa(node.Val)
		if node.Left != nil {
			dfs(node.Left, r+1, c-1<<(height-r-1))
		}
		if node.Right != nil {
			dfs(node.Right, r+1, c+1<<(height-r-1))
		}
	}
	dfs(root, 0, (n-1)/2)
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
