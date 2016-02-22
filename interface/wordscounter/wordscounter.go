package main

import (
	"bufio"
	"fmt"
	"strings"
)

// WordConter count words in the line
type WordConter int

func (w *WordConter) Write(p []byte) (n int, err error) {
	if len(p) > 0 {
		*w += WordConter(1)
	}
	return len(p), nil
}

func main() {
	line := "hello hello word world"

	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)

	var w WordConter

	for scanner.Scan() {
		w.Write(scanner.Bytes())
	}

	fmt.Println(w)
}
