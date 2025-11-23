package first

import (
	"math/rand"
	"strings"
	"testing"
)

type Simple struct{}

func (op Simple) first(i int) int { return i + 1 }
func (op Simple) rest(i int) int  { return i + 1 }

func BenchmarkIsFirst(b *testing.B) {
	t := make([]int, 1e9)
	var op Simple
	for b.Loop() {
		isFirst := true
		for _, v := range t {
			if isFirst {
				op.first(v)
				isFirst = false
				continue
			}
			op.rest(v)
		}
	}
}

func BenchmarkIndexEqualZero(b *testing.B) {
	var op Simple
	t := make([]int, 1e9)
	for i := range t {
		t[i] = rand.Intn(1e3)
	}
	for b.Loop() {
		for i, v := range t {
			if i == 0 {
				op.first(v)
				continue
			}
			op.rest(v)
		}
	}
}

type StringBuilder struct{}

// func (op StringBuilder) first(b strings.Builder, s string) {
// 	b.WriteString("> ")
// 	b.WriteString(s)
// 	b.WriteRune('\n')
// }
// func (op StringBuilder) rest(b strings.Builder, s string) {
// 	b.WriteString("  ")
// 	b.WriteString(s)
// 	b.WriteRune('\n')
// }

func (op StringBuilder) rand(alphabet []rune) string {
	length := rand.Intn(50)
	t := make([]rune, length)
	for i := range t {
		t[i] = '!' + rune(rand.Intn(len(alphabet)))
	}
	return string(t)
}
func (op StringBuilder) alphabet() []rune {
	alphabet := make([]rune, 94)
	i := 0
	for r := '!'; r <= '~'; r++ {
		alphabet[i] = r
		i++
	}
	return alphabet
}

func TestStringBuilderRand(t *testing.T) {
	var op StringBuilder
	for _, r := range op.rand(op.alphabet()) {
		if r > '~' || r < '!' {
			t.Fatal()
		}
	}
}

func (op StringBuilder) seed() []string {
	t := make([]string, 1e6)
	alphabet := op.alphabet()
	for i := range t {
		t[i] = op.rand(alphabet)
	}
	return t
}

func BenchmarkOpStringBuilderFirstContinue(b *testing.B) {
	var op StringBuilder
	choices := op.seed()
	for b.Loop() {
		var s strings.Builder
		var first = true
		for _, v := range choices {
			if first {
				s.WriteString("> ")
				s.WriteString(v)
				s.WriteRune('\n')
				first = false
				continue
			}
			s.WriteString("  ")
			s.WriteString(v)
			s.WriteRune('\n')
		}
	}
}

func BenchmarkOpStringBuilderIndexContinue(b *testing.B) {
	var op StringBuilder
	choices := op.seed()
	for b.Loop() {
		var s strings.Builder
		for i, v := range choices {
			if i == 0 {
				s.WriteString("> ")
				s.WriteString(v)
				s.WriteRune('\n')
				continue
			}
			s.WriteString("  ")
			s.WriteString(v)
			s.WriteRune('\n')
		}
	}
}

func BenchmarkOpStringBuilderFirstAlt(b *testing.B) {
	var op StringBuilder
	choices := op.seed()
	for b.Loop() {
		var s strings.Builder
		first := true
		for _, v := range choices {
			if first {
				s.WriteRune('>')
				first = false
			} else {
				s.WriteRune(' ')
			}
			s.WriteRune(' ')
			s.WriteString(v)
			s.WriteRune('\n')
		}
	}
}
func BenchmarkOpStringBuilderIndexAlt(b *testing.B) {
	var op StringBuilder
	choices := op.seed()
	for b.Loop() {
		var s strings.Builder
		for i, v := range choices {
			if i == 0 {
				s.WriteRune('>')
			} else {
				s.WriteRune(' ')
			}
			s.WriteRune(' ')
			s.WriteString(v)
			s.WriteRune('\n')
		}
	}
}
