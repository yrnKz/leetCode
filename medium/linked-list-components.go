package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = struct{}{}
	}
	p := head
	maxc := 0
	flag := false
	for p != nil {
		if _, ok := m[p.Val]; !ok {
			if flag {
				maxc++
			}
			flag = false
		} else {
			flag = true
		}
		p = p.Next
	}

	if flag {
		maxc++
	}
	return maxc
}
