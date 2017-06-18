package utils

import (
	"fmt"
)

type A struct {
	a int
}

type B struct {
	a, b int
}

type C struct {
	A
	B
}

type D struct {
	B
	b float32
}

func HelloAnonymous() {
	d := D{B{1, 2}, 3.0}
	fmt.Println(d, d.B.a, d.B.b, d.b)

	c := C{A{1}, B{2, 3}}
	fmt.Println(c, c.A.a, c.B.a, c.B.b)
}
