package utils

import (
	"fmt"
	"reflect"
)

type TypeTag struct {
	name   string "type_tag.name"
	height int    "type_tag.height"
}

func refTag(typeTag TypeTag) {
	t := reflect.TypeOf(typeTag)
	fmt.Printf("Align=%v, fieldAlign=%v, numField=%v\nField.tags\n", t.Align(), t.FieldAlign(), t.NumField())
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("\t%v\n", t.Field(i).Tag)
	}
}

func HelloTypeTag() {
	typeTag := TypeTag{"180", 180}
	refTag(typeTag)
}
