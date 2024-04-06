package main

import (
	"fmt"
	"sync"
)

func abc() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}
func main() {
	mycounter := abc()
	fmt.Println(mycounter())
	fmt.Println(mycounter())
	fmt.Println(mycounter())
	fmt.Println(mycounter())
	mycounter = abc()
	fmt.Println(mycounter())
	fmt.Println(mycounter())

	ab(nil)

	forLoopObj()

}

func forLoopObj() {
	fmt.Println("forLoopObj start")
	arr := []*int{
		getPtr(1),
		getPtr(2),
		getPtr(3),
		getPtr(4),
		getPtr(5),
		getPtr(6),
		getPtr(7),
		getPtr(8),
		getPtr(9),
		getPtr(10),
		getPtr(11),
		getPtr(12),
		getPtr(13),
		getPtr(14),
		getPtr(15),
		getPtr(16),
		getPtr(17),
		getPtr(18),
		getPtr(19),
		getPtr(20),
		getPtr(21),
		getPtr(22),
		getPtr(23),
		getPtr(24),
		getPtr(25),
	}

	ch := make(chan func() int)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for f := range ch {
			f()
		}
	}()
	go func() {
		defer wg.Done()
		for f := range ch {
			f()
		}
	}()
	func() {
		for _, v := range arr {
			cv := v
			ch <- func() int {
				fmt.Println(*cv)
				return *cv
			}
		}
		close(ch)
	}()
	wg.Wait()
	fmt.Println("forLoopObj end")
}
func getPtr[T any](t T) *T {
	return &t
}
func ab(t *int64) {

}
