package main

import "fmt"

//go:generate go install github.com/mna/pigeon@latest
//go:generate pigeon -o expr.peg.go ./expr.peg

func main() {
	parse, err := Parse("", []byte("(ab,xyz,)"), Debug(false))
	if err != nil {
		panic(err)
	}
	fmt.Println(parse)
}
