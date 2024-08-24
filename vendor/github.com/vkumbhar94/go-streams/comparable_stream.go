package streams

import "sync/atomic"

type ComparableStream[T comparable] struct {
	Stream[T]
}

func ToComparableStream[T comparable](s *Stream[T]) *ComparableStream[T] {
	return &ComparableStream[T]{
		Stream: Stream[T]{
			data: s.data,
			run:  s.run,
			ran:  atomic.Bool{},
		},
	}
}
func (s *ComparableStream[T]) CollectToSet() map[T]struct{} {
	s.Run()
	ans := make(map[T]struct{})
	for t := range s.data {
		ans[t] = struct{}{}

	}
	return ans
}

func (s *ComparableStream[T]) Distinct() *Stream[T] {
	ch := make(chan T)
	return &Stream[T]{
		data: ch,
		run: func() {
			s.Run()
			defer close(ch)
			seen := make(map[T]struct{})
			for t := range s.data {
				if _, ok := seen[t]; !ok {
					seen[t] = struct{}{}
					ch <- t
				}
			}
		},
	}
}
func (s *ComparableStream[T]) DistinctAndThen() *ComparableStream[T] {
	ch := make(chan T)
	return &ComparableStream[T]{
		Stream[T]{
			data: ch,
			run: func() {
				s.Run()
				defer close(ch)
				seen := make(map[T]struct{})
				for t := range s.data {
					if _, ok := seen[t]; !ok {
						seen[t] = struct{}{}
						ch <- t
					}
				}
			},
		},
	}
}
