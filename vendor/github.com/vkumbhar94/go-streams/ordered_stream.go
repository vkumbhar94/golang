package streams

import (
	"cmp"
	"sort"
	"sync/atomic"
)

type OrderedStream[T cmp.Ordered] struct {
	Stream[T]
}

func ToOrderedStream[T cmp.Ordered](s *Stream[T]) *OrderedStream[T] {
	return &OrderedStream[T]{
		Stream: Stream[T]{
			data: s.data,
			run:  s.run,
			ran:  atomic.Bool{},
		},
	}
}

func (s *OrderedStream[T]) Sorted(order SortOrder) *Stream[T] {
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
			if order == DESC {
				sort.Slice(data, func(i, j int) bool {
					return data[i] > data[j]
				})
			} else {
				sort.Slice(data, func(i, j int) bool {
					return data[i] < data[j]
				})
			}
			for _, t := range data {
				ch <- t
			}
		},
	}
}

type MOrderedStream[K cmp.Ordered, V any] struct {
	Stream[MapEntry[K, V]]
}

func ToMOrderedStream[K cmp.Ordered, V any](s *Stream[MapEntry[K, V]]) *MOrderedStream[K, V] {
	return &MOrderedStream[K, V]{
		Stream: Stream[MapEntry[K, V]]{
			data: s.data,
			run:  s.run,
			ran:  atomic.Bool{},
		},
	}
}

func (s *MOrderedStream[K, V]) Sorted(order SortOrder) *Stream[MapEntry[K, V]] {
	ch := make(chan MapEntry[K, V])
	return &Stream[MapEntry[K, V]]{
		data: ch,
		run: func() {
			s.Run()
			defer close(ch)
			var data []MapEntry[K, V]
			for t := range s.data {
				data = append(data, t)
			}
			if order == DESC {
				sort.Slice(data, func(i, j int) bool {
					return data[i].K > data[j].K
				})
			} else {
				sort.Slice(data, func(i, j int) bool {
					return data[i].K < data[j].K
				})
			}
			for _, t := range data {
				ch <- t
			}
		},
	}
}
