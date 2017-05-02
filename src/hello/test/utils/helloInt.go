package utils

import "fmt"

func HelloInt() {
	// init
	var iInt16 int16 = 30
	var iInt32 int32
	iInt32 = int32(iInt16)

	// convert
	// iInt32 = iInt16 + iInt16
	iInt32 = iInt32 + 100

	// output
	fmt.Printf("int16= %d\n", iInt16)
	fmt.Printf("int32= %d\n", iInt32)
}
