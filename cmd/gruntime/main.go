package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("start")
	// goroutine will always run on same os thread
	// no other goroutine on the os thread
	// goroutine exits, os thread will terminate
	runtime.LockOSThread()
	runtime.UnlockOSThread()

	abc()

	fmt.Println("end")
}

func abc() {
	xyz()
}

func xyz() {
	// get immediate caller info
	pc, file, line, ok := runtime.Caller(1)
	fmt.Println(pc, file, line, ok)
	fmt.Println(runtime.Caller(1))

	// get callers' pc - need to give size of slice
	callers := make([]uintptr, 2)
	runtime.Callers(1, callers)
	fmt.Println(callers)

	runtime.CallersFrames(callers)
}
