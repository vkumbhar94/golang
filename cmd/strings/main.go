package main

import (
	"fmt"
	"strconv"
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

	fmt.Println(strconv.Unquote(`"abc"`))

	// using first argument multiple times in string format
	const rdsCertURLFormat = "https://truststore.pki.rds.amazonaws.com/%[1]s/%[1]s-bundle.pem"
	fmt.Println(fmt.Sprintf(rdsCertURLFormat, "us-gov-west-1"))
}
