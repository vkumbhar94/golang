package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	Explain()
}

func resource() func() int {
	i := 0
	return func() int {
		fmt.Println("returning new")
		i++
		return i
	}
}

var p = resource()

func getNew() any {
	return p()
}

func worker(id int, pool *sync.Pool, result *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		a := pool.Get()
		val, _ := result.LoadOrStore(fmt.Sprintf("r%d", a), new(int64))
		atomic.AddInt64(val.(*int64), 1)
		//fmt.Printf("%2d : %2d\n", id, a)
		//time.Sleep(time.Millisecond)
		pool.Put(a)
	}
}
func Explain() {
	fmt.Println("Starting execution...")
	pool := &sync.Pool{New: getNew}

	result := &sync.Map{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i, pool, result, wg)
	}
	wg.Wait()
	result.Range(func(k, v any) bool {
		fmt.Printf("%s : %2d\n", k, *(v.(*int64)))
		return true
	})
	fmt.Println("Completed execution.")
}
