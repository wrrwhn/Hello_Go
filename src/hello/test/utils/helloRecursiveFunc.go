package utils

import "fmt"

func HelloRecursive() {
	fibonacci(3)
	for i := 0; i <= 5; i++ {
		fmt.Printf("fibonacci(%d)= %d\n", i, fibonacci(i))
	}
}

func fibonacci(i int) int {
	if i <= 0 {
		return 1
	} else {
		return fibonacci(i-1) + fibonacci(i-2)
	}
}
