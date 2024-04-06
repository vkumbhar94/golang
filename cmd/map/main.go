package main

import "fmt"

func main() {
	fmt.Println("Map is pass by reference in Go")
	m := make(map[string]int)
	m["abc"] = 1
	m["ijk"] = 2
	m["xyz"] = 3
	fun1(m)
	fmt.Println("main", m)
}

func fun1(m map[string]int) {
	fmt.Println("Inside fun1")
	fmt.Println("fun1 before", m)
	fun2(m)
	fmt.Println("fun1 after", m)
}

func fun2(m map[string]int) {
	m["abc"] = 100

}
