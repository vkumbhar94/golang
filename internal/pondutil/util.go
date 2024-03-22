package pondutil

func MapToTasks[I any, R any](in <-chan I, mapper func(I) Task[R]) <-chan Task[R] {
	out := make(chan Task[R])
	go func() {
		for res := range in {
			out <- mapper(res)
		}
		close(out)
	}()
	return out
}
