package main

import (
	"bufio"
	"strings"
	"fmt"
)

//LineCounter counte line number
type LineCounter int

func (l *LineCounter) Write(p []byte) (n int, err error)  {
    *l += LineCounter(1)
    return len(p), nil
}

func main() {
    lines := `Hello
            Hello
            World
            World`
    
    scanner := bufio.NewScanner(strings.NewReader(lines))
    scanner.Split(bufio.ScanLines)
    
    var l LineCounter
    for scanner.Scan() {
        l.Write(scanner.Bytes())
    }
    
    fmt.Println(l)
}