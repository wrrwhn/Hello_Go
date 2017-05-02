package utils

import (
	"strings"
)

var compareStr string

func init() {
	compareStr = "test"
}

func HelloTest(strVal string) bool {
	return 0 == strings.Compare(strVal, compareStr)
}
