package utils

import "fmt"

func HelloFuncFunc() {

	iVal := 10
	fmt.Printf("ixi(%d)= %d", iVal, callFuncArgs(iVal, ixi))
}

func callFuncArgs(i int, f func(int, int) int) int {
	return f(i, i)
}

func ixi(i, j int) int {
	return i * j
}
