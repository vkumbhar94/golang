package main

import "testing"

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

func BenchmarkAdd2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add2(1, 2)
	}
}
