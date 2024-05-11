package main

import "fmt"

type Thread interface {
	Run()
}

type Thread1 struct{}
type Thread2 struct{}

func (t Thread1) Run() {
	fmt.Println("Thread1")
}

func (t Thread2) Run() {
	fmt.Println("Thread2")
}

func main() {
	fmt.Println("Start")
	Run(Thread1{})
	Run(Thread2{})
	fmt.Println("End")
}

type NamedThread interface {
	Thread
	Id() string
}

func Run[T Thread](t T) {
	t.Run()
}

func RunNamedThread[T NamedThread](t T) {
	fmt.Println(t.Id())
	t.Run()
}
