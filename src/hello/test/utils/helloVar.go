package utils

import (
    "fmt"
    "runtime"
    "os"
)

func HelloVar(){
    var str string = runtime.GOOS
    fmt.Printf("System = %s\n", str)
    path := os.Getenv("PATH")
    fmt.Printf("Path = %s\n", path)
}