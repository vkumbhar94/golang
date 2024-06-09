package main

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/parquet-go/parquet-go"
	"github.com/vkumbhar94/golang/internal/util"
)

func main() {
	fmt.Println("starting the application...")
	const path = "./cmd/parquet/arb-multitype/file-multitype.parquet"

	st := reflect.StructOf([]reflect.StructField{
		{
			Name: "FirstName_String",
			Type: reflect.TypeOf(func() *string { s := ""; return &s }()),
		},
		{
			Name: "LastName_String",
			Type: reflect.TypeOf(func() *string { s := ""; return &s }()),
		},
		{
			Name: "Id_Int64",
			Type: reflect.TypeOf(func() *int64 { s := int64(0); return &s }()),
		},
		{
			Name: "Age_Int64",
			Type: reflect.TypeOf(func() *int64 { s := int64(0); return &s }()),
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
	// e1.FieldByName("FirstName").SetPointer(unsafe.Pointer(&b))
	e1.FieldByName("FirstName_String").Set(reflect.ValueOf(util.GetPtr("Bob")))
	e1.FieldByName("LastName_String").Set(reflect.ValueOf(util.GetPtr("Smith")))
	e1.FieldByName("Id_Int64").Set(reflect.ValueOf(util.GetPtr(int64(123))))
	e1.FieldByName("Age_Int64").Set(reflect.ValueOf(util.GetPtr(int64(30))))
	e2 := reflect.New(st).Elem()
	e2.FieldByName("FirstName_String").Set(reflect.ValueOf(util.GetPtr("Alice")))
	e2.FieldByName("LastName_String").Set(reflect.ValueOf(util.GetPtr("Doe")))
	e2.FieldByName("Id_Int64").Set(reflect.ValueOf(util.GetPtr(int64(456))))
	e2.FieldByName("Age_Int64").Set(reflect.ValueOf(util.GetPtr[int64](25)))

	e3 := reflect.New(st).Elem()
	e3.FieldByName("FirstName_String").Set(reflect.ValueOf(util.GetPtr("John")))
	e3.FieldByName("LastName_String").Set(reflect.ValueOf(util.GetPtr("Doe")))
	e3.FieldByName("Id_Int64").Set(reflect.ValueOf(util.GetPtr(int64(789))))
	e3.FieldByName("Age_Int64").Set(reflect.ValueOf(util.GetPtr[int64](35)))

	// fmt.Println(sl)
	// fmt.Printf("%v %T %T", reflect.Indirect(reflect.ValueOf(sl)).Interface(), reflect.ValueOf(sl).Interface(), sl)
	write, err := w.Write([]any{e1.Interface(), e2.Interface(), e3.Interface()})
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
