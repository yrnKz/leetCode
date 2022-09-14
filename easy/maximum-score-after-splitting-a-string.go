package main

func maxScore(s string) int {
	max := 0
	for i := 1; i <= len(s)-1; i++ {
		var temp int
		for j := 0; j < len(s); j++ {
			if j < i && s[j] == '0' {
				temp++
			} else if j >= i && s[j] == '1' {
				temp++
			}
		}
		if temp > max {
			max = temp
		}
	}
	return max
}
