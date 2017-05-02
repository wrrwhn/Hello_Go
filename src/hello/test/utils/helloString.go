package utils

import (
	"fmt"
	"unicode/utf8"
)

func HelloString() {
	str := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("[%s].bytes.length= %d\n", str, len(str))
	fmt.Printf("[%s].chars.length= %d\n", str, utf8.RuneCountInString(str))
}
