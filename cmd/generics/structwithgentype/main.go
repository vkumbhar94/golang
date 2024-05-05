package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(t T) {
	s.data = append(s.data, t)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.data) == 0 {
		var zero T
		return zero, errors.New("empty stack")
	}
	t := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return t, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if len(s.data) == 0 {
		var zero T
		return zero, errors.New("empty stack")
	}
	return s.data[len(s.data)-1], nil
}

func main() {
	/**
	You must instantiate generic type with a type.
	You cannot do like `var s Stack`
	**/
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}
