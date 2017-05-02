package utils

import (
	"fmt"
	"regexp"
	// "strconv"
)

func HelloRegexp() {
	helloRegexp()
}

func helloRegexp() {

	search := "John: 2578.34 William: 4567.23 Steve: 5632-18"
	pattern := "([0-9]+)(.)([0-9]+)"

	// floatFmt := func(str string) string {

	// 	val, _ := strconv.ParseFloat(str, 32)
	// 	return strconv.FormatFloat(val, 'f', 2, 32)
	// }
	fmt.Println(search)
	if ok, _ := regexp.Match(pattern, []byte(search)); !ok {
		return
	}

	// 0.compile
	regex, _ := regexp.Compile(pattern)
	// 1.find
	strs := regex.FindAllString(search, -1)
	fmt.Printf("\tMatches= %v\n", strs)
	// 1.replace
	search = regex.ReplaceAllString(search, "$3$2$1")
	fmt.Println(string(search))
}
