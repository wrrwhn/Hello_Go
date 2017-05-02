package utils

import "fmt"

func HelloFunc() {
	helloDefer()
}

// return
func helloFuncReturn() {

	var iStartVal int = 10
	printFunc(getXx2AndYx3(iStartVal))
	x3, y4 := getXx3AndYx4(iStartVal)
	printFunc(iStartVal, x3, y4)
}

func getXx2AndYx3(iStartVal int) (int, int, int) {
	return iStartVal, iStartVal * 2, iStartVal * 3
}

func getXx3AndYx4(iStartVal int) (x3, y4 int) {
	x3 = iStartVal * 3
	y4 = iStartVal * 4

	return
}

func printFunc(x, y, z int) {
	fmt.Printf("x= %d, y= %d, z= %d\n", x, y, z)
}

// arg = point
func helloFuncPtr() {

	var iVal int = 10
	var iResult int = 0

	squarePtr(iVal, &iResult)
	// square(iVal, iResult)
	fmt.Printf("%d^2= %d\n", iVal, iResult)
}

func squarePtr(iVal int, iResult *int) {

	*iResult = iVal * iVal
}

func square(iVal int, iResult int) {
	iResult = iVal * iVal
}

// args...
func helloFuncArgsChange() {

	iMin := min(10, 11, 9)
	fmt.Println(iMin)
}

func min(args ...int) int {

	if 0 == len(args) {
		return 0
	}

	min := args[0]
	for _, val := range args {
		if min > val {
			min = val
		}
	}

	return min
}

// args...
func helloFuncEmptyArgsChange() {
	typeCheck(1, "str", 1.1, 'c')
}

func typeCheck(values ...interface{}) {

	for _, val := range values {
		switch val.(type) {
		case int:
			fmt.Printf("[%v](%T) -> int\n", val, val)
		case float64:
			fmt.Printf("[%v](%T) -> float64\n", val, val)
		case *int:
			fmt.Printf("[%v](%T) -> *int\n", val, val)
		case string:
			fmt.Printf("[%v](%T) -> string\n", val, val)
		default:
			fmt.Printf("[%v](%T) -> default\n", val, val)
		}
	}
}
