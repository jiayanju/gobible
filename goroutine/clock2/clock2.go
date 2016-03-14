package main

import (
	"net"
	"log"
	"io"
	"time"
	"flag"
	"fmt"
)

var port = flag.Int("port", 8080, "server port listening")

func main() {
    flag.Parse()
    log.Printf("port: %d", *port)
    listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
    if err != nil {
        log.Fatal(err)
    }
    
    for {
        c, err := listener.Accept()
        if err != nil {
            log.Fatalf("Connection error: %v", err)
            continue
        }
        
        go handleConn(c)
    }
}

func handleConn(c net.Conn)  {
    defer c.Close()
    for {
        _, err := io.WriteString(c, time.Now().Format("Mon Jan _2 15:04:05 2006\n"))
        if err != nil {
            return
        }
        time.Sleep(1 * time.Second)
    }
}


