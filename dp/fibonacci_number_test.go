package dp

import "testing"

// go test -v -bench=BenchmarkFib
func BenchmarkFibRecursive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib_rec(50)
	}
}

func BenchmarkFibMemo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib_memo(50)
	}
}
func BenchmarkFibDP(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(50)
	}
}

func BenchmarkFibDPOpt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib_opt(50)
	}
}
