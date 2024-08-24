package main

import (
	"fmt"
	"os"
)
import "github.com/kelseyhightower/envconfig"

type Abc struct {
	A map[string]int      `envconfig:"ABC"`
	B map[string]struct{} `envconfig:"ABCD"`
}

func main() {
	var a Abc
	os.Setenv("ABC", "abc:123,xyz:345")
	os.Setenv("ABCD", "abc:,xyz:")
	envconfig.Process("", &a)
	fmt.Println(a)
	if _, ok := a.B["abc"]; ok {
		fmt.Println("abc exists")
	}

	if _, ok := a.B["ijk"]; ok {

	} else {
		fmt.Println("ijk does not exist")
	}

}
