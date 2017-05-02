package utils

import "fmt"

// Public
func HelloSlice() {
	copyAndAppendSlice()
}

// private
func initSlice() {

	// init
	var array [10]int
	var slice []int = array[5:10]

	// init.data
	for i := 0; i < 10; i++ {
		array[i] = i
	}

	// read
	printSlice(&slice)

	// slice of slice
	slicePart := slice[0 : len(slice)-2]
	printSlice(&slicePart)
}

func makeSlice() {

	// init
	slice := make([]int, 10)
	for i := 0; i < len(slice); i++ {
		slice[i] = i * i
	}

	printSlice(&slice)
}

func copyAndAppendSlice() {
	initSlice := []int{1, 2, 3}
	copySlice := make([]int, 10)
	copyCnt := copy(copySlice, initSlice)
	fmt.Printf("Copy.cnt= %d\n", copyCnt)
	printSlice(&copySlice)

	appendSlice := append(copySlice, 998)
	printSlice(&copySlice)
	printSlice(&appendSlice)
}

// common
func printSlice(slice *[]int) {

	fmt.Printf("\nslice.len= %d, cap= %d, datas= [", len(*slice), cap(*slice))
	for _, val := range *slice {
		fmt.Printf("%d ", val)
	}
	fmt.Print("]\n")
}

// test
func testSlice() {
	array := []int{1, 2, 3, 4}
	slice := array[2:]
	printArray("array.init", array)
	printSlice(&slice)

	slice[1] = 9
	printArray("array.update", array)
	printSlice(&slice)
}
