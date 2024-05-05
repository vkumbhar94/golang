package main

import "fmt"

//1go:generate go install github.com/mna/pigeon@latest
//go:generate pigeon -o expr.peg.go ./expr.peg
func main() {
	a, e := Parse("", []byte(``), Debug(false), Entrypoint("EmptyQuery"))
	if e != nil {
		fmt.Println(e)
	}
	b, e2 := Parse("", []byte(`a=`), Debug(false), Entrypoint("AssignQuery"))
	if e2 != nil {
		fmt.Println(e2)
	}
	fmt.Println(a, b)

	m := make(map[string]string)
	m["ns1"] = "ns1 id"
	m["ns2"] = "ns2 id"
	c, e3 := Parse("", []byte(`ns1`), Debug(false), Entrypoint("Input"), GlobalStore("namespaces", m))
	if e3 != nil {
		fmt.Println(e3)
	}
	fmt.Println(c)
	d, e4 := Parse("", []byte(`ns3`), Debug(false), Entrypoint("Input"), GlobalStore("namespaces", m))
	if e4 != nil {
		fmt.Println(e4)
	}
	fmt.Println(d)
}
