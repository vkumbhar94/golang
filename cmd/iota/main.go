package main

import (
	"fmt"
)

type EnumA int

const (
	_ EnumA = iota
	A1
	A2
	A3
	_
	A5
)

type EnumB int

// when we define first enum as unknown, it helps to handle error better.
// Typically, parse methods return error.
// but when the handler/consumer wants to know whether it is given or set as default while soft handling parse error

const (
	Unknown EnumB = iota
	B1
	B2
)

// Always implement String methods on enums, never ever export enums underneath integer values outside/over the wire.
// underneath int value must never be stored to db nor communicate with other service
func (b EnumB) String() string {
	switch b {
	case B1:
		return "b1"
	case B2:
		return "b2"
	case Unknown:
		return "unknown"
	default:
		return "unknown"
	}
}

type EnumC int

// When you don't want to define unknown right away, as least start sequencing from 1.
// So that if required, it keeps hole to insert unknown.
const (
	C1 EnumC = iota + 1
	C2
)

func (c EnumC) String() string {
	switch c {
	case C1:
		return "c1"
	case C2:
		return "c2"
	default:
		// invalid
		return "unknown"
	}
}

func (c EnumC) String2() string {
	switch c {
	case C1:
		return "c1"
	case C2:
		return "c2"
	}
	// invalid
	return "unknown"
}

func main() {
	fmt.Println("start")

	fmt.Println("A1: ", A1, A2, A3)
	fmt.Println("B1: ", B1, B2)
	var b EnumB
	fmt.Println("init enum: B", b)

	fmt.Println("end")
}
