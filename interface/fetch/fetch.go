package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
    for _, url := range os.Args[1:] {
        res, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch %s: %s\n", url, err)
            os.Exit(1)
        }
        defer res.Body.Close()
        
        b, err := ioutil.ReadAll(res.Body)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch %s: Reading error %s\n", url, err)
            os.Exit(1)
        }
        fmt.Printf("%s", b)
    }
}