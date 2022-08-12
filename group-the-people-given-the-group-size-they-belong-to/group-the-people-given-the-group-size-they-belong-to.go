package main

import "fmt"

func main() {
	a := []int{3, 3, 3, 3, 3, 1, 3}
	fmt.Println(groupThePeople(a))
}

func groupThePeople(groupSizes []int) [][]int {
	var dp [][]int
	var dlen []int
	for i := 0; i < len(groupSizes); i++ {
		flag := false
		for j := 0; j < len(dp); j++ {
			if cap(dp[j]) == groupSizes[i] && dlen[j] < cap(dp[j]) {
				dp[j][dlen[j]] = i
				dlen[j]++
				flag = true
				break
			}
		}
		if !flag {
			temp := make([]int, groupSizes[i])
			temp[0] = i
			dlen = append(dlen, 1)
			dp = append(dp, temp)
		}
	}
	return dp
}
