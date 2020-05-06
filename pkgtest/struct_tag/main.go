package main

import (
	"fmt"
	"reflect"
)

// 自定义tag

const tagName = "validate"

type User struct {
	Id    int    `validate:"-"json:"id"`
	Name  string `validate:"presence,min=2,max=32"`
	Email string `validate:"email,required"`
}

func main() {
	user := User{
		Id:    1,
		Name:  "John Doe",
		Email: "john@example",
	}
	t := reflect.TypeOf(user)
	// Get the type and kind of our user variable
	fmt.Println("Type: ", t.Name())
	fmt.Println("Kind: ", t.Kind())
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)
		//Get the field tag value
		fmt.Println(reflect.TypeOf(field.Tag))
		tag := field.Tag.Get(tagName)
		fmt.Printf("%d. %v(%v), tag:'%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}
}
