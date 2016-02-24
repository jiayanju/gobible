package main

import (
	"net/http"
	"fmt"
)

const port = 8088

func main() {
    db := database{"shoe": 50, "socks": 5}
   http.ListenAndServe(fmt.Sprintf(":%d", port), db)
}

type dollars float32

func (d dollars) String() string  {
    return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
    for item, price := range db {
        fmt.Fprintf(w, "%s : %s\n", item, price)
    }
}
