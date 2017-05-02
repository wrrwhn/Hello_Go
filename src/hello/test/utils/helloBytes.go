package utils

import (
	"bytes"
	"fmt"
)

func HelloBytes() {
	helloBytes()
}
func helloBytes() {

	var buffer bytes.Buffer
	array := []string{"https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/07.2.md", "if s, ok := getNextString(); ok { //method getNextString() not shown here", "或者通过函数：func NewBuffer(buf []byte) *Buffer", "上一节：声明和初始化"}

	for i := 0; i < len(array); i++ {
		buffer.WriteString(array[i])
	}
	fmt.Println(buffer.String())
}
