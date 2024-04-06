package main

import (
	"embed"
	"fmt"
)

//go:embed dist/*
var dist embed.FS

func main() {
	file, err := dist.ReadFile("dist/sample.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	dir, err := dist.ReadDir("dist")
	fmt.Println(dir, err)
	for _, d := range dir {
		fmt.Println(d.Name(), d.IsDir())
	}
}
