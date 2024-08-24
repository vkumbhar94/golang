# Java equivalent stream APIs in go

- This library provides a set of APIs to perform operations on a collection of elements similar to Java Stream APIs.

- Similar to Java, intermediate operations are lazy and terminal operations are eager i.e. intermediate operations are not performed until a terminal operation is called.

- when intermediate operations are performed, a new stream is returned, and the original stream is not modified.

- Intermediate operations that can short circuit will halt once the condition is met, and the rest of the elements will not be processed. To avoid go routine leaks, the stream is cleared after obtaining the result.

- The library make use of go routines to perform operations in parallel.
- The library is designed to be used with a collection of elements, and not with a channel.
- The library is not thread safe.
- The library is not designed to be used with infinite streams.
- Stream can only be used once, and it is not reusable.


## Usage

```go
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
}

```