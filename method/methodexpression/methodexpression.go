package main

import (
	"fmt"
	"github.com/jiayanju/gobible/method/point"
)

func main() {
    p := point.Point{X: 1, Y: 2}
    q := point.Point{X: 4, Y: 6}
    
    distance := point.Point.Distance
    fmt.Println(distance(p, q))
    fmt.Printf("%T\n", distance)
    
    scale := (*point.Point).ScaleBy
    scale(&p, 2)
    fmt.Println(p)
    fmt.Printf("%T\n", scale)
}