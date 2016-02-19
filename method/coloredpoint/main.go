package main

import (
	"fmt"
	"image/color"
)

func main() {
    var cp ColoredPoint
    cp.X = 1
    fmt.Println(cp.Point.X)
    
    cp.Point.Y = 2
    fmt.Println(cp.Y)
    
    red := color.RGBA{255, 0, 0, 255}
    blue := color.RGBA{0, 0, 255, 255}
    var p = ColoredPoint{Point{1, 1}, red}
    var q = ColoredPoint{Point{5, 4}, blue}
    fmt.Println(p.Distance(q.Point))
    
    p.ScaleBy(2)
    q.ScaleBy(2)
    fmt.Println(p.Distance(q.Point))
    
    // Compile Error: cannot use q(ColoredPoint) as Point
    // p.Distance(q)
    
    var cpp ColoredPointPtr
    // runtime error 
    // cpp.X = 1 
    // cpp.Point.X = 1 
    // (*cpp.Point).X = 1
    
    cpp.Point = &Point{1, 2}
    fmt.Println(cpp.Point.X)

    var pptr = ColoredPointPtr{&Point{1, 1}, red}
    var qptr = ColoredPointPtr{&Point{5, 4}, blue}
    fmt.Println(pptr.Distance(*qptr.Point))
    
    pptr.Point = qptr.Point
    qptr.ScaleBy(2)
    fmt.Println(*pptr.Point, *qptr.Point)
}