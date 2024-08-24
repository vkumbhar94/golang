package streams

import (
	"cmp"
	"fmt"
	"sort"
)

type MapEntry[K comparable, V any] struct {
	K K
	V V
}

func (t MapEntry[K, V]) Key() K {
	return t.K
}

func (t MapEntry[K, V]) Value() V {
	return t.V
}

func (t MapEntry[K, V]) String() string {
	return fmt.Sprintf("(%v, %v)", t.K, t.V)
}

func MNew[K comparable, V any](data map[K]V) *Stream[MapEntry[K, V]] {
	ch := make(chan MapEntry[K, V])
	return &Stream[MapEntry[K, V]]{
		data: ch,
		run: func() {
			func(tasks map[K]V) {
				defer close(ch)
				for k, v := range tasks {
					ch <- MapEntry[K, V]{k, v}
				}
			}(data)
		},
	}
}

func MKeys[K comparable, V any](data map[K]V) *Stream[K] {
	ch := make(chan K)
	return &Stream[K]{
		data: ch,
		run: func() {
			func(tasks map[K]V) {
				defer close(ch)
				for k := range tasks {
					ch <- k
				}
			}(data)
		},
	}
}

func MValues[K comparable, V any](data map[K]V) *Stream[V] {
	ch := make(chan V)
	return &Stream[V]{
		data: ch,
		run: func() {
			func(tasks map[K]V) {
				defer close(ch)
				for _, v := range tasks {
					ch <- v
				}
			}(data)
		},
	}
}

func MCollect[K comparable, V any](stream *Stream[MapEntry[K, V]]) map[K]V {
	stream.Run()
	result := make(map[K]V)
	for t := range stream.data {
		result[t.K] = t.V
	}
	return result
}

func MSorted[K cmp.Ordered, V any](s *Stream[MapEntry[K, V]]) *Stream[MapEntry[K, V]] {
	ch := make(chan MapEntry[K, V])
	return &Stream[MapEntry[K, V]]{
		data: ch,
		run: func() {
			s.Run()
			defer close(ch)
			result := Collect(s)
			sort.Slice(result, func(i, j int) bool {
				return result[i].K < result[j].K
			})
			for _, r := range result {
				ch <- r
			}
		},
	}
}
