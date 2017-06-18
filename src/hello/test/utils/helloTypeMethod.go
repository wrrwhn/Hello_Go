package utils

import (
	"fmt"
)

type Class struct {
	name string
}

func (this *Class) Name() string {
	return this.name
}
func (this *Class) SetName(newName string) {
	this.name = newName
}

func HelloTypeMethod() {

	var this Class = Class{"yao"}
	fmt.Println(this, this.name)

	(&this).SetName("lu")
	fmt.Println(this, (&this).Name())
}
