package check_if_number_is_a_sum_of_powers_of_three

func checkPowersOfThree(n int) bool {
	if n == 0 {
		return false
	}
	nums := [20]int{}
	for i := 0; i < len(nums); i++ {
		nums[i] = pow(i)
	}
	return dfs(n, 0, nums)
}

func dfs(n, i int, nums [20]int) bool {
	if n == 0 {
		return true
	}
	if n < 0 || i >= len(nums) {
		return false
	}
	if dfs(n-nums[i], i+1, nums) {
		return true
	}
	return dfs(n, i+1, nums)
}

func pow(n int) int {
	if n == 0 {
		return 1
	}
	ans := 1
	for i := 0; i < n; i++ {
		ans *= 3
	}
	return ans
}
