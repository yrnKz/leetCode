package main

func validateStackSequences(pushed []int, popped []int) bool {
	var r, index int
	stake := make([]int, len(pushed))

	for i := 0; i < len(pushed); i++ {
		stake[r] = pushed[i]
		for r >= 0 && index < len(popped) && stake[r] == popped[index] {
			index++
			r--
		}
		r++
	}
	r--
	for r >= 0 {
		if stake[r] != popped[index] {
			return false
		}
		r--
		index++
	}
	return true
}
