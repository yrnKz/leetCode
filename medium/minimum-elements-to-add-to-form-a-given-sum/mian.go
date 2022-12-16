package minimum_elements_to_add_to_form_a_given_sum

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	sum = sum - goal
	if sum < 0 {
		sum = -sum
	}
	if limit < 0 {
		limit = -limit
	}
	r := sum / limit
	if sum%limit != 0 {
		r++
	}
	return r
}
