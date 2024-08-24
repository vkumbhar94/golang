package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("vim-go")
}

func divide(a, b int) int {
	return a / b
}

func onlyChars(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
		} else {
			return false
		}
	}
	return true
}

func onlyCharsRegex(s string) bool {
	m, err := regexp.MatchString("^[a-z]*$", s)
	if err != nil {
		return false
	}
	return m
}
