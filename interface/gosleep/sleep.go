package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var period = flag.Duration("period", 1*time.Second, "input the period seconds to sleep")
	flag.Parse()
	fmt.Printf("Sleep for : %v...", *period)
	time.Sleep(*period)
	fmt.Println("Done")
}
