package main

import (
	"net"
	"log"
	"os"
	"io"
)

func main() {
    c, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
    defer c.Close()
    
    mustCopy(os.Stdout, c)
}

func mustCopy(dst io.Writer, src io.Reader)  {
    _, err := io.Copy(dst, src)
    if err != nil {
        log.Fatal(err)
    }
}