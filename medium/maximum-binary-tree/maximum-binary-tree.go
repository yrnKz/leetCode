package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	max := 0
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	t := &TreeNode{
		Val:   max,
		Left:  nil,
		Right: nil,
	}
	if index != 0 {
		t.Left = constructMaximumBinaryTree(nums[:index])
	}
	if index != len(nums)-1 {
		t.Right = constructMaximumBinaryTree(nums[index+1:])
	}
	return t
}
