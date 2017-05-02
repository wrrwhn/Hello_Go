package utils

import (
	"fmt"
	"time"
)

// https://my.oschina.net/u/943306/blog/149395
func HelloTime() {

	// helloNow()
	// helloAdd()
	// helloSleep()
	// helloAfter()
	// helloTicker()
	// helloAfterFunc()
	helloBeforeAndAfter()

	// time.format
	// fmt.Println("time.format")

	// time.string.format
	// fmt.Println("time.string.format")

}

func helloNow() {

	fmt.Println("time.now")
	tNow := time.Now()
	fmt.Printf("\t%04d-%02d-%02d\n", tNow.Year(), tNow.Month(), tNow.Day())
}

func helloAdd() {

	fmt.Println("time.Add")
	tNow := time.Now()
	tAfterDay := tNow.Add(time.Hour * 24)
	fmt.Printf("\t%v\n", tNow)
	fmt.Printf("\t%v\n", tAfterDay)
}

func helloSleep() {

	fmt.Println("time.sleep")
	fmt.Printf("\tstart sleep at %v\n", time.Now())
	time.Sleep(200 * time.Microsecond)
	fmt.Printf("\tend sleep at %v\n", time.Now())
}

// 当调用 <-tCAfter 获取管道数据时，期间已经过了超过一秒，里面已然有了数据
func helloAfter() {

	fmt.Println("time.after")
	fmt.Printf("\t %v\n", time.Now())
	tCAfter := time.After(time.Second)
	fmt.Printf("\t %v\n", time.Now())
	time.Sleep(time.Second)
	fmt.Printf("\t %v\n", time.Now())
	<-tCAfter
	fmt.Printf("\t %v\n", time.Now())
}

// ticker= after 可重复使用版本
func helloTicker() {

	fmt.Println("time.ticker")
	fmt.Printf("\tstart triker at %v\n", time.Now())
	timeTick := time.NewTicker(200 * time.Microsecond)
	go func() {
		for tickTime := range timeTick.C {
			fmt.Printf("\t\t %v\n", tickTime)
		}
	}()
	time.Sleep(600 * time.Microsecond)
	fmt.Printf("\tend triker at %v\n", time.Now())
}

// 定时任务
func helloAfterFunc() {

	fmt.Println("time.afterFunc")
	tSubThread := func() {
		fmt.Printf("\t call sub thread func at %v\n", time.Now())
	}
	fmt.Printf("%v\n", time.Now())
	time.AfterFunc(time.Second, tSubThread)
	time.Sleep(1 * time.Second)
	fmt.Printf("%v\n", time.Now())
}

func helloBeforeAndAfter() {
	fmt.Println("time.before&after")
	tNow := time.Now()
	tAfter := tNow.Add(time.Second)
	fmt.Printf("now= \t\t %v\n", tNow)
	fmt.Printf("after= \t %v\n", tAfter)
	var sEqual string
	if tNow.After(tAfter) {
		sEqual = ">"
	} else {
		sEqual = "<="
	}
	fmt.Printf("tNow - tAfter %v 0\n", sEqual)
}
