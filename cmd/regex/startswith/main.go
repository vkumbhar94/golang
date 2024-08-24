package main

import (
	"fmt"
	"regexp"
)

func main() {
	p := regexp.MustCompile(`^abc.*`)
	fmt.Println(p.LiteralPrefix())
	fmt.Println(p.String())
	fmt.Println(p.MatchString(`kjdkfabclkjsd`))
}
