package main

import (
	"fmt"
    "github.com/jiayanju/gobible/method/set"
)

func main() {
    var x, y set.IntSet
    x.Add(1)
    x.Add(144)
    x.Add(9)
    fmt.Println(x.String())
    fmt.Println(x)
    fmt.Println(&x)
    
    y.Add(9)
    y.Add(43)
    fmt.Println(y.String())
    
    x.UnionWith(&y)
    fmt.Println(x.String())
    
    fmt.Println(x.Has(9), x.Has(123))
    
    fmt.Printf("x length %d\n", x.Len())
    
    x.Remove(9)
    fmt.Printf("x length %d\n", x.Len())
    
    x.Remove(123) 
    fmt.Printf("x length %d\n", x.Len())
    
    x.Clear()
    fmt.Printf("x length %d\n", x.Len())
    
    x.AddAll(1, 4, 9, 123, 12345)
    fmt.Printf("x length %d\n", x.Len())
}