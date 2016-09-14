package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

type Foo2 struct {
	Uno string `tag_name:"tag f"`
	Dos  string `tag_name:"tag a"`
}

func reflection(f interface{}) {
	fmt.Println("type:", reflect.TypeOf(f))
	val := reflect.ValueOf(f).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
	}
}

func main() {
	f := &Foo{
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
	}
	f2 := &Foo2{
		Uno: "11",
		Dos:  "22",
	}

	reflection(f)
	reflection(f2)
}
