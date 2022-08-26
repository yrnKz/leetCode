package main

import "math"

func main() {
	a := 0xfffffff
	println(a)
}

func findClosestElements(arr []int, k int, x int) []int {
	a := 0xfffffff
	index := 0
	for i := 0; i < len(arr); i++ {
		if int(math.Abs(float64(arr[i]-x))) < a {
			index = i
			a = int(math.Abs(float64(arr[i] - x)))
		}
	}
	res := make([]int, 0)
	res = append(res, arr[index])
	l := index - 1
	r := index + 1
	for len(res) < k {
		if l < 0 {
			res = append(res, arr[r])
			r++
			continue
		}
		if r >= len(arr) {
			res = append([]int{arr[l]}, res...)
			l--
			continue
		}
		if int(math.Abs(float64(arr[l]-x))) <= int(math.Abs(float64(arr[r]-x))) {
			res = append([]int{arr[l]}, res...)
			l--
		} else {
			res = append(res, arr[r])
			r++
		}
	}
	return res
}
