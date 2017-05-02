package utils

import (
	"fmt"
	"strings"
)

func HelloStrings() {

	// init
	var str string = "4.7 strings 和 strconv 包"
	sPrefix := "4.7"
	sSuffix := "suffix"
	sContain := "strings"

	// query
	fmt.Printf("[%s].prefix[%s] = %t\n", str, sPrefix, strings.HasPrefix(str, sPrefix))
	fmt.Printf("[%s].suffix[%s] = %t\n", str, sSuffix, strings.HasSuffix(str, sSuffix))
	fmt.Printf("[%s].contain[%s] = %t\n", str, sContain, strings.Contains(str, sContain))
	fmt.Printf("[%s].index[%s] = %d\n", str, sContain, strings.Index(str, sContain))
	fmt.Printf("[%s].count[%s] = %d\n", str, sPrefix, strings.Count(str, sPrefix))

	// replace
	fmt.Printf("[%s].replace[%s] to [%s] = %s\n", str, sPrefix, sSuffix, strings.Replace(str, sPrefix, sSuffix, -1))
	fmt.Printf("[%s].ToUpper = %s\n", str, strings.ToUpper(str))
	fmt.Printf("[%s].ToLower = %s\n", str, strings.ToLower(str))
	str = "    " + str + "  "
	fmt.Printf("[%s].trimSpace = [%s]\n", str, strings.TrimSpace(str))
	fmt.Printf("[%s].trimLeft = [%s]\n", str, strings.TrimLeft(str, " "))
	str = "Chicken Duck Dog Egg Mobile"
	splits := strings.Fields(str)
	sJoin := ", "
	fmt.Printf("[%s].split().join[%s] = %s\n", str, sJoin, strings.Join(splits, sJoin))
}
