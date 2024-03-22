package main

import (
	"fmt"

	"golang/internal/util"
)

type AllVars1 struct {
	Name      string
	Addresses []string
	Int       int32
	Float64   float64
	Bool      bool
}

type AllVars2 struct {
	Addresses []string
	Bool      bool
	Int       int32
	Name      string
	Float64   float64
	B         bool
}

type AllVars3 struct {
	v2        AllVars2
	Bool      bool
	Addresses []string
	Name      string
	Float64   float64
	Int       int
	Int32     int32
}

func main() {
	fmt.Println("🏃 exploring go memory model...")

	obj1 := AllVars1{}
	obj2 := AllVars2{}
	obj3 := AllVars3{}

	util.PPSize(obj1)
	util.PPSize(obj2)
	util.PPSize(obj3)

	var arr []string
	sort(arr)
	fmt.Println("✅ Done")
}

func sort(arr []string) {

}
