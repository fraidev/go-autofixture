package main

import (
	"fmt"
	"reflect"
)

type Fixture struct {
}

func NewFixture() *Fixture {
	return &Fixture{}
}

func (f *Fixture) Create(obj interface{}) {

    v := reflect.ValueOf(obj)
    indirect := reflect.Indirect(v)
    t := indirect.Type()
	for i := 0; i < t.NumField(); i++ {
		switch s := (t.Field(i).Type.Name()); s {
			case "int":
				reflect.ValueOf(obj).Elem().Field(i).SetInt(24)
			case "string":
				reflect.ValueOf(obj).Elem().Field(i).SetString("Felipe")
			case "bool":
				reflect.ValueOf(obj).Elem().Field(i).SetBool(true)
			default:
				fmt.Printf("I don't know about type %T!\n", v)
		}
		fmt.Printf("%+v\n", t.Field(i).Type)
	}

	// reflect.ValueOf(obj).Elem().Field(0).SetString("Felipe")
	// reflect.ValueOf(obj).Elem().Field(1).SetInt(24)

}
