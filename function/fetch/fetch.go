package main

import (
    "fmt"
	"net/http"
	"path"
	"os"
	"io"
)

func main() {
    for _, url := range os.Args[1:] {
        filename, lenght, err := fetch(url)
        if err != nil {
            fmt.Printf("fetch %s error: %v", url, err)
        }
        fmt.Printf("fileName: %s length: %d", filename, lenght)
    }
}

func fetch(url string) (filename string, length int64, err error)  {
    resp, err := http.Get(url)
    if err != nil {
        return "", 0, err
    }
    defer resp.Body.Close()
    
    local := path.Base(resp.Request.URL.Path)
    if local == "/" {
        local = "index.html"
    }
    
    file, err := os.Create(local)
    if err != nil {
        return "", 0, err
    }
    
    length, err = io.Copy(file, resp.Body)
    if closeErr := file.Close(); err == nil {
        err = closeErr
    }
    
    return local, length, err
    
}