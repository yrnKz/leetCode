package main

func canBeEqual(target []int, arr []int) bool {
	m := make([]int, 1001)

	for i := 0; i < len(target); i++ {
		m[target[i]]++
	}
	for _, v := range arr {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}

	return true
}
