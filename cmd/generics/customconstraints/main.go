package main

import "fmt"

// OnlyInt32Or64 tilde allows only int32 or int64
type OnlyInt32Or64 interface {
	int32 | int64
}

// OnlyInt32Or64OrDescendants tilde allows all types whos underlying type is int32 or int64
type OnlyInt32Or64OrDescendants interface {
	~int32 | ~int64
}

func Add[T OnlyInt32Or64](a, b T) T {
	return a + b
}

func Add2[T OnlyInt32Or64OrDescendants](a, b T) T {
	return a + b
}

type MyInt32 int32

func main() {
	// fmt.Println(Add(10,20)) // fails to work as int is not in allowed types
	fmt.Println(Add(int32(10), int32(20))) // works
	// fmt.Println(Add(MyInt32(10), MyInt32(20))) // doesn't work because constraint is not allowing descendants
	fmt.Println(Add2(MyInt32(10), MyInt32(20))) // works because constraint allows descendants
}
