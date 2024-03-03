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
}

type AllVars3 struct {
	Addresses []string
	Name      string
	Float64   float64
	Int32     int32
	Int       int
	Bool      bool
}

func main() {
	fmt.Println("🏃 exploring go memory model...")

	obj1 := AllVars1{}
	obj2 := AllVars2{}
	obj3 := AllVars3{}

	util.PPSize(obj1)
	util.PPSize(obj2)
	util.PPSize(obj3)

	fmt.Println("✅ Done")
}
