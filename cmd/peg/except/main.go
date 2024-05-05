package main

import "fmt"

//go:generate go install github.com/mna/pigeon@latest
//go:generate pigeon -o expr.peg.go ./expr.peg

func main() {
	parse, err := Parse("", []byte(`abc="wxy"`), Debug(false))
	if err != nil {
		panic(err)
	}
	fmt.Println(parse)
	arr := []string{
		`abc="abc"`,
		`ijk="ijk"`,
		`lmn="lmn"`,
		`xyz="xyz"`,
	}

	for _, v := range arr {
		parse, err := Parse("", []byte(v), Debug(false))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(parse)
	}
}
