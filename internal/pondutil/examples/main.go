package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/vkumbhar94/golang/internal/pondutil"
)

func main() {
	p := pondutil.New(2, 3)
	result, err := pondutil.RunTasksWithSupplierFunc(p, context.Background(), func(ch chan<- pondutil.Task[int64]) {
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			for i := 1; i <= 10; i++ {
				v := i
				ch <- func() (int64, error) {
					return int64(v), nil
				}
			}
			wg.Done()
		}()
		go func() {
			for i := 11; i <= 20; i++ {
				v := i
				ch <- func() (int64, error) {
					return int64(v), nil
				}
			}
			wg.Done()
		}()
		wg.Wait()
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
