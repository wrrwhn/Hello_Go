package utils

import (
	"fmt"
	"sort"
)

func HelloSort() {
	helloInverted()
}

func helloSort() {
	var iArr []int = []int{1, 3, 2, 4}
	printArray(fmt.Sprintf("Array[sorted= %v]", sort.IntsAreSorted(iArr)), iArr)

	sort.Ints(iArr)
	printArray(fmt.Sprintf("Array[sorted= %v]", sort.IntsAreSorted(iArr)), iArr)

	var iIdx int = 3
	fmt.Printf("Array.indexOf(%d)= %d\n", iIdx, sort.SearchInts(iArr, iIdx))
}

func helloInverted() {
	array := []int{1, 2, 3, 4, 5}
	printArray("Init", array)

	invertedArray(&array)
	printArray("Inverted", array)
}

func invertedArray(array *[]int) {
	arrLen := len(*array)
	for i := 0; i <= arrLen/2; i++ {
		exchangeIdx := arrLen - 1 - i
		(*array)[i], (*array)[exchangeIdx] = (*array)[exchangeIdx], (*array)[i]
	}
}
