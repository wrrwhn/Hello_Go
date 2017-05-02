package utils

import (
	"flag"
	"fmt"
)

func HelloFlag() {

	url := flag.String("url", "127.0.0.1", "input the url, plz")
	port := flag.Int("port", 8888, "input the port, plz")
	flag.Parse()

	fmt.Printf("%v %v\n", *url, *port)
}
