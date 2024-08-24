package streams

import "golang.org/x/exp/constraints"

func drain[T any](ch <-chan T) {
	for range ch {
	}
}

func Collect[T any](s *Stream[T]) []T {
	return Reduce(s, []T{}, func(ans []T, i T) []T {
		return append(ans, i)
	})
}

func Count[T any](s *Stream[T]) int64 {
	return Reduce(s, int64(0), func(ans int64, i T) int64 {
		return ans + 1
	})
}

func Sum[T constraints.Integer | constraints.Float](s *Stream[T]) T {
	return Reduce(s, 0, func(ans T, i T) T {
		return ans + i
	})
}

func CollectToSet[T comparable](stream *Stream[T]) map[T]struct{} {
	return Reduce(stream, map[T]struct{}{}, func(ans map[T]struct{}, i T) map[T]struct{} {
		ans[i] = struct{}{}
		return ans
	})
}
