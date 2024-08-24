package main

import "fmt"

func main() {
	// #Tip Number 1
	arr := [2]int{1, 2}
	fmt.Println(arr)

	// what if you add one more element to the array
	// you need to make changes at 2 places, one to increase array size and add element
	// arr2 := [3]int{1,2,3}

	// Instead, you can define array like following
	// compiler will calculate length of array for you
	arr2 := [...]int{1, 2, 3}

	fmt.Println(arr2)

	// #Tip Number 2

	// Say, you need to define some large number for some calculation, or constant for limits
	// you can use underscore to make it more readable
	// 1_00_000 is same as 100000
	// you can easily count zeros in 1_00_000
	maxRecords := 1_00_000 // 1 lac
	fmt.Println(maxRecords)

	// #Tip Number 3
	name := "Bob"
	fmt.Printf("My name is %s. Yes, you heard that right: %s\n", name, name)

	// Now see, how we can pass same argument multiple times.
	// In `%[1]s` 1 is order of argument in list
	fmt.Printf("My name is %[1]s. Yes, you heard that right: %[1]s\n", name)

	// you can decide order of arguments in list
	fmt.Printf("%[2]s's age is %[1]d\n", 25, "Bob")
	// You will get output like `Bob's age is 25`

	fmt.Println("That's all for today. Happy coding!")
}
