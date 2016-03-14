package main

import (
	"flag"
	"net"
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
    defer c.Close()
    
    go mustCopy(os.Stdout, c)
    mustCopy(c, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader)  {
    _, err := io.Copy(dst, src)
    if err != nil {
        log.Fatal(err)
    }
}