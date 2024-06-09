package main

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/parquet-go/parquet-go"
)

func main() {
	fmt.Println("starting the application...")
	const path = "./cmd/parquet/arb/file.parquet"
	st := reflect.StructOf([]reflect.StructField{
		{
			Name: "FirstName",
			Type: reflect.TypeOf(""),
		},
		{
			Name: "LastName",
			Type: reflect.TypeOf(""),
		},
	})
	fmt.Println("st: ", st)

	schema := parquet.SchemaOf(reflect.New(st).Elem().Interface())
	fmt.Println("schema: ", schema)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)

	}
	w := parquet.NewGenericWriter[any](f, schema)
	// sl := reflect.MakeSlice(reflect.SliceOf(st), 0, 10)
	e1 := reflect.New(st).Elem()
	e1.FieldByName("FirstName").SetString("Bob")
	e1.FieldByName("LastName").SetString("Smith")
	e2 := reflect.New(st).Elem()
	e2.FieldByName("FirstName").SetString("Alice")
	e2.FieldByName("LastName").SetString("Doe")

	// fmt.Println(sl)
	// fmt.Printf("%v %T %T", reflect.Indirect(reflect.ValueOf(sl)).Interface(), reflect.ValueOf(sl).Interface(), sl)
	write, err := w.Write([]any{e1.Interface(), e2.Interface()})
	if err != nil {
		panic(err)
	}
	fmt.Println(write)
	err = w.Flush()
	if err != nil {
		panic(err)
	}
	err = w.Close()
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second)
	ofile, err := os.Open(path)
	if err != nil {
		panic(err)

	}
	reader := parquet.NewReader(ofile, schema)
	re1 := reflect.New(st)
	err = reader.Read(re1.Interface())
	if err != nil {
		panic(err)
	}

	fmt.Println(re1.Elem().Interface())

	sl := make([]any, 0)
	for i := 0; i < 1; i++ {
		re := reflect.New(st)
		err = reader.Read(re.Interface())
		if err != nil {
			panic(err)
		}
		sl = append(sl, re.Elem().Interface())
	}
	fmt.Println(sl)

	fmt.Println("ending the application...")
}
func typeOf[T any]() reflect.Type {
	var v T
	return reflect.TypeOf(v)
}
