package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	arr := []int{1, 3, 2, 6, 4, 5}

	for _, v := range slices.Backward(arr) {
		println(v)
	}
	for v := range slices.Values(arr) {
		println(v)
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range maps.All(m) {
		println(k, v)
	}

	for k := range maps.Keys(m) {
		println(k)
	}

	for v := range maps.Values(m) {
		println(v)
	}
	println("End")

	arr2 := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr2 = append(arr2, i)
	}

	fmt.Println(len(arr2))
}

// func reverse[T any](arr []T) iter.Seq[T] {
//	return func(yield func(T)) {
//		for i := len(arr) - 1; i >= 0; i-- {
//			yield(arr[i])
//		}
//	}
// }
