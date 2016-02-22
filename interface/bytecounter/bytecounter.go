package main

import (
	"fmt"
	"io"
	"os"
)

//ByteCounter used for count byte number
type ByteCounter int

func (c *ByteCounter)Write(p []byte) (n int, err error)  {
    *c += ByteCounter(len(p))
    return len(p), nil
}

func main() {
    var c ByteCounter
    c.Write([]byte("hello"))
    fmt.Println(c)
    
    c = 0
    var name = "George"
    fmt.Fprintf(&c, "hello, %s", name)
    fmt.Println(c)
    
    out := os.Stdout;
    c, count := CountingWriter(out)
    out.Write([]byte("hello"))
    fmt.Println(c)
    
    out.Write([]byte("world"))
    fmt.Println(c)
}

//CountingWriter used to return new writer and wirte bytes
func CountingWriter(w io.Writer) (io.Writer, *int64)  {
    var c ByteCounter
    return c, *c
    
}