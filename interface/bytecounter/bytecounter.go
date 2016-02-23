package main

import (
	"fmt"
	"io"
	"os"
)

//ByteCounter used for count byte number
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (n int, err error)  {
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
    bc, count := CountingWriter(out)
    bc.Write([]byte("hello"))
    fmt.Println(*count)
    
    bc.Write([]byte("world\n"))
    fmt.Println(bc)
    fmt.Println(*count)
}

//CountingWriter used to return new writer and wirte bytes
func CountingWriter(w io.Writer) (io.Writer, *int64)  {
    bc := ByteCounter2 {
        Writer: w,
        Count: 0,
    }
    
    return &bc, &bc.Count
}

//ByteCounter2 count bytes
type ByteCounter2 struct {
    Writer io.Writer
    Count int64
}

func (b *ByteCounter2) Write(p []byte) (n int, err error)  {
    n, err = b.Writer.Write(p)
    if err != nil {
        return n, err
    }
    
    b.Count += int64(n)
    return n, err
}