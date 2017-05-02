package utils

import (
	"fmt"
	"os"
	"strconv"
)

func HelloMultiReturn() {
	var sErrVal string = "93"
	simpleReturn(sErrVal)
}

func normalReturn(sErrVal string) {

	iErrVal, err := strconv.Atoi(sErrVal)
	if nil == err {
		fmt.Printf("Normal Format: fail to convert %v[%T] to int: %v\n", sErrVal, sErrVal, err)
		// return
		// return err
		os.Exit(1)
	}
	fmt.Printf("convert %v[%T] to %v[%T]\n", sErrVal, sErrVal, iErrVal, iErrVal)
}

func simpleReturn(sErrVal string) {

	if _, ok := strconv.Atoi(sErrVal); nil != ok {
		fmt.Printf("Simple Format: fail to convert %v[%T] to int: %v\n", sErrVal, sErrVal, ok)
		os.Exit(1)
	}
	// fmt.Printf("convert %v[%T] to %v[%T]\n", sErrVal, sErrVal, iErrVal, iErrVal)
	fmt.Println("simpleReturn.success")
}
