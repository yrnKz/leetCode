package main

func main() {

}

func shuffle(nums []int, n int) []int {
	ans := make([]int, 2*n)
	k := 0
	for i := 0; i < n; i++ {
		ans[k] = nums[i]
		ans[k+1] = nums[n+i]
		k += 2
	}
	return ans
}

func shuffle2(nums []int, n int) []int {
	k := 0
	for i := 0; i < n; i++ {
		nums[k] |= (nums[i] & 1023) << 10
		nums[k+1] |= (nums[n+i] & 1023) << 10
		k += 2
	}
	for i := 0; i < 2*n; i++ {
		nums[i] = nums[i] >> 10
	}
	return nums
}
