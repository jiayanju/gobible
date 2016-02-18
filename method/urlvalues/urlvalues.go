package main

import (
	"net/url"
	"fmt"
)

func main() {
    m := url.Values{"lang": {"en"}}
    m.Add("item", "1")
    m.Add("item", "2")
    
    fmt.Println(m.Get("lang"))
    fmt.Println(m.Get("q"))
    fmt.Println(m.Get("item"))
    fmt.Println(m["item"])
    
    // m = nil
    // fmt.Println(m.Get("item"))
    // m.Add("item", "3")
    
    reassignValues(m)
    fmt.Printf("Reassign values out of func: %v\n", m)
    
    changeValues(m)
    fmt.Printf("Values is changed: %v", m)
}

func reassignValues(values url.Values)  {
    n := url.Values{"test": {"value"}}
    values = n
    fmt.Printf("Reassign values In func: %v\n", values)
}

func changeValues(values url.Values)  {
    values.Add("item1", "1")
    values.Del("item")
}
