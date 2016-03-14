package main

import (
	"net"
	"flag"
	"fmt"
	"log"
	"io"
	"os"
)

var port = flag.Int("port", 8080, "server port")

func main() {
    flag.Parse()
    c, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", *port))
    if err != nil {
        log.Fatal(err)
    }
    
    done := make(chan struct{})
    go func ()  {
        io.Copy(os.Stdout, c)
        log.Println("Done")
        done <- struct{}{}
    }()
    mustCopy(c, os.Stdin)
    c.Close()
    <- done
}

func mustCopy(dst io.Writer, src io.Reader)  {
    _, err := io.Copy(dst, src)
    if err != nil {
        log.Fatal(err)
    }
}