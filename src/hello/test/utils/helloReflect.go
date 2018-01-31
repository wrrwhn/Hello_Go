package utils

import (
	"fmt"
	"reflect"
)

func HelloReflect() {
	type MyInt int
	var x MyInt = 7
	v := reflect.ValueOf(x)
	fmt.Println(v.Type(), v.Type().Kind(), v.CanSet())

	y := 3.4
	v = reflect.ValueOf(y)
	fmt.Println(v.Type(), v.Type().Kind(), v.CanSet())

	v = reflect.ValueOf(&y)
	fmt.Println(v.Type(), v.Type().Kind(), v.CanSet(), v.Elem().CanSet())

	type T struct {
		A int
		B string
	}
	t := T{1, "A"}
	e := reflect.ValueOf(&t).Elem()
	tType := e.Type()
	fmt.Println(tType)
	for i := 0; i < tType.NumField(); i++ {
		f := e.Field(i)
		fmt.Println(tType.Field(i).Name, f.Type(), f.Interface(), f.CanSet())
		if 1 == i {
			f.SetString("B")
		}
	}
	fmt.Println(t)
}
