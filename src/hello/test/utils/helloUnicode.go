package utils

import (
	"fmt"
	"unicode"
)

func HelloUnicode() {
	cLetter := 'L'
	cNum := '9'
	cSpace := '_'

	fmt.Printf("%c [letter= %t, digit= %t, space= %t]\n", cLetter, unicode.IsLetter(cLetter), unicode.IsDigit(cLetter), unicode.IsSpace(cLetter))
	fmt.Printf("%c [letter= %t, digit= %t, space= %t]\n", cNum, unicode.IsLetter(cNum), unicode.IsDigit(cNum), unicode.IsSpace(cNum))
	fmt.Printf("%c [letter= %t, digit= %t, space= %t]\n", cSpace, unicode.IsLetter(cSpace), unicode.IsDigit(cSpace), unicode.IsSpace(cLetter))
}
