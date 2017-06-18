package utils

import (
	"fmt"
	"time"
)

func HelloChannel() {
	helloChannelRange()
}

func helloGoruntime() {

	c := make(chan string)

	go func() { c <- "one" }()
	go func() { c <- "two" }()
	go func() { c <- "three" }()

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(len(c))
}

func helloChannelBuf() {
	c := make(chan string, 2)

	c <- "one"
	c <- "two"
	fmt.Println(len(c))
	<-c
	c <- "two"
	close(c)

	for v := range c {
		fmt.Println(v)
	}
}

func helloChannelSync() {

	done := make(chan bool)
	go worker(done)
	<-done
}
func worker(done chan bool) {

	fmt.Print("working...")
	time.Sleep(500 * time.Millisecond)
	fmt.Print("done")

	done <- true
}

// point the direction for channel to read or write
func helloChannelDirection() {

	ping := make(chan string)
	pong := make(chan string)
	go push(ping, "msg")
	go send(ping, pong)
	fmt.Print(<-pong)
}
func push(ping chan<- string, msg string) {
	ping <- msg
}
func send(ping <-chan string, pong chan<- string) {
	msg := <-ping
	pong <- msg
}

// select{} wait for channel
func helloChannelSelect() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "goruntime-1"
	}()
	go func() {
		time.Sleep(time.Second * 4)
		c2 <- "goruntime-2"
	}()

	fmt.Println("start at", time.Now())
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-c1:
			fmt.Println(time.Now(), m1)
		case m2 := <-c2:
			fmt.Println(time.Now(), m2)
		case <-time.After(time.Second * 2):
			fmt.Println("timeout 2")
		}
	}
}

func helloChannelClose() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// more will be false when jobs channel closed
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}
	close(jobs)
	// read after close
	fmt.Println(<-jobs)
	fmt.Println("sent all jobs")

	<-done
}

func helloChannelRange() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for val := range queue {
		fmt.Println(val)
	}
}
