package main

import "fmt"

type Foo struct {
	a int
	b string
	c float64
}

type FooBuilder Foo

func (fb *FooBuilder) SetA(a int) *FooBuilder {
	fb.a = a
	return fb
}

func (fb *FooBuilder) SetB(b string) *FooBuilder {
	fb.b = b
	return fb
}

func (fb *FooBuilder) SetC(c float64) *FooBuilder {
	fb.c = c
	return fb
}

func main() {
	fmt.Println("starting zero alloc builder...")
	var f Foo

	// golang type system is the beauty here.
	// It just casts the object to another type whereas the referred underneath memory is the same
	// here pointer to object of Foo is type casted to pointer of FooBuilder
	(*FooBuilder)(&f).
		SetA(1).
		SetB("hello").
		SetC(3.14)

	fmt.Println(f)
	fmt.Println("ending zero alloc builder...")
}
