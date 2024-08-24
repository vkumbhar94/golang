package streams

import (
	"sync/atomic"

	"golang.org/x/exp/constraints"
)

type NumberStream[T constraints.Integer | constraints.Float] struct {
	Stream[T]
}

func ToNumberStream[T constraints.Integer | constraints.Float](s *Stream[T]) *NumberStream[T] {
	return &NumberStream[T]{
		Stream: Stream[T]{
			data: s.data,
			run:  s.run,
			ran:  atomic.Bool{},
		},
	}
}
func (s *NumberStream[T]) Sum() (result T) {
	s.Run()
	for t := range s.data {
		result += t
	}
	return result
}
func (s *NumberStream[T]) Average() (result float64) {
	s.Run()
	var count int
	for t := range s.data {
		result += float64(t)
		count++
	}
	if count == 0 {
		return 0
	}
	return result / float64(count)
}
func (s *NumberStream[T]) Max() (result *T) {
	s.Run()
	for t := range s.data {
		if result == nil || t > *result {
			result = &t
		}
	}
	return result
}

func (s *NumberStream[T]) Min() (result *T) {
	s.Run()
	for t := range s.data {
		if result == nil || t < *result {
			result = &t
		}
	}
	return result
}

func (s *NumberStream[T]) Count() (result int64) {
	s.Run()
	for range s.data {
		result++
	}
	return result
}
