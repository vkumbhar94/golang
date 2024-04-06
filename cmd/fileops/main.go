package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("starting the application...")
	data := []string{"first", "second", "third"}
	os.WriteFile()

	f, _ := os.Open("abc")
	f.Sync()
	f.Read(data)
	fmt.Println("ending the application...")
}
