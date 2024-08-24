package streams

import (
	"sync/atomic"
)

func (s *Stream[T]) Filter(filter FilterFun[T]) *Stream[T] {
	ch := make(chan T)
	return &Stream[T]{
		data: ch,
		run: func() {
			s.Run()
			defer close(ch)
			for t := range s.data {
				if filter(t) {
					ch <- t
				}
			}
		},
	}
}

func (s *Stream[T]) Limit(i int) *Stream[T] {
	ch := make(chan T)
	return &Stream[T]{
		data: ch,
		run: func() {
			s.Run()
			func() {
				defer close(ch)
				for t := range s.data {
					if i > 0 {
						ch <- t
						i--
					} else {
						break
					}
				}
			}()

			for range s.data {
			}

		},
	}
}

func (s *Stream[T]) ForEach(f func(i T)) {
	s.Run()
	for t := range s.data {
		f(t)
	}
}

func (s *Stream[T]) AllMatch(f func(T) bool) bool {
	s.Run()
	for t := range s.data {
		if !f(t) {
			go drain(s.data)
			return false
		}
	}
	return true
}

func (s *Stream[T]) NotAllMatch(f func(T) bool) bool {
	return !s.AllMatch(f)
}

func (s *Stream[T]) AnyMatch(f func(T) bool) bool {
	s.Run()
	for t := range s.data {
		if f(t) {
			go drain(s.data)
			return true
		}
	}
	return false
}

func (s *Stream[T]) NoneMatch(f func(T) bool) bool {
	return !AnyMatch(s, f)
}

func (s *Stream[T]) DropWhile(f func(T) bool) *Stream[T] {
	ch := make(chan T)
	return &Stream[T]{
		data: ch,
		run: func() {
			s.Run()
			defer close(ch)
			dropping := true
			for t := range s.data {
				if dropping && f(t) {
					continue
				}
				dropping = false
				ch <- t
			}
		},
	}
}

func (s *Stream[T]) TakeWhile(f func(T) bool) *Stream[T] {
	ch := make(chan T)
	return &Stream[T]{
		data: ch,
		run: func() {
			s.Run()
			func() {
				defer close(ch)
				for t := range s.data {
					if f(t) {
						ch <- t
					} else {
						break
					}
				}
			}()
			go drain(s.data)
		},
	}
}

func (s *Stream[T]) Peek(f func(T)) *Stream[T] {
	ch := make(chan T)
	return &Stream[T]{
		data: ch,
		run: func() {
			s.Run()
			defer close(ch)
			for t := range s.data {
				f(t)
				ch <- t
			}
		},
	}
}

type OrStream[T any] struct {
	Stream[T]
}

func (s *OrStream[T]) Or(or T) T {
	s.Run()
	for t := range s.data {
		go drain(s.data)
		return t
	}
	return or
}

func (s *Stream[T]) FindFirst() *T {
	s.Run()
	for t := range s.data {
		go drain(s.data)
		return &t
	}
	return nil
}

func (s *Stream[T]) FindFirstOr() *OrStream[T] {
	s.Run()
	for t := range s.data {
		go drain(s.data)
		ch := make(chan T)
		return &OrStream[T]{
			Stream[T]{
				data: ch,
				run: func() {
					defer close(ch)
					ch <- t
				},
				ran: atomic.Bool{},
			},
		}
	}
	ch := make(chan T)
	return &OrStream[T]{
		Stream[T]{
			data: ch,
			run: func() {
				defer close(ch)
			},
			ran: atomic.Bool{},
		},
	}
}

func (s *Stream[T]) Skip(n int) *Stream[T] {
	ch := make(chan T)
	return &Stream[T]{
		data: ch,
		run: func() {
			s.Run()
			defer close(ch)
			for t := range s.data {
				if n > 0 {
					n--
					continue
				}
				ch <- t
			}
		},
	}
}

type ElseStream[T any] struct {
	Stream[T]
}

func (s *ElseStream[T]) Else(action func(t T)) {
	s.Run()
	for t := range s.data {
		action(t)
	}
}

func (s *Stream[T]) IfAllMatch(f func(T) bool, action func(t T)) *ElseStream[T] {
	s.Run()
	allMatch := true
	var data []T
	for t := range s.data {
		if allMatch && !f(t) {
			allMatch = false
		}
		data = append(data, t)
	}
	if allMatch {
		for _, t := range data {
			action(t)
		}
	}
	ch := make(chan T)
	return &ElseStream[T]{
		Stream: Stream[T]{
			data: ch,
			run: func() {
				defer close(ch)
				for _, t := range data {
					ch <- t
				}
			},
			ran: atomic.Bool{},
		},
	}
}

func (s *Stream[T]) Collect() []T {
	return Reduce(s, []T{}, func(ans []T, i T) []T {
		return append(ans, i)
	})
}

type UnaryMapFun[T any] func(T) T

func (s *Stream[T]) Map(mapper UnaryMapFun[T]) *Stream[T] {
	ch := make(chan T)
	return &Stream[T]{
		data: ch,
		run: func() {
			s.Run()
			defer close(ch)
			for t := range s.data {
				ch <- mapper(t)
			}
		},
	}
}

func (s *Stream[T]) Reduce(result T, f func(ans T, i T) T) T {
	s.Run()
	for t := range s.data {
		result = f(result, t)
	}
	return result
}

func (s *Stream[T]) Count() (cnt int64) {
	s.Run()
	for range s.data {
		cnt++
	}
	return
}

func (s *Stream[T]) Reverse() *Stream[T] {
	ch := make(chan T)
	return &Stream[T]{
		data: ch,
		run: func() {
			s.Run()
			defer close(ch)
			var data []T
			for t := range s.data {
				data = append(data, t)
			}
			for i := len(data) - 1; i >= 0; i-- {
				ch <- data[i]
			}
		},
	}
}
