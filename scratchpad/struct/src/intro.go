package main

import (
	"fmt"
	"reflect"
)

// Person is a mock struct
type Person struct {
	Name    string
	Address *Address
}

// Address is a mock struct
type Address struct {
	Street string
	Region string
}

func main() {
	person := Person{}
	// Recursive call printField to print each field

	str := printField(reflect.TypeOf(person), "")

	fmt.Println(str)
}

func printField(structType reflect.Type, str string) string {
	// get struct name and kind
	structName := structType.Name()
	structKind := structType.Kind()
	str = fmt.Sprintf("%s \n Struct type is %s, kind is %s, and it has fields:", str, structName, structKind)

	// get number of field of this struct
	fieldNum := structType.NumField()
	for i := 0; i < fieldNum; i++ {
		// get each field's name, type and kind
		fieldName := structType.Field(i).Name
		fieldType := structType.Field(i).Type
		fieldKind := fieldType.Kind()

		// if kind is a pointer, get use Elem() to get the actual kind
		if fieldKind == reflect.Ptr {
			fieldType = fieldType.Elem()
		}

		// if a field is a struct go into next loop to print that struct infos
		if fieldType.Kind() == reflect.Struct {
			str = fmt.Sprintf("%s %s is a type of %s,", str, fieldName, fieldType.Name())
			str = printField(fieldType, str)
		} else {
			str = fmt.Sprintf("%s %s is a type of %s,", str, fieldName, fieldKind)
		}
	}

	return str
}
