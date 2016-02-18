package main

import (
    "math"
	"fmt"
)

func main() {
    p := Point{1, 2}
    q := Point{x: 4, y: 6}
    fmt.Println(p.Distance(q))
    fmt.Println(Distance(p, q))
    
    fmt.Printf("function type %T\n", Distance)
    fmt.Printf("method type %T\n", Point.Distance)
    fmt.Printf("point type %T\n", p)
    
    path1 := Path{
        Point{1, 2},
        Point{4, 6},
    }
    
    fmt.Printf("type : %T\n", path1)
    
    sum := path1.Distance()
    fmt.Printf("Distance of points %v : %f\n", path1, sum)
    
    path2 := Path {
        {1, 2},
        {3, 4}}
    path3 := Path {
        {1, 2},
        {3, 4},
    }
    sum = path2.Distance()
    fmt.Printf("Distance of points %v : %f\n", path2, sum)
    sum = path3.Distance()
    fmt.Printf("Distance of points %v : %f\n", path3, sum)
    
    // test method receiver value possiblility
    fmt.Println("----------------- Test method receiver value ----------------")
    r := &Point{1, 2}
    r.ScaleBy(2)
    fmt.Println(r)
    fmt.Println(*r)
    
    pptr := &p
    pptr.ScaleBy(2)
    fmt.Println(pptr)
    fmt.Println(*pptr)
    
    (&p).ScaleBy(2)
    fmt.Println(p)
    
}

// P for testing
type P *int

/* method declaration are not permitted on named type that are themselves pointer type
func (P) f()  {
    
}
*/


// Point Type
type Point struct {x, y float64}

// Distance Return distance of tow point
func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.x - p.x, q.y - p.y)
}

// ScaleBy - expain pointer receiver
func (p *Point) ScaleBy(factor float64)  {
    p.x *= factor
    p.y *= factor
}

// Distance return distance of two points
func Distance(p, q Point) float64 {
    return math.Hypot(q.x - p.x, q.y - p.y)
}

// Path Point slice 
type Path []Point

// Distance return the distance of slice point
func (path Path) Distance() float64  {
    sum := 0.0
    for i := range path {
        if i > 0 {
            sum += path[i - 1].Distance(path[i])
        }
    }
    return sum
}