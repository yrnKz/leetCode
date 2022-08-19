package main

func main() {
	nums := []int{2, 2, 1, 1, 5, 3, 3, 5}
	maxEqualFreq(nums)
}

func maxEqualFreq(nums []int) int {
	count := map[int]int{}
	freq := map[int]int{}
	var max, ans int
	for i, num := range nums {
		if count[num] > 0 {
			freq[count[num]]--
		}
		count[num]++
		freq[count[num]]++
		if count[num] > max {
			max = count[num]
		}
		if max == 1 ||
			(freq[max]*max+freq[max-1]*(max-1) == i+1 && freq[max] == 1) ||
			(freq[max]*max+1 == i+1 && freq[1] == 1) {
			ans = i
		}

	}
	return ans
}
