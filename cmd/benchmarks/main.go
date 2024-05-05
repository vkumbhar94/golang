package main

import "golang.org/x/exp/constraints"

func main() {

}

func Add(a, b int) int {
	return a + b
}

func Add2[T constraints.Integer | constraints.Float](a, b T) T {
	return a + b
}
