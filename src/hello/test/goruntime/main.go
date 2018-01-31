package main

import (
	"fmt"
)

var c chan int = make(chan int)

func main() {
	serial()
}

func serialLoop() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	c <- 0
}

//
func serial() {

	go serialLoop()
	go serialLoop()

	for i := 0; i < 2; i++ {
		<-c
	}
}
