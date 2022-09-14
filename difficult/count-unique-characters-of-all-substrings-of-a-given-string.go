package main

func uniqueLetterString(s string) int {
	count := 0

	for i := 0; i < len(s); i++ {
		m := make(map[string]struct{})
		for j := i; j < len(s); j++ {
			m[string(s[j])] = struct{}{}
			count += len(m)
		}
	}
	return count
}

func uniqueLetterString1(s string) (ans int) {
	idx := map[rune][]int{}
	for i, c := range s {
		idx[c] = append(idx[c], i)
	}
	for _, arr := range idx {
		arr = append(append([]int{-1}, arr...), len(s))
		for i := 1; i < len(arr)-1; i++ {
			ans += (arr[i] - arr[i-1]) * (arr[i+1] - arr[i])
		}
	}
	return
}
