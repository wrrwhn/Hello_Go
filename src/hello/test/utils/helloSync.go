package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

func HelloSync() {
	helloWaitGroup()
}

func helloOnce() {

	// init
	var once sync.Once
	done := make(chan int)
	printer := func() {
		fmt.Println("printer().once.Do()")
	}
	printerArgs := func(val int) {
		fmt.Printf("\tprinter(%v)\n", val)
	}

	// call
	for i := 0; i < 10; i++ {

		// 纤程
		go func() {
			once.Do(printer)
			done <- i
		}()
		go func(val int) {
			printerArgs(val)
		}(i)
		// println(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("\tdone[%v]= %v\n", i, <-done)
	}
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}
func bufLog(w io.Writer, key, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	b.WriteByte('\n')
	w.Write(b.Bytes())
	bufPool.Put(b)
}
func hisLog(w io.Writer) {
	for i := 0; i < 10; i++ {

		buffer := bufPool.Get().(*bytes.Buffer)
		var val string = fmt.Sprintf("[%v] %v\n", i, buffer.String())
		w.Write([]byte(val))
	}
}
func helloPool() {

	bufLog(os.Stdout, "path", "/search?q=golang")
	hisLog(os.Stdout)

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	ptrA := &a
	fmt.Println(ptrA)
	for i, v := range ptrA {
		a[3] = 100
		if 3 == i {
			fmt.Println(i, v)
		}
	}
}

var wg sync.WaitGroup
var urls = []string{
	"http://www.baidu.com",
	"https://www.sogou.com/",
	"http://www.bing.com",
}

func helloWaitGroup() {

	fmt.Println(timeNow().UTC().Format(time.RFC3339Nano))
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fmt.Printf("GET %v\n", url)
			time.Sleep(time.Second)
		}(url)
	}

	fmt.Println(timeNow().UTC().Format(time.RFC3339Nano))
	wg.Wait()
	fmt.Println(timeNow().UTC().Format(time.RFC3339Nano))
}
