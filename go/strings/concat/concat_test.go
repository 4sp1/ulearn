package concat

import (
	"strings"
	"testing"
)

func TestConcat(t *testing.T) {
	s := "hell"
	s += "o"
	if s != "hello" {
		t.Fatalf(`"hell"+"o" != "hello"`)
	}
}

func TestConcatJoin(t *testing.T) {
	s := "hell"
	s = strings.Join([]string{s, "o"}, "")
	if s != "hello" {
		t.Fatalf(`strings join "hell","o" != "hello"`)
	}
}

func BenchmarkConcatSum(b *testing.B) {
	for b.Loop() {
		s := "hell"
		s += "o"
		if s != "hello" {
			b.Fatal()
		}
	}
}

func BenchmarkConcatJoin(b *testing.B) {
	for b.Loop() {
		s := "hell"
		s = strings.Join([]string{s, "o"}, "")
		if s != "hello" {
			b.Fatal()
		}
	}
}
