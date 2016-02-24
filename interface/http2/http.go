package main

import (
	"fmt"
	"net/http"
)

const port = 8088

func main() {
	db := database{"shoe": 50, "socks": 5}
	http.ListenAndServe(fmt.Sprintf(":%d", port), db)
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s : %s\n", item, price)
		}

	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s", price)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", r.URL)
	}
}
