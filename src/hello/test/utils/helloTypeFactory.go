package utils

import (
	"errors"
	"fmt"
)

// private
type person struct {
	name   string
	height int
	weight int
}

// InitPerson
func InitPerson(name string, height int, weight int) (*person, error) {
	if height < 0 {
		return nil, errors.New("height less than 0")
	} else if weight < 0 {
		return nil, errors.New("weight less than 0")
	}

	return &person{name, height, weight}, nil
}

// HelloTypeFactory Type工厂类
func HelloTypeFactory() {

	var p, _ = InitPerson("yao", 180, 75)
	fmt.Println(*p)
}
