package main

import "fmt"

func abc() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}
func main() {
	mycounter := abc()
	fmt.Println(mycounter())
	fmt.Println(mycounter())
	fmt.Println(mycounter())
	fmt.Println(mycounter())
	mycounter = abc()
	fmt.Println(mycounter())
	fmt.Println(mycounter())

	ab(nil)

}

func ab(t *int64) {

}
