package main

import (
	"fmt"
	"reflect"
)

type Lol struct {
	Lol  *string
	Name string
}

func main() {
	l := Lol{}
	val := reflect.ValueOf(l)
	fmt.Println(reflect.ValueOf(l).NumField())
	for i := 0; i < 2; i++ {
		fmt.Println(val.Field(i))
	}
	//fmt.Println(reflect.TypeOf(l.Lol))
	//fmt.Println(reflect.ValueOf(l.Lol).Kind().String())
	//fmt.Println(reflect.ValueOf(l.Lol).IsNil())
}
