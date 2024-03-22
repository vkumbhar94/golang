package main

import "fmt"

func main() {
	fmt.Println("Starting the application...")

	arr := [...]string{"a", "b", "c", "d", "e"}
	fmt.Printf("%T\n", arr)
	slice := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("%T\n", slice)
	fmt.Println("Ending the application...")
}
