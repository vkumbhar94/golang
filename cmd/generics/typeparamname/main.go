package main

import "fmt"

func main() {
	fmt.Println("Start")
	Print[int]()
	Print[[]int]()
	fmt.Println("End")
}

type Typ interface {
	~int | ~string | ~[]int
}

func Print[T Typ]() {
	var v any = *new(T)
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case []int:
		fmt.Println("[]int")
	}

	fmt.Printf("%T\n", v)
}
