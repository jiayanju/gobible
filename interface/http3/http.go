package main

import (
	"fmt"
	"net/http"
	"log"
)

const port = 8088

func main() {
	db := database{"shoe": 50, "socks": 5}
    serverMux := http.NewServeMux()
    serverMux.HandleFunc("/list", db.list)
    serverMux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), serverMux))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s : %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s", price)
}
