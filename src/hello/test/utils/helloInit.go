package utils

import (
	"fmt"
)

type pMap map[string]string
type pPerson struct {
	name   string
	height int
}

func HelloInit() {

	m := make(pMap)
	m["g"] = "google.com"
	m["b"] = "baidu.com"
	fmt.Println(m)

	p := new(pPerson)
	(*p).name = "yao"
	(*p).height = 180
	fmt.Println(*p)
}
