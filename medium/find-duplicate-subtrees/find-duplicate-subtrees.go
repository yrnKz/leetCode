package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var ans []*TreeNode
	f(root.Right, root.Left, &ans)
	return ans
}

func f(r1, r2 *TreeNode, ans *[]*TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if (r1 == nil && r2 != nil) || (r2 == nil && r1 != nil) {
		return false
	}
	//f(r1, r2.Right, ans)
	//f(r1, r2.Left, ans)
	f(r2, r1.Left, ans)
	f(r2, r1.Right, ans)
	s1 := f(r1.Right, r2.Right, ans)
	s2 := f(r1.Left, r2.Left, ans)
	if r1.Val == r2.Val && s1 && s2 {
		*ans = append(*ans, r1)
		return true
	}
	return false
}
func main() {
	r := &TreeNode{Val: 1}
	l := &TreeNode{Val: 1}
	root := &TreeNode{
		Val:   2,
		Left:  l,
		Right: r,
	}

	ans := findDuplicateSubtrees(root)
	for _, i2 := range ans {
		print(i2.Val)
	}
}
