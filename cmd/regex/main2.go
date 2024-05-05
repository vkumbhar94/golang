package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	ex := regexp.MustCompile(`(?:ab(\d+))`)
	// res := ex.FindAllStringSubmatch("ab123ab78", -1)
	// fmt.Println(res)

	val := []byte("ab123ab78")
	indAll := ex.FindAllSubmatchIndex(val, -1)
	fmt.Println(indAll)
	i := 0
	replaceString := "*"
	var replaced []byte
	for _, ind := range indAll {

		if len(ind) < 3 {
			panic("less than 3")
		}
		iter := toIterator(ind[2:])

		for {
			next := iter()
			if next == nil {
				replaced = append(replaced, val[i:]...)
				break
			}
			v := *next
			replaced = append(replaced, val[i:v]...)
			i = *iter()

			replaced = append(replaced, []byte(strings.Repeat(replaceString, i-v))...)
		}
	}
	fmt.Println("replaced", string(replaced))
}

func toIterator[T any](vs []T) func() *T {
	i := 0
	return func() *T {
		if i >= len(vs) {
			return nil
		}
		v := vs[i]
		i++
		return &v
	}
}
