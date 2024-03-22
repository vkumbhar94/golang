package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abc/t1/ns1"
	arr := strings.SplitN(s, "/", 2)
	fmt.Println(arr)

	arr2 := []string{"a", "b", "c", "d", "e"}
	for i, v := range arr2 {
		arr2[i] = "\"" + v + "\""
	}
	fmt.Println(strings.Join(arr2, ", "))
}
