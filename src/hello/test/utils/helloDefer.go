package utils

import (
	"fmt"
	"log"
)

func HelloDefer() {
	anonymousFunc()
}

// defer
func helloDefer() {
	testDefer(1)
	defer testDefer(2)
	defer testDefer(3)
	testDefer(4)
}

func testDefer(iVal int) {
	fmt.Print(iVal)
}

// defer 匿名函数
func anonymousFunc() {
	defer func() {
		log.Println("anonymous.defer")
	}()
	log.Println("anonymousFunc() finish")
}
