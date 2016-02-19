package point

import (
	"math"
	"image/color"
)

// Point geomath point
type Point struct {X, Y float64}

//Distance return distance of two points
func (p Point) Distance(q Point) float64  {
    return math.Hypot(q.X - p.X, q.Y - p.Y)
}

//ScaleBy scale pointer by factor
func (p *Point) ScaleBy(factor float64)  {
    p.X *= factor
    p.Y *= factor
}

//ColoredPoint point with color
type ColoredPoint struct {
    Point
    Color color.RGBA
}

//ColoredPointPtr colored pointer with pointer anonymouse field
type ColoredPointPtr struct {
    *Point
    Color color.RGBA
}