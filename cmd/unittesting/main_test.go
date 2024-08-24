package main

import (
	"testing"
)

func TestDivide(t *testing.T) {
	got := divide(6, 3)
	if got != 2 {
		t.Errorf("divide(6, 3) = %d; want 2", got)
	}
}

func FuzzDivide(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b int) {
		if b == 0 {
			t.Skip()
		}
		divide(a, b)
	})
}

func TestOnlyCharacters(t *testing.T) {
	got := onlyChars("abc123")
	if got != false {
		t.Errorf("onlyChars(%q) = %v; want false", "abc123", got)
	}
}

func FuzzOnlyCharacters(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		r1 := onlyChars(s)
		r2 := onlyCharsRegex(s)
		if r1 != r2 {
			t.Errorf("onlyChars(%q) = %v; onlyCharsRegex(%q) = %v; want both to be equal", s, r1, s, r2)
		}
	})
}
