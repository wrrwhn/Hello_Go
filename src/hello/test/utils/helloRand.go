package utils

import (
	"fmt"
	"math/rand"
)

func HelloRand() {
	fmt.Printf("%d \n", rand.Int())
	fmt.Printf("%d \n", rand.Intn(8))
	// rand.Float32()= [0.0, 1.0)
	fmt.Printf("%2.2f", rand.Float32()*100)

	sSource := rand.NewSource(100)
	rNew := rand.New(sSource)
	fmt.Print(rNew.Intn(100))
}
