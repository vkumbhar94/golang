package main

import (
	"fmt"
)

import "github.com/parquet-go/parquet-go"

type RowType struct{ FirstName, LastName string }

func main() {
	fmt.Println("starting the application...")
	if err := parquet.WriteFile("file.parquet", []RowType{
		{FirstName: "Bob"},
		{FirstName: "Alice"},
	}); err != nil {
		fmt.Println("failed to write parquet file", err)
	}

	rows, err := parquet.ReadFile[RowType]("file.parquet")
	if err != nil {
		fmt.Println("failed to write parquet file", err)
	}

	for _, row := range rows {
		fmt.Println(row.FirstName)
	}

	fmt.Println("ending the application...")
}
