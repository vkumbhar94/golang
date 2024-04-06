package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("starting the application...")
	data := []string{"first", "second", "third"}
	const name = "./cmd/fileops/file.txt"
	err := os.WriteFile(name, []byte(strings.Join(data, ", ")), 0644)

	if err != nil {
		return
	}

	f, _ := os.Open(name)
	// sync writes buffer to disk
	// f.Sync()

	var read []byte
	_, err = f.Read(read)
	if err != nil {
		return
	}
	fmt.Println(string(read))
	fmt.Println("ending the application...")
}
