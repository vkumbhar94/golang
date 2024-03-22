package main

import (
	"fmt"

	"github.com/vkumbhar94/go-streams"
)

func main() {
	stream := streams.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	filtered := streams.Filter(stream, func(i int) bool {
		return i%2 == 0
	})
	mapped := streams.Map(filtered, func(i int) int {
		return i * 2
	})
	limited := streams.Limit(mapped, 3)

	collected := streams.Collect(limited)
	fmt.Println(collected)

	abc()
	abc2()
	abc3()
	abc4()
	abc5()
	abc6()
}

func abc6() {
	input := []int{1, 2, 3, 5, 6, 4, 7, 9, 10, 8}
	collected := streams.ToNumberStream(streams.New(input...).Filter(func(i int) bool {
		return i%2 == 0
	}).Map(func(i int) int {
		return i * 2
	})).Average()

	fmt.Println(collected)
}

func abc5() {
	input := []int{1, 2, 3, 5, 6, 4, 7, 9, 10, 8}
	collected := streams.ToNumberStream(streams.New(input...).Filter(func(i int) bool {
		return i%2 == 0
	}).Map(func(i int) int {
		return i * 2
	})).Sum()

	fmt.Println(collected)
}

func abc4() {
	input := []int{1, 2, 3, 5, 6, 4, 7, 9, 10, 8}
	collected := streams.New(input...).Filter(func(i int) bool {
		return i%2 == 0
	}).Map(func(i int) int {
		return i * 2
	}).Limit(3).Reverse().Collect()

	fmt.Println(collected)
}

func abc3() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	collected := streams.ToOrderedStream(
		streams.New(input...).Filter(func(i int) bool {
			return i%2 == 0
		}).Map(func(i int) int {
			return i * 2
		}).Limit(3),
	).Sorted(streams.ASC).Reverse().Collect()

	fmt.Println(collected)

}

func abc2() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	collected := streams.New(input...).Filter(func(i int) bool {
		return i%2 == 0
	}).Map(func(i int) int {
		return i * 2
	}).Limit(3).Collect()

	fmt.Println(collected)
}

func abc() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	collected := streams.Collect(
		streams.Limit(
			streams.Map(
				streams.Filter(
					streams.New(input...),
					func(i int) bool {
						return i%2 == 0
					},
				),
				func(i int) int {
					return i * 2
				},
			),
			3,
		),
	)

	fmt.Println(collected)
}
