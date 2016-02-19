package main

import (
	"fmt"
    "github.com/jiayanju/gobible/method/point"
	"time"
)

func main() {
    p := point.Point{X: 1, Y: 2}
    q := point.Point{X: 4, Y: 6}
    
    distanceFromP := p.Distance
    fmt.Println(distanceFromP(q))
    fmt.Printf("%T\n", distanceFromP)
    
    var origin point.Point
    fmt.Println(distanceFromP(origin))
    
    scaleP := p.ScaleBy
    scaleP(2)
    fmt.Println(p)
    
    scaleP(3)
    fmt.Println(p)
    
    scaleP(10)
    fmt.Println(p)
    
    r := new(Rocket)
    time.AfterFunc(10 * time.Second, r.launch)
    
    // make console wait for the timer to see the result
    time.Sleep(20 * time.Second)
}

//Rocket rocket test
type Rocket struct {}

func (r *Rocket) launch()  {
    fmt.Println("Launch....")
}
