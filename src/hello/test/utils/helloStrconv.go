package utils

import (
	"fmt"
	"strconv"
)

func HelloStrconv() {

	svIStr := "10"
	if svRet, svErr := strconv.Atoi(svIStr); svErr == nil {
		fmt.Printf("%T, %v\n", svRet, svRet)
	} else {
		fmt.Println(svErr)
	}

	svFVal := 123.4567
	// fval, type[b|E|f|g|G], prec, size[32|64]
	svFSstr := strconv.FormatFloat(svFVal, 'f', 2, 32)
	fmt.Printf("%v(%T)<>= %v(%T)\n", svFVal, svFVal, svFSstr, svFSstr)

	svStr := "☺"
	fmt.Printf("quote(%v)= %v\n	", svStr, strconv.Quote(svStr))

	test := func(s string) {
		t, err := strconv.Unquote(s)
		if err != nil {
			fmt.Printf("Unquote(%#v): %v\n", s, err)
		} else {
			fmt.Printf("Unquote(%#v) = %v\n", s, t)
		}
	}
	svStr = `\"Fran & Freddie's Diner\t\u263a\"\"`
	test("`" + svStr + "`")
	test(`"` + svStr + `"`)
}
