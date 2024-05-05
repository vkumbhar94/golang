package main

import "fmt"

type Play interface {
	Play()
}

type tmp struct {
	v Play
}

type Top struct {
	tmp tmp
}

type Player struct {
}

func (p *Player) Play() {
	println("play")
}

func main() {
	t := &Top{
		tmp: tmp{
			v: nil,
		},
	}
	if t.tmp.v != nil {
		t.tmp.v.Play()
	}

	main2()
}

type myerr string

func (err myerr) Error() string {
	return "an error ocurred"
}

func do() *myerr {
	return nil // returns nil pointer
}

func wrap() error {
	return do() // the information about nil pointer is dissolved
}

func main2() {
	err := wrap()
	fmt.Println(err == nil) // prints `false` because underneath is nil pointer not empty interface

	// Explaination
	var e *myerr = nil
	var e2 error = e
	// because e2 is not nil, it is a nil pointer of *myerr
	fmt.Printf("%T\n", e2)
	fmt.Println(e2)
	fmt.Println(e2 == nil) // prints `false` because underneath is nil pointer not empty interface
}
