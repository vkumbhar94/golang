package main

import "testing"

func BenchmarkFooBuilder(b *testing.B) {
	var f Foo
	for i := 0; i < b.N; i++ {
		(*FooBuilder)(&f).
			SetA(1).
			SetB("hello").
			SetC(3.14)
	}
}
