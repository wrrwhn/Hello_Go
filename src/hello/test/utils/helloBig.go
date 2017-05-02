package utils

import (
	"fmt"
	"math"
	"math/big"
)

func HelloBig() {
	helloBig()
}

func helloBig() {
	biMaxVal := big.NewInt(math.MaxInt64)
	biVal := big.NewInt(1)
	biOne := big.NewInt(1)
	biTwo := big.NewInt(2)
	biThree := big.NewInt(3)
	biVal.Mul(biTwo, biThree).Div(biOne, biOne).Add(biTwo, biThree)
	fmt.Println(biMaxVal, biVal)

	riMaxVal := big.NewRat(math.MaxInt64, 1956)
	riMinVal := big.NewRat(-1956, math.MaxInt64)
	riVal1 := big.NewRat(19, 56)
	riVal2 := big.NewRat(111, 222)
	riVal3 := big.NewRat(1, 1)
	riVal3.Mul(riMaxVal, riMinVal).Add(riVal3, riVal1).Mul(riVal3, riVal2)
	fmt.Println(riVal3)
}
