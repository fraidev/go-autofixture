package main

import (
	"fmt"
	"reflect"
	"unicode"
)

type Fixture struct {
}

func NewFixture() *Fixture {
	return &Fixture{}
}

func (f *Fixture) Create(obj interface{}) {
	t := reflect.Indirect(reflect.ValueOf(obj))
	valueOf := reflect.ValueOf(obj)
	typeOf := reflect.TypeOf(obj)

	for i := 0; i < t.NumField(); i++ {

		valueOfField := t.Field(i)
		kind := valueOfField.Kind()

		typeOfField := typeOf.Elem().Field(i)
		typeOfFieldName := typeOfField.Name
		fmt.Println(typeOfFieldName)

		//Ignore private fields
		if !unicode.IsUpper(rune(typeOfFieldName[0])) {
			return
		}

		switch kind {
		case reflect.Bool:
			valueOf.Elem().Field(i).SetBool(true)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uintptr:
			valueOf.Elem().Field(i).SetInt(24)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			valueOf.Elem().Field(i).SetUint(24)
		case reflect.Float32, reflect.Float64:
			valueOf.Elem().Field(i).SetFloat(24)
		case reflect.Complex64, reflect.Complex128:
			valueOf.Elem().Field(i).SetComplex(complex(6, 2))
		case reflect.Array:
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Slice, reflect.UnsafePointer:
		case reflect.String:
			valueOf.Elem().Field(i).SetString("Felipe")
		case reflect.Ptr:
			h := valueOf.Elem().Field(i).Interface()
			fmt.Printf("%+v\n", valueOfField)
			f.Create(h)
			fmt.Printf("%+v\n", valueOfField)
		case reflect.Struct:
			fmt.Println("Ignoring struct values")
		case reflect.Invalid:
			fmt.Println("Invalid Type")
		default:
			fmt.Println("Type not expected")
		}
	}
}
