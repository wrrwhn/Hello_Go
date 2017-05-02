package utils

import "fmt"

func HelloArray() {
	arrayCopy()
}

func writeAndReadArray() {
	const ARRAY_LENGTH int = 5
	// init-4
	var arr [ARRAY_LENGTH]int

	for i := 0; i < ARRAY_LENGTH; i++ {
		arr[i] = i
	}

	for idx, val := range arr {
		fmt.Printf("Array[%d]= %d\n", idx, val)
	}
}

func arrayInit() {
	// init-2/3
	// sArr := [...]string{"yqj", "lmm"}
	sArr := [...]string{"yqj", "lmm"}
	for _, val := range sArr {
		fmt.Println(val)
	}
}

//init-1
var aInit = []int{0, 2: 1, 4: 2}

func arrayCopy() {

	printArray("Init", aInit)
	arrayParanCopy(aInit)
	aInit[0] += 1
	printArray("\nInit", aInit)
	arrayPoint(&aInit)
	printArray("\nInit", aInit)
}

func arrayPoint(arr *[]int) {

	arrLen := len(*arr)
	(*arr)[arrLen-1] = 9
	printArray("Array Point", *arr)
}

func arrayParanCopy(arr []int) {
	arr[len(arr)-1] = 9
	printArray("Array Param Copy", arr)
}

func printArray(title string, arr []int) {

	fmt.Printf("%v\n\t", title)
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
