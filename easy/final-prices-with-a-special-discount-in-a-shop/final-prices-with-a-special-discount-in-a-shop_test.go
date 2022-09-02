package final_prices_with_a_special_discount_in_a_shop

import (
	"math/rand"
	"testing"
	//"time"
)

var nums [][]int

var ns []int

func init() {
	for j := 0; j < 500; j++ {
		n := rand.Int31n(500) + 1
		num := make([]int, n)
		for i := 0; i < int(n); i++ {
			num[i] = rand.Intn(1001)
		}
		nums = append(nums, num)
	}
	ns = make([]int, 50000)
	for i := 0; i < int(50000); i++ {
		ns[i] = rand.Intn(1001)
	}
}

func BenchmarkFinalPrices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FinalPrices(ns)
	}
}

func BenchmarkFinalPrices2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FinalPrices2(ns)
	}
}

func BenchmarkFinalPrices3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FinalPrices3(ns)
	}
}
