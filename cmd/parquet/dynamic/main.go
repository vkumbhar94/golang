package main

import (
	"encoding/json"
	"fmt"
	"log"

	dynamicstruct "github.com/ompluscator/dynamic-struct"
)

import "github.com/parquet-go/parquet-go"

type RowType struct{ FirstName, LastName string }

func main() {
	fmt.Println("starting the application...")
	const path = "./cmd/parquet/dynamic/file.parquet"
	sl := dynamicstruct.NewStruct().
		AddField("FirstName", "", `json:"FirstName"`).
		AddField("LastName", "", `json:"LastName"`).
		Build().NewSliceOfStructs()
	data := []byte(`
[
	{
		"FirstName": "Aishwarya",
		"LastName": "Kumbhar"
	},
	{
		"FirstName": "Vaibhav",
		"LastName": "Kumbhar"
	}
]
`)
	err := json.Unmarshal(data, &sl)
	if err != nil {
		log.Fatal(err)
	}
	// if err := parquet.WriteFile(path, sl); err != nil {
	// 	fmt.Println("failed to write parquet file", err)
	// }

	rows, err := parquet.ReadFile[RowType](path)
	if err != nil {
		fmt.Println("failed to write parquet file", err)
	}

	for _, row := range rows {
		fmt.Println(row.FirstName)
	}

	fmt.Println("ending the application...")
}
