package utils

import "fmt"

func HelloPoint() {
	changePointToVal()
}

func readPoint() {

	var iPointVal = 1
	fmt.Printf("%v %p\n", iPointVal, &iPointVal)
	var ptrPointVal *int = &iPointVal
	fmt.Printf("%p \n", ptrPointVal)
}

func changePointToVal() {
	sPointVal := "init-val"
	var ptrStrVal = &sPointVal
	fmt.Printf("[%T](%p)= %v \t%p\n", sPointVal, &sPointVal, sPointVal, ptrStrVal)

	*ptrStrVal = "change-val"
	fmt.Printf("[%T](%p)= %v \t%p\n", sPointVal, &sPointVal, sPointVal, ptrStrVal)
}
