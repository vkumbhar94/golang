package streams

import (
	"cmp"
	"sort"
	"sync/atomic"
)

type MapFun[T, R any] func(T) R

type FilterFun[T any] func(T) bool

type Stream[T any] struct {
	data chan T
	run  func()
	ran  atomic.Bool
}

func (s *Stream[T]) Run() {
	if s.ran.CompareAndSwap(false, true) {
		go s.run()
	}
}

func New[T any](data ...T) *Stream[T] {
	ch := make(chan T)
	return &Stream[T]{
		data: ch,
		run: func() {
			func(tasks []T) {
				defer close(ch)
				for _, task := range tasks {
					ch <- task
				}
			}(data)
		},
	}
}

func Map[T, R any](s *Stream[T], mapper MapFun[T, R]) *Stream[R] {
	ch := make(chan R)
	return &Stream[R]{
		data: ch,
		run: func() {
			s.Run()
			defer close(ch)
			for t := range s.data {
				r := mapper(t)
				ch <- r
			}
		},
	}
}

func Filter[T any](s *Stream[T], filter FilterFun[T]) *Stream[T] {
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

func Limit[T any](s *Stream[T], i int) *Stream[T] {
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

type SortOrder int

const (
	ASC SortOrder = iota
	DESC
)

func Sorted[T cmp.Ordered](s *Stream[T], order SortOrder) *Stream[T] {
	ch := make(chan T)
	return &Stream[T]{
		data: ch,
		run: func() {
			s.Run()
			defer close(ch)
			result := Collect(s)
			sort.Slice(result, func(i, j int) bool {
				if order == DESC {
					return result[i] > result[j]
				}
				return result[i] < result[j]
			})
			for _, r := range result {
				ch <- r
			}
		},
	}
}

func Reduce[T any, R any](s *Stream[T], result R, f func(ans R, i T) R) R {
	s.Run()
	for t := range s.data {
		result = f(result, t)
	}
	return result
}

func ForEach[T any](stream *Stream[T], f func(i T)) {
	stream.Run()
	for t := range stream.data {
		f(t)
	}
}

// Distinct returns a new Stream with distinct elements from the input Stream.
// Stateful Intermediate Operation.
func Distinct[T comparable](s *Stream[T]) *Stream[T] {
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

func AllMatch[T any](s *Stream[T], f func(T) bool) bool {
	s.Run()
	for t := range s.data {
		if !f(t) {
			go drain(s.data)
			return false
		}
	}
	return true
}

func NotAllMatch[T any](s *Stream[T], f func(T) bool) bool {
	return !AllMatch(s, f)
}

func AnyMatch[T any](s *Stream[T], f func(T) bool) bool {
	s.Run()
	for t := range s.data {
		if f(t) {
			go drain(s.data)
			return true
		}
	}
	return false
}

func NoneMatch[T any](s *Stream[T], f func(T) bool) bool {
	return !AnyMatch(s, f)
}

func DropWhile[T any](s *Stream[T], f func(T) bool) *Stream[T] {
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

func TakeWhile[T any](s *Stream[T], f func(T) bool) *Stream[T] {
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

func Peek[T any](s *Stream[T], f func(T)) *Stream[T] {
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

func FindFirst[T any](s *Stream[T]) *T {
	s.Run()
	for t := range s.data {
		go drain(s.data)
		return &t
	}
	return nil
}

func FlatMap[T, R any](s *Stream[T], f func(T) *Stream[R]) *Stream[R] {
	ch := make(chan R)
	return &Stream[R]{
		data: ch,
		run: func() {
			s.Run()
			defer close(ch)
			for t := range s.data {
				is := f(t)
				is.Run()
				for r := range is.data {
					ch <- r
				}
			}
		},
	}
}

func Min[T cmp.Ordered](s *Stream[T]) *T {
	s.Run()
	var minVal *T
	for t := range s.data {
		if minVal == nil || t < *minVal {
			minVal = &t
		}
	}
	return minVal
}

func Max[T cmp.Ordered](s *Stream[T]) *T {
	s.Run()
	var maxVal *T
	for t := range s.data {
		if maxVal == nil || t > *maxVal {
			maxVal = &t
		}
	}
	return maxVal
}

func Skip[T any](s *Stream[T], n int) *Stream[T] {
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

func IfAllMatch[T any](s *Stream[T], f func(T) bool, action func(t T)) {
	s.Run()
	allMatch := true
	var data []T
	for t := range s.data {
		if !f(t) {
			allMatch = false
			go drain(s.data)
			break
		}
		data = append(data, t)
	}
	if allMatch {
		for _, t := range data {
			action(t)
		}
	}
}
