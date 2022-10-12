package advantage_shuffle

import "sort"

func advantageCount(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	ans := make([]int, len(nums2))
	for i := 0; i < len(nums2); i++ {
		j := 0
		for ; j < len(nums1); j++ {
			if nums1[j] > nums2[i] {
				ans[i] = nums1[j]
				nums1[j] = -1
				break
			}
		}
		if j >= len(nums1) {
			for k := 0; k < len(nums1); k++ {
				if nums1[k] != -1 {
					ans[i] = nums1[k]
					nums1[k] = -1
					break
				}
			}
		}
	}
	return ans
}
