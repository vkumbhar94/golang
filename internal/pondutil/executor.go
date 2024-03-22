package pondutil

import (
	"context"
)

func RunTasksAndPipe[R any](p *Pool, ctx context.Context, tasks ...Task[R]) (<-chan R, error) {
	return RunTasksWithSupplierFuncAndPipe(p, ctx, func(out chan<- Task[R]) {
		for _, task := range tasks {
			out <- task
		}
	})
}
func RunTasksWithSupplierFuncAndPipe[R any](p *Pool, ctx context.Context, fn func(chan<- Task[R])) (<-chan R, error) {
	inch := make(chan Task[R])
	out, err := RunTasksWithSupplierChanAndPipe(p, ctx, inch)
	if err != nil {
		// returning on failure, make sure to propagate to error to outside function's variable
		return nil, err
	}
	go func() {
		fn(inch)
		close(inch)
	}()
	return out, err

}

func RunTasksWithSupplierChanAndPipe[R any](p *Pool, ctx context.Context, inch <-chan Task[R]) (<-chan R, error) {
	ch := make(chan R)
	go func(out chan R) {
		tg, _ := p.GroupContext(ctx)
		for task := range inch {
			cpy := task
			tg.Submit(func() error {
				r, err := cpy()
				if err != nil {
					// TODO: silently ignore error
					// support logger function to log error
					// return err
					return nil
				}
				out <- r
				return nil
			})
		}
		defer close(out)
		// need to wait to close channel after all tasks completion
		err := tg.Wait() // wait for all tasks to complete
		if err != nil {
			return
		}

	}(ch)
	return ch, nil
}
