package main

// EmptyTypeSet tilde allows no types
// no type can have both int32 as well as float64
type EmptyTypeSet interface {
	int32
	float64
}

func Add[T Number](a, b T) T {
	return a + b
}
func main() {

}
