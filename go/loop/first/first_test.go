package first

import "testing"

func opfirst(i int) int { return i + 1 }
func oprest(i int) int  { return i + 1 }

func BenchmarkIsFirst(b *testing.B) {
	for b.Loop() {
		isFirst := true
		t := make([]int, 1e9)
		for _, v := range t {
			if isFirst {
				opfirst(v)
				isFirst = false
			}
			oprest(v)
		}
	}
}

func BenchmarkIndexEqualZero(b *testing.B) {
	for b.Loop() {
		t := make([]int, 1e9)
		for i, v := range t {
			if i == 0 {
				opfirst(v)
			}
			oprest(v)
		}
	}
}
