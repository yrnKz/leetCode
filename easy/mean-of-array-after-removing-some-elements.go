package main

import (
	"sort"
)

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	k := int(float32(len(arr)) * 0.05)
	println(k)
	sum := 0
	for i := k; i < len(arr)-k; i++ {
		sum += arr[i]
	}

	return float64((sum * 1.0) / ((len(arr) - 2*k) * 1.0))
}

//func main() {
//	a := []int{6, 0, 7, 0, 7, 5, 7, 8, 3, 4, 0, 7, 8, 1, 6, 8, 1, 1, 2, 4, 8, 1, 9, 5, 4, 3, 8, 5, 10, 8, 6, 6, 1, 0, 6, 10, 8, 2, 3, 4}
//	println(len(a))
//	println(trimMean(a))
//}
