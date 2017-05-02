package utils

import "fmt"

func HelloClosure() {
	helloFuncRet()
}

func helloFuncCaller() {

	func() { fmt.Println("start...") }()

	funcCaller := func(iVal int) { fmt.Printf("closure.funcCaller(%d)\n", iVal) }
	for i := 0; i < 4; i++ {
		funcCaller(i)
	}
}

func helloFuncRet() {

	iVal := 3
	fAdd2 := add2()
	fmt.Printf("add2()(%d)= %d\n", iVal, fAdd2(iVal))

	fAddB := addB(iVal)
	fmt.Printf("addB(%d)(%d)= %d", iVal, iVal, fAddB(iVal))
}

func add2() func(iVal int) int {
	return func(iVal int) int {
		return iVal + 2
	}
}

func addB(a int) func(b int) int {
	return func(b int) int {
		return b + a
	}
}
