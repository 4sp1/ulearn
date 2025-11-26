package lc274

import "testing"

func TestHIndex(t *testing.T) {
	for _, test := range []struct {
		citations []int
		h         int
	}{
		{[]int{3, 0, 6, 1, 5}, 3},
		{[]int{1, 3, 1}, 1},
		{[]int{100}, 1},
		{[]int{0, 0, 2}, 1},
	} {
		h := hIndex(test.citations)
		if test.h != h {
			t.Logf("h-index of %v is %d not %d", test.citations, test.h, h)
			t.Fail()
		}
	}
}
