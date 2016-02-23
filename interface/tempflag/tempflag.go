package main

import (
    "github.com/jiayanju/gobible/interface/tempconv"
	"fmt"
	"flag"
)

func main() {
    var temp = tempconv.CelsiusFlag("temp", 20.0, "tempearture")
    flag.Parse()
    fmt.Println(*temp)
}