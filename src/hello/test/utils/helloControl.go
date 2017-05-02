package utils

import (
	"fmt"
	"math/rand"
	"runtime"
)

func HelloControl() {

	helloLabel()
}

func helloLabel() {

	for i := 0; i <= 5; i++ {

		if 3 == i {
			goto BREAK_POINT
		}
		fmt.Print(i, " ")
	}

BREAK_POINT:
	fmt.Println("\njump to break_point")
}

func helloFor() {
	for i := 0; i <= 5; i++ {
		fmt.Print(i)
	}
}

func helloForRange() {
	str := "hello 够"
	for idx, rune := range str {
		fmt.Printf("%-2d \t %d \t %U \t '%c' %X\n", idx, rune, rune, rune, []byte(string(rune)))
	}
}

func ifElseTest() {
	// rVal := rand.Intn(100)
	if sSystemOS := runtime.GOOS; sSystemOS == "windows" {
		fmt.Println("Hello windows!")
	} else {
		fmt.Println("Hello Unix!")
	}
}

func Abs() {
	iAbsVal := 1
	fmt.Printf("abs(%v)= %v\n", iAbsVal, abs(iAbsVal))
	iAbsVal = -1
	fmt.Printf("abs(%v)= %v\n", iAbsVal, abs(iAbsVal))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func helloSwitch() {

	iRandVal := rand.Intn(4)
	// iRandVal = 3
	switch iRandVal {
	case 0:
		fmt.Println("0")
	case 1:
		fallthrough
	case 2, 3:
		fmt.Println("1/2/3")
	default:
		fmt.Println("default")
	}
}

func helloConditionSwitch() {

	iRandVal := 10

	switch {
	case iRandVal < 0:
		fmt.Printf("[%v](%T) is negative\n", iRandVal, iRandVal)
	case iRandVal >= 0 && iRandVal < 10:
		fmt.Printf("[%v](%T) is between 0 and 10\n", iRandVal, iRandVal)
	default:
		fmt.Printf("[%v](%T) is greater or equal to 10\n", iRandVal, iRandVal)
	}
}

// 6/ 7/ 8/ default
func testSwitch() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
